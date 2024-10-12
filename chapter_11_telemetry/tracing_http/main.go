package main

/*

go run .

creats trace.out
Check that file using the following:
go tool trace trace.out
*/

import (
	"fmt"
	"net/http"
	"os"
	"runtime/trace"
)

func TraceHandler(inner http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, task := trace.NewTask(r.Context(), r.URL.Path)
		defer task.End()

		trace.Log(ctx, "HTTP Method", r.Method)
		trace.Log(ctx, "URL", r.URL.String())

		inner(w, r.WithContext(ctx))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Tracing!")
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	http.HandleFunc("/", TraceHandler(handler))
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
