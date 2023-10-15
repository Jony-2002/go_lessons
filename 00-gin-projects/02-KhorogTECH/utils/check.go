package utils

import (
	"gin/KhorogTECH/types"
	"log"
	"reflect"
)

func CheckData(obj types.Registration) types.IsOKType {
	values := reflect.ValueOf(obj)
	typeOfS := values.Type()
	log.Printf("The data which came from client: %v", obj)

	var isOK bool = true
	var Message string

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Interface() == "" || values.Field(i).Interface() == 0 {
			isOK = false
			Message = "Field " + typeOfS.Field(i).Name + " is required"
			log.Printf("Field %v is required\n", typeOfS.Field(i).Name)
			break
		}
	}

	if isOK {
		var res types.IsOKType = types.IsOKType{
			Result: true,
			Message: Message,
		}
		
		return res 
	} else {
		var res types.IsOKType = types.IsOKType{
			Result: false,
			Message: Message,
		}
		
		return res
	}
}
