package main

import (
	"context"
	"log"
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	tracer.Start(
		tracer.WithService("service"),
		tracer.WithEnv("env"),
	)
	defer tracer.Stop()
	if err := profiler.Start(
		profiler.WithService("service"),
		profiler.WithEnv("env"),
	); err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	s, _ := tracer.StartSpanFromContext(context.Background(), "tester"+".opname", tracer.ResourceName("tester"+".resname"))
	s.Finish()
	log.Printf("Hello")

	// Create a traced mux router.
	mux := muxtrace.NewRouter()

	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", mux)
}
