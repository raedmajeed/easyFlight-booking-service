package utils

import (
	"encoding/json"
	"fmt"
	"github.com/raedmajeed/booking-service/pkg/DOM"
)

func ErrorEmail(err error, function, pkg, service string) []byte {
	email := DOM.EmailMessage{
		Email:   "raedam786@gmail.com",
		Subject: fmt.Sprintf("ERROR IN %v SERVICE", service),
		Content: fmt.Sprintf("ERROR IN %v method, %v package, please check it out. error is defined: %v", function, pkg, err),
	}
	marshal, _ := json.Marshal(email)
	return marshal
}
