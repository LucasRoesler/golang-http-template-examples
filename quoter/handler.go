package function

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)

		input = body
	}

	logrus.Debug(fmt.Sprintf("Handle; %q", input))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Body: %q", string(input))))
}
