package main

/*
curl -X POST -H "Content-Type: application/json" -d '{"key":"foo","value":"bar123"}' -i http://localhost:8080/set
curl –i "http://localhost:8080/get?key=foo"
*/
import (
	"fmt"
	"net/http"

	"example.com/spewg-cache/example"
	"example.com/spewg-cache/spewg"
)

func main() {
	example.ExamplePrinter()
	cs := spewg.NewCacheServer()

	http.HandleFunc("/set", cs.SetHandler)
	http.HandleFunc("/get", cs.GetHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
