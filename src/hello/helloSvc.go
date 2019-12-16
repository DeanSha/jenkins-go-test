package main
import ("fmt"
	"log"
	"net/http"
	"time"
)

func testRootEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test End Point was hit at " + time.Now().String())
	log.Print("Request for Test Eng Point was served at " + time.Now().String())
}

func handleRequests() {
	http.HandleFunc("/", testRootEndPoint)
	log.Fatal(http.ListenAndServe(":8087", nil)) 
}

func main() {
	handleRequests()
}
