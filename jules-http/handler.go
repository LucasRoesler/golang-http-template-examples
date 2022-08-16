package function

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	handler "github.com/openfaas/templates-sdk/go-http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	input := req.Body
	logrus.Debug(fmt.Sprintf("They said: %q", input))

	message := fmt.Sprintf("Say %q one more time!", string(input))

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
