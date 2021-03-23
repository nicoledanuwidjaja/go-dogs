package main

import (
	// "context"
	"log"
	"math/rand"
	"time"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"github.com/beefsack/go-rate"
)

func main() {
	tracer.Start(
		tracer.WithRuntimeMetrics(),
		tracer.WithService("test-spans"),
		tracer.WithEnv("nicole.danuwidjaja"),
	)
	defer tracer.Stop()

	rl := rate.New(10, time.Second)
	// ctx := context.Background()
	for {
		rl.Wait()
		startTrace()
	}
}

func startTrace() {
	now := time.Now()
	start := now.Add(-time.Duration(rand.Intn(500)) * time.Millisecond)
	root := tracer.StartSpan("root", tracer.StartTime(start))
	defer root.Finish()

	n := rand.Intn(5)

	for i := 0; i < n; i++ {
		log.Printf("new child: %d", i)
		tracer.StartSpan("child", tracer.StartTime(start), tracer.ChildOf(root.Context())).Finish()
	}
}