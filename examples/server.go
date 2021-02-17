package main

import (
	"context"
	"log"
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const serviceName = "dd-server"
const agentAddr = "127.0.0.1:8126"
const debug = true

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func startServer(ctx context.Context) {
	s, ctx := tracer.StartSpanFromContext(ctx, "Server starting")
	mux := muxtrace.NewRouter()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":8080", mux)
	s.Finish()
}

func main() {
	log.SetPrefix(serviceName + ":")
	// start DD trace agent
	tracer.Start(
		tracer.WithServiceName(serviceName),
		tracer.WithAgentAddr(agentAddr),
		tracer.WithDebugMode(debug),
	)
	defer tracer.Stop()
	s, ctx := tracer.StartSpanFromContext(context.Background(), serviceName+".opname", tracer.ResourceName(serviceName+".resname"))
	startServer(ctx)
	s.Finish()
	log.Printf("Server ran!")

}
