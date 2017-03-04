package main
// Unpredictable server: sometimes slow and sometimes quick
import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)
func main() {
	http.HandleFunc("/", UnpredictableServer)
	http.ListenAndServe(":8082", nil)
}
// sometimes quick, sometimes slow server
func UnpredictableServer(w http.ResponseWriter, req *http.Request) {
	headOrTails := rand.Intn(2)

	if headOrTails == 0 {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(w, "I am S3 in slow mode %v", headOrTails)
		fmt.Printf("I am in S3 slow mode %v\n", headOrTails)
		return
	}

	fmt.Fprintf(w, "I am S3 in quick mode %v", headOrTails)
	fmt.Printf("I am S3 in quick mode %v\n", headOrTails)
	return
}
