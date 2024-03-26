package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/VitoNaychev/simple-chat/gateway/handlers"
)

func TestEchoHandler(t *testing.T) {
	t.Run("echos request", func(t *testing.T) {
		want := handlers.EchoMessage{"Hello from Gateway!"}

		request, _ := http.NewRequest(http.MethodGet, "/echo", nil)
		response := httptest.NewRecorder()

		handlers.EchoHandler(response, request)

		var got handlers.EchoMessage
		json.NewDecoder(response.Body).Decode(&got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
