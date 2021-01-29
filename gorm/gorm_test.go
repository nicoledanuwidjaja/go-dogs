package gorm

import (
	"log"
	"testing"

	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/internal/sqltest"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	gormtrace "gopkg.in/jinzhu/gorm.v1"
)

// TestPostgreSQL tests Gorm functionality.
func TestPostgreSQL(t *testing.T) {
	// Register augments the provided driver with tracing, enabling it to be loaded by gormtrace.Open.
	sqltrace.Register("postgres", &pq.Driver{}, sqltrace.WithServiceName("my-service"))

	// Open the registered driver, allowing all uses of the returned *gorm.DB to be traced.
	db, err := gormtrace.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	testConfig := &sqltest.Config{
		DB:         db.DB(),
		DriverName: "mysql",
		TableName:  "testgorm",
		ExpectName: "mysql.query",
		ExpectTags: map[string]interface{}{
			ext.ServiceName: "mysql-test",
			ext.SpanType:    ext.SpanTypeSQL,
			ext.TargetHost:  "127.0.0.1",
			ext.TargetPort:  "3306",
			"db.user":       "test",
			"db.name":       "test",
		},
	}
	sqltest.RunAll(t, testConfig)
}
