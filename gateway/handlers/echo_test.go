package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/VitoNaychev/simple-chat/gateway/handlers"
)

func TestEchoHandler(t *testing.T) {
	t.Run("echos request", func(t *testing.T) {
		want := handlers.EchoMessage{"Hello, World!"}

		body := bytes.NewBuffer([]byte{})
		json.NewEncoder(body).Encode(want)

		request, _ := http.NewRequest(http.MethodPost, "/echo", body)
		response := httptest.NewRecorder()

		handlers.EchoHandler(response, request)

		var got handlers.EchoMessage
		json.NewDecoder(response.Body).Decode(&got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
