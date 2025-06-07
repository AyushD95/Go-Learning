package main

import (
	"Day2/mod2"
	"fmt"
)

func main(){

	fmt.Printf("Welcome! \n")

	msg, err := greet.Greet("Ayush")  //as this greet package return string and eroro we have to assign 2 variable msg for greet msg
									 // and err for error msg 

		if err != nil {
			fmt.Println("Error:", err)
		 
		} else {
			fmt.Println(msg)
		
		}

}