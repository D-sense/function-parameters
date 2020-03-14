package main

import (
	"fmt"
	"github.com/d-sense/function-paramters/functions"
)

func main(){
	// Primitive/Basic Types  ----- STARTS
	fmt.Println("INT")
	age := 30
	fmt.Println("Before function call: ", age)              // 30
	fmt.Println("Function call:", functions.ModifyInt(age))           // 35
	fmt.Println("After function call: ", age)               // 30

	fmt.Println(".........................")

	fmt.Println("FLOAT")
	cash := 10.50
	fmt.Println("Before function call: ", cash)             // 10.5
	fmt.Println("Function call:", functions.ModifyFloat(cash))        // 15.5
	fmt.Println("After function call: ", cash)              // 10.5

	fmt.Println(".........................")

	fmt.Println("BOOLEAN")
	old := false
	fmt.Println("Before function call: ", old)              // false
	fmt.Println("Function call:",  functions.ModifyBool(old))         // true
	fmt.Println("After function call: ", old)               // false

	fmt.Println(".........................")

	fmt.Println("STRING")
	message := "Go"
	fmt.Println("Before function call: ", message)         // Go
	fmt.Println("Function call:", functions.ModifyString(message))   // My favourite language is: Go
	fmt.Println("After function call: ", message)          // Go

	fmt.Println(".........................")

	fmt.Println("ARRAY")
	country := [3]string{"nigeria", "egypt", "sweden"}
	fmt.Println("Before function call: ", country)         // [nigeria egypt sweden]
	fmt.Println("Function call:", functions.ModifyArray(country))    // [nigeria egypt germany]
	fmt.Println("After function call: ", country)          // [nigeria egypt sweden]

	fmt.Println(".........................")

	fmt.Println("STRUCT")
	myProfile := functions.Profile{
		Age:          15,
		Name:         "Adeshina",
		Salary:       300,
		TechInterest: false,
	}
	fmt.Println("Before function call: ", myProfile)         // Go
	fmt.Println("Function call:", functions.ModifyStruct(myProfile))   // My favourite language is: Go
	fmt.Println("After function call: ", myProfile)          // Go

	// Primitive/Basic Types ----- ENDS

	fmt.Println(".........................")

	// Concrete Types ----- STARTS
	fmt.Println("SLICE")
	coffeeBox := []string{"egyptian_coffee", "kenyan_coffee", "brazilian_coffee"}
	fmt.Println("Before function call: ", coffeeBox)         // [egyptian_coffee kenyan_coffee brazilian_coffee]
	fmt.Println("Function call:", functions.ModifySlice(coffeeBox))    // [egyptian_coffee turkish_coffee brazilian_coffee]
	fmt.Println("After function call: ", coffeeBox)          // [egyptian_coffee turkish_coffee brazilian_coffee]


	fmt.Println(".........................")

	// Concrete Types ----- STARTS
	fmt.Println("MAP")
	expenses := make(map[string]int, 0)
	expenses["transport"] = 30
	expenses["food"] = 300
	expenses["rent"] = 100
	fmt.Println("Before function call: ", expenses)         //  map[food:300 rent:100 transport:30]
	fmt.Println("Function call:", functions.ModifyMap(expenses))      //  map[food:4500 rent:100 transport:30]
	fmt.Println("After function call: ", expenses)          //  map[food:4500 rent:100 transport:30]

	fmt.Println(".........................")

	fmt.Println("POINTER")
	m := "Go"
	fmt.Println("Before function call: ", m, age, cash, old, country, myProfile)                              //
	functions.ModifyBasicTypes(&m, &age, &cash, &old, &country, &myProfile) //
	fmt.Println("After function call: ", m, age, cash, old, country, myProfile)                               //

	// Composite Types ----- ENDS


}






