package greet

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error){ // a function which returns a 

		if name=="" {
			return "",errors.New("no name")
		}

		msg:=fmt.Sprintf("Hi %s",name)
		return msg,nil
}