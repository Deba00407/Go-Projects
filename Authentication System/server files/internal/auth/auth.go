package auth

import (
	"fmt"
	"io"
	"net/http"
)

func MakePostRequestToAPIEndpoint(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST request allowed on this route", http.StatusBadRequest)
	}

	response, err := io.ReadAll(req.Body)
	if err == nil {
		data := string(response)
		res.Write([]byte(data))
		fmt.Println("The data received was: ", data)
	}else{
		fmt.Println("Error receiving data: ", err)
	}

}
