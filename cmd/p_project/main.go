package main

import (
	"net/http"

	"github.com/Czirko/Go_p_api/pkg/handlers"
)

func main() {
	runAPI()

}
func runAPI() {
	http.HandleFunc("/test", handlers.TestHandle)
	http.ListenAndServe(":8081", nil)
}
