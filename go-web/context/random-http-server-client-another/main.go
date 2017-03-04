package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

type responseDetails struct {
	r   *http.Response
	err error
}
type configuration struct {
	Timeout                 time.Duration
	UnpredictableServerURL  string
	UnpredictableServerURL2 string
	UnpredictableServerURL3 string
}

var config configuration

// initializer: reads configuration file to assign values to prop vars
func init() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Configuration file not fount")
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	config = configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Config:", config)
}

// main function
func main() {
	http.HandleFunc("/getMeSome", requestHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8090", nil)
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
	// context with timeout of 2 seconds, will wait for work to be completed within 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout*time.Second)
	defer cancel() // before request handler completes it calls cancel
	// function so that context is canclelled and every one is notified

	fmt.Println("Doing some work")
	workResp := make(chan string)
	go work(ctx, workResp, cancel)
	msgFromWork := <-workResp // waiting for work function to return

	fmt.Fprintf(w, "Response: %v.\n", msgFromWork)

}

// to do actual work of request processing - creates Http client Request to our
// unpredicatble server
func work(ctx context.Context, workResp chan string, cancel context.CancelFunc) {
	ctxCancel, canc := context.WithCancel(ctx)
	defer canc()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	// channel between work and go routine that will be
	// responsible to fetch data from unpredicatble server
	c := make(chan responseDetails)
	cErr := make(chan responseDetails)

	req, _ := http.NewRequest("GET", config.UnpredictableServerURL, nil)
	req1, _ := http.NewRequest("GET", config.UnpredictableServerURL2, nil)
	req2, _ := http.NewRequest("GET", config.UnpredictableServerURL3, nil)
	// hands off the responsibility to fetch data from unpredictable server
	// in this go routine and waits on channel
	go func() {
		resp, err := client.Do(req)
		fmt.Println("Doing http request is hard sometimes")
		// packing resp and err in anonymous struct
		pack := responseDetails{resp, err}
		// when because of timeout tr.CancelRequest called, resp will be nil and err non-nil
		// when server responds in time, resp will be non-nil and err nil - see the console output.
		//fmt.Printf("Writing client.Do results - \nResponse: %v AND Error: %v\n", pack.r, pack.err)

		select {
		case <-ctxCancel.Done():
			fmt.Println("No need to continue...")
			return
		default:
			// putting resp and err as responseDetails struct on channel
			if resp != nil {
				c <- pack
				return

			}
			fmt.Println("Error...", err)
			cErr <- pack
		}

	}()
	go func() {
		resp, err := client.Do(req1)
		fmt.Println("Doing http request is hard sometimes")
		// packing resp and err in anonymous struct
		pack := responseDetails{resp, err}
		// when because of timeout tr.CancelRequest called, resp will be nil and err non-nil
		// when server responds in time, resp will be non-nil and err nil - see the console output.
		//fmt.Printf("Writing client.Do results - \nResponse: %v AND Error: %v\n", pack.r, pack.err)

		select {
		case <-ctxCancel.Done():
			fmt.Println("No need to continue...")
			return
		default:
			// putting resp and err as responseDetails struct on channel
			if resp != nil {
				c <- pack
				return

			}
			fmt.Println("Error...", err)
			cErr <- pack
		}

	}()
	go func() {
		resp, err := client.Do(req2)
		fmt.Println("Doing http request is hard sometimes")
		// packing resp and err in anonymous struct
		pack := responseDetails{resp, err}
		// when because of timeout tr.CancelRequest called, resp will be nil and err non-nil
		// when server responds in time, resp will be non-nil and err nil - see the console output.
		//fmt.Printf("Writing client.Do results - \nResponse: %v AND Error: %v\n", pack.r, pack.err)

		select {
		case <-ctxCancel.Done():
			fmt.Println("No need to continue...")
			return
		default:
			// putting resp and err as responseDetails struct on channel
			if resp != nil {

				c <- pack
				return

			}
			fmt.Println("Error...", err)
			cErr <- pack
		}

	}()

	select {
	case ok := <-c:
		err := ok.err
		resp := ok.r
		// in case error sent by unpredicatble server
		if err != nil {
			fmt.Println("Error------------ ", err)
			workResp <- err.Error() // to get string values of error
			return                  // no point trying to read from response
		}

		defer resp.Body.Close()
		out, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Server Response: %s\n", out)
		tr.CancelRequest(req) // cancelling http client request, so that client.Do returns
		// with  Get http://localhost:8080: net/http: request canceled Error
		tr.CancelRequest(req1)
		tr.CancelRequest(req2)
		canc()
		//cancel()
		workResp <- string(out)

		return
	case <-ctx.Done():
		fmt.Println("Explicit Cancel called")
		tr.CancelRequest(req) // cancelling http client request, so that client.Do returns
		// with  Get http://localhost:8080: net/http: request canceled Error
		tr.CancelRequest(req1)
		tr.CancelRequest(req2)
		//<-c // Wait for client.Do
		//<-c
		//<-c
		fmt.Println("Cancelling the context, it has timed out")
		workResp <- ctx.Err().Error()
		return
		// selecting on channel expecting results from go routine
		// to which it handed off the task to fetch data from unpredictable server

	case <-cErr:
	//close(cErr)
		var errStr []string
		for n := range cErr {
			fmt.Printf("Error000000%v\n", n.err.Error())
			errStr = append(errStr, n.err.Error())
		}
		
		fmt.Println("Am i here!!!")
		workResp <- strings.Join(errStr, "")
		return
	}

}
