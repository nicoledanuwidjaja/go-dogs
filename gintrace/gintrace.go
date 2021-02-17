package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// Service A
func LivenessGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := tracer.StartSpanFromContext(c.Request.Context(), c.Request.RequestURI)
		defer span.Finish()

		ctx := c.Request.Context()

		// I tried some variations, such as using http.Get("url") instead,
		// or removing "StartSpanFromContext" since the
		// gintrace middleware already has it (https://github.com/DataDog/dd-trace-go/blob/v1.27.1/contrib/gin-gonic/gin/gintrace.go#L23),
		// but it just wouldn't work right.
		// ND: this duplicates the request as it's already set in Middleware.
		req, err := http.NewRequestWithContext(ctx, "GET", "local:8085/path/xyz", nil)
		err = tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(req.Header))

		// These tags don't seem to be set? method and status_code
		span.SetTag("http.method", req.Method)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "request error"})
			span.SetTag("http.status_code", http.StatusBadRequest)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		content, _ := ioutil.ReadAll(resp.Body)
		span.SetTag("http.status_code", http.StatusOK)
		c.JSON(http.StatusOK, string(content))
	}
}

func main() {
	tracer.Start(
		tracer.WithServiceName("geo-service"),      // external services
		tracer.WithGlobalTag("env", "development"), // development
	)
	defer tracer.Stop()

	r := gin.Default()

	// service A's main.go
	r.Use(gintrace.Middleware("geo-service"))
	r.GET("/liveness", LivenessGet())
}
