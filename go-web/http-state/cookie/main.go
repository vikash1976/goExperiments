package main

import(
    "net/http"
    "log"
    "strconv"
    "fmt"
)

//var counter int
func main() {
    http.HandleFunc("/", handleRoot)
    
    log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
    log.Println("In  handleRoot")
    myCounter, err := req.Cookie("visitCount")
    if err == http.ErrNoCookie {
        log.Println("Cookie Not Found")
        myCounter =  &http.Cookie {
            Name: "visitCount",
            Value: "0",
        }

    }
    count, _ := strconv.Atoi(myCounter.Value)
    count++
    myCounter.Value = strconv.Itoa(count)
    http.SetCookie(w, myCounter)
    fmt.Fprintf(w, "You are seeing me %s times", myCounter.Value)
    //counter++
    //http.Redirect(w, req, "/set", http.StatusTemporaryRedirect)
}

