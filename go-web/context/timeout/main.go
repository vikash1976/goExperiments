package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// request processor
func foo(w http.ResponseWriter, req *http.Request) {
	// getting request context
    ctx := req.Context()
    

    // adding identifier to req context for other methods in the chain to access it
	ctx = context.WithValue(ctx, "userID", 777)
	// pretend that this db function is to fetch data for given identifier
	results, err := dbAccess(ctx)
	if err != nil {
        fmt.Println("Flagged Here too")
        // primary DB timeout request process expected in 1 second
        results, err = dbAccessAnother(ctx) // kind of handing off to secondary DB instance for request processing
        if err != nil {
            fmt.Println("Flagged Here finally")
            http.Error(w, "Request timed out. Server busy", http.StatusRequestTimeout)
            return
        }
	}

	fmt.Fprintln(w, results)
}

// may be the primary DB instance used here to fetch data
func dbAccess(ctx context.Context) (int, error) {
    ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		// function to fetch data for provided uid (assume)
		uid := ctx.Value("userID").(int)
		time.Sleep(2 * time.Second)

		// check to make sure we're not running unnecessarily
		// if ctx.Done()
		if ctx.Err() != nil {
            fmt.Println("Flagged Here")
			return
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done():
        fmt.Println("Done Here")
        return 0, ctx.Err()
    case i := <-ch:
		return i, nil
	}
}

// may be a secondary DB instance used here to fetch data
func dbAccessAnother(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // time out changed on context, should respond in 3 seconds
	defer cancel()

	ch := make(chan int)

	go func() {
		// function to fetch data for provided uid (assume)
		uid := ctx.Value("userID").(int)
		time.Sleep(2 * time.Second)
        fmt.Println("Timedout Here Another")

		// check to make sure we're not running in vain
		// if ctx.Done()
		if ctx.Err() != nil {
            fmt.Println("Flagged Here Another")
			return
		}
        
		ch <- uid
	}()

	select {
	case <-ctx.Done():
    // we will land here when time out expires before go func writes the uid onto channel
        fmt.Println("Done Here Another")
        return 0, ctx.Err()
    
    // if we come here within expected time out then req is processed fine. For this
    // instance its line# 82's Sleep controls the behavior
    case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}