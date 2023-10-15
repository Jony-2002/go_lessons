package controllers

import (
	"encoding/json"
	"fmt"
	"khorogtech/data"
	"khorogtech/types"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n127.0.0.1:8888/getUsers - GET Request")
	
	res := types.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data.Users,
	}
	
	response, _ := json.Marshal(res)
	
	fmt.Printf("%+v\n", data.Users)
	w.WriteHeader(http.StatusOK) 				// StatusOK - 200
	w.Write(response)
}