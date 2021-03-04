// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v8"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// To start tracing Redis, simply create a new client using the library and continue
// using as you normally would.
func main() {
	tracer.Start(
		tracer.WithService("dumb-service-name"),
		tracer.WithEnv("nicole.danuwidjaja"),
	)
	defer tracer.Stop()

	ctx := context.Background()
	// create a new Client
	opts := &redis.Options{Addr: "127.0.0.1", Password: "", DB: 0}
	c := redistrace.NewClient(opts)
	// optionally, create a new root span
	root, ctx := tracer.StartSpanFromContext(context.Background(), "parent.request",
		tracer.SpanType(ext.SpanTypeRedis),
		tracer.ServiceName("web"),
		tracer.ResourceName("/home"),
	)

	// commit further commands, which will inherit from the parent in the context.
	c.Set(ctx, "test_key", "test_value", 0)

	// set the context on the client
	c = c.WithContext(ctx)

	c.Set(ctx, "food", "cheese", 0)
	root.Finish()
}

// You can also trace Redis Pipelines. Simply use as usual and the traces will be
// automatically picked up by the underlying implementation.
func Example_pipeliner() {
	ctx := context.Background()
	// create a client
	opts := &redis.Options{Addr: "127.0.0.1", Password: "", DB: 0}
	c := redistrace.NewClient(opts, redistrace.WithServiceName("my-redis-service"))

	// open the pipeline
	pipe := c.Pipeline()

	// submit some commands
	pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	// execute with trace
	pipe.Exec(ctx)
}
