package views

import (
	"encoding/json"
	"net/http"
)

func Contact(response http.ResponseWriter, request *http.Request) {
	// var contact models.Contact
	// response.Header().Set("Content-Type", "application/json")
	// _ = json.NewDecoder(request.Body).Decode(&contact)
	// json.NewEncoder(response).Encode(request)
	frmname := request.PostFormValue("name")
	json.NewEncoder(response).Encode(frmname)
	// fmt.Println(frmCategory)
}
