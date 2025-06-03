package main

import (
	maths "Day1_Test_Code/Maths"
	"fmt"

	"rsc.io/quote"
)



var name2 = "Monty"


func main() {

    var a, b int
    name := "Ayush"

    greeting := fmt.Sprintf("Hi %s \n",name) // Use of Sprintf

    fmt.Print(greeting)
    fmt.Println(quote.Go())

    fmt.Printf("Hello %s (%s)!\n",name,name2)

    fmt.Printf("Enter three numbers: ")
	fmt.Scanf("%d %d", &a, &b)

     addResult := maths.Add(a, b)
     subResult := maths.Sub(a, b)
     divResult := maths.Div(a, b)
     modResult := maths.Mod(a, b)
     mulResult := maths.Mul(a, b)


     fmt.Printf("Result:\n Add: %d \n Sub: %d \n Div: %d \n Mod: %d \n Mul: %d", addResult,subResult,divResult,modResult,mulResult)
   
}


