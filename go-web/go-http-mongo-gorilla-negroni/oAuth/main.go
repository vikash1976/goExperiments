package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	//"log"
    "net/http"
    "io/ioutil"
)

func main() {
    
    googleconf := &oauth2.Config{
        ClientID:     "ztObSscrjGBe16qOc3ZNRvIzy",
        ClientSecret: "BmPCldhh60LBJ90XhQ2ULbCaJPBlsiWHibQU3USkS7yYiHPUXS",
        RedirectURL:  "http://localhost:300/googlelogin",
        /*Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.profile",
            "https://www.googleapis.com/auth/userinfo.email",
        }*/,
        Endpoint: github.Endpoint,
    }

http.HandleFunc("/googleloginrequest", func(response http.ResponseWriter, request *http.Request) {

    url := googleconf.AuthCodeURL("state")
    fmt.Printf("Visit the URL for the auth dialog: %v", url)

    http.Redirect(response, request, url, http.StatusTemporaryRedirect)

})

http.HandleFunc("/googlelogin", func(response http.ResponseWriter, request *http.Request) {

        authcode := request.FormValue("code")

        tok, err := googleconf.Exchange(oauth2.NoContext, authcode)
        if err != nil {
            fmt.Println("err is", err)
        }

        fmt.Println("token is ", tok)
        resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)

        defer resp.Body.Close()
        contents, err := ioutil.ReadAll(resp.Body)
        response.Write(contents)
    })
   http.ListenAndServe(":8080", nil)
}