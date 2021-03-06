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
	http.ListenAndServe(":8080", nil)
}
// sometimes quick, sometimes slow server
func UnpredictableServer(w http.ResponseWriter, req *http.Request) {
	headOrTails := rand.Intn(2)

	if headOrTails == 0 {
		time.Sleep(4 * time.Second)
		fmt.Fprintf(w, "I am S1 in slow mode %v", headOrTails)
		fmt.Printf("I am in S1 slow mode %v\n", headOrTails)
		return
	}

	fmt.Fprintf(w, "I am S1 in quick mode %v", headOrTails)
	fmt.Printf("I am in S1 quick mode %v\n", headOrTails)
	return
}
