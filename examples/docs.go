package main

import (
	"context"
	"log"
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func spanner(ctx context.Context) {
	s, ctx := tracer.StartSpanFromContext(ctx, "child", tracer.Tag("tester", 1))
	defer s.Finish()
}

func main() {
	tracer.Start(
		tracer.WithService("dumb-service-name"),
		tracer.WithEnv("nicole.danuwidjaja"),
	)
	defer tracer.Stop()
	if err := profiler.Start(
		profiler.WithService("dumb-service-name"),
		profiler.WithEnv("nicole.danuwidjaja"),
	); err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	s, ctx := tracer.StartSpanFromContext(context.Background(), "tester"+".opname", tracer.ResourceName("tester"+".resname"))
	spanner(ctx)
	s.Finish()
	log.Printf("Hello")

	// Create a traced mux router.
	mux := muxtrace.NewRouter()

	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	mux.HandleFunc("/blah", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("boo!"))
	})
	http.ListenAndServe(":8080", mux)
}
