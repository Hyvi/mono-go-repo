package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	sayHello := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!!")
	}
	req := httptest.NewRequest("GET", "http://example.com/hello", nil)
	w := httptest.NewRecorder()
	sayHello(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	require.EqualValues(t, "Hello!!", string(body))
}
