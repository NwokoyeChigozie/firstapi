package views

import (
	"encoding/json"
	"net/http"

	"github.com/gregoflash05/firstapi/models"
)

func Contact(response http.ResponseWriter, request *http.Request) {
	var contact models.Contact
	// response.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(request.Body).Decode(&contact)
	json.NewEncoder(response).Encode(request)
}
