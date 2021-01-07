package main

import (
	"context"
	"log"
	"strconv"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const serviceName = "dd-search"
const agentAddr = "127.0.0.1:8126"
const debug = true
const depth = 4
const target = 5

var items = []int{1, 1, 2, 3, 5, 8, 13, 21, 34}

func search(ctx context.Context, n int) {
	s, ctx := tracer.StartSpanFromContext(ctx, "Found number:"+strconv.Itoa(n)+"")
	defer s.Finish()

	index := binarySearch(items, n)
	log.Println("num found at index:", index)
}

func binarySearch(list []int, t int) int {
	l := 0
	r := len(list) - 1
	if l <= r {
		m := (r + l) / 2
		if t < list[m] {
			return binarySearch(list[:m], t)
		} else if t > list[m] {
			return binarySearch(list[m+1:], t)
		} else {
			return 1
		}
	}

	return -1
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
	search(ctx, target)
	s.Finish()
}
