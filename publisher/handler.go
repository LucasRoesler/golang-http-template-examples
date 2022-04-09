package function

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	nats "github.com/nats-io/nats.go"
	handler "github.com/openfaas/templates-sdk/go-http"

	"handler/function/pkg/version"
)

var (
	subject        = "nats-test"
	defaultMessage = "Hello World"
)

// Handle a function invocation
func Handle(r handler.Request) (handler.Response, error) {
	var err error

	fmt.Println(version.Version)

	msg := defaultMessage
	if r.Body != nil {
		msg = string(bytes.TrimSpace(r.Body))
	}

	natsURL := nats.DefaultURL
	val, ok := os.LookupEnv("nats_url")
	if ok {
		natsURL = val
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		errMsg := fmt.Sprintf("can not connect to nats: %s", err)
		log.Print(errMsg)

		return handler.Response{
			Body:       []byte(errMsg),
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	defer nc.Close()

	s := r.Header.Get("nats-subject")
	if s != "" {
		subject = s
	}

	log.Printf("Publishing %d bytes to: %q\n", len(msg), subject)

	err = nc.Publish(subject, []byte(msg))
	if err != nil {
		log.Println(err)

		return handler.Response{
			Body:       []byte(fmt.Sprintf("can not publish to nats: %s", err)),
			StatusCode: http.StatusInternalServerError,
		}, err

	}
	return handler.Response{
		Body:       []byte(fmt.Sprintf("Published %d bytes to: %q", len(msg), subject)),
		StatusCode: http.StatusOK,
	}, err
}
