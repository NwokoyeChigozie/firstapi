package views

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Contact(response http.ResponseWriter, request *http.Request) {
	// var contact models.Contact
	// response.Header().Set("Content-Type", "application/json")
	// _ = json.NewDecoder(request.Body).Decode(&contact)
	// json.NewEncoder(response).Encode(request)
	frmname := request.PostFormValue("name")
	json.NewEncoder(response).Encode(frmname)
	fmt.Println("<div class'alert alert sucess text-center'> " + frmname + "</div>")
}
