package main

import (
	"context"
	"log"
	"strconv"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const serviceName = "dd-fib"
const agentAddr = "127.0.0.1:8126"
const debug = true
const depth = 4

func fib(ctx context.Context, n int) int {
	s, ctx := tracer.StartSpanFromContext(ctx, "fib("+strconv.Itoa(n)+")")
	defer s.Finish()
	if n <= 2 {
		return 1
	}
	return fib(ctx, n-1) + fib(ctx, n-2)
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

	// returns new span as parent
	s, ctx := tracer.StartSpanFromContext(context.Background(), serviceName+".opname", tracer.ResourceName(serviceName+".resname"))
	fibn := fib(ctx, depth)
	s.Finish()
	log.Printf("fib(%d) = %d", depth, fibn)
}
