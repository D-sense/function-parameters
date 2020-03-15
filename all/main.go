package main

import (
	"fmt"
	"strings"
)

type profile struct {
	Age int
	Name string
	Salary float64
	TechInterest bool
}

type detail struct {
	Age int
	Name string
	Cash float64
	TechInterest bool
	Countries [3]string
	MyProfile profile
}

func main(){
	// Primitive/Basic Types  ----- STARTS
	fmt.Println("INT")
	age := 30
	fmt.Println("Before function call: ", age)              // 30
	fmt.Println("Function call:", modifyInt(age))           // 35
	fmt.Println("After function call: ", age)               // 30

	fmt.Println(".........................")

	fmt.Println("FLOAT")
	cash := 10.50
	fmt.Println("Before function call: ", cash)             // 10.5
	fmt.Println("Function call:", modifyFloat(cash))        // 15.5
	fmt.Println("After function call: ", cash)              // 10.5

	fmt.Println(".........................")

	fmt.Println("BOOLEAN")
	old := false
	fmt.Println("Before function call: ", old)              // false
	fmt.Println("Function call:",  modifyBool(old))         // true
	fmt.Println("After function call: ", old)               // false

	fmt.Println(".........................")

	fmt.Println("STRING")
	message := "Go"
	fmt.Println("Before function call: ", message)         // Go
	fmt.Println("Function call:", modifyString(message))   // My favourite language is: Go
	fmt.Println("After function call: ", message)          // Go

	fmt.Println(".........................")

	fmt.Println("ARRAY")
	country := [3]string{"nigeria", "egypt", "sweden"}
	fmt.Println("Before function call: ", country)         // [nigeria egypt sweden]
	fmt.Println("Function call:", modifyArray(country))    // [nigeria egypt germany]
	fmt.Println("After function call: ", country)          // [nigeria egypt sweden]

	fmt.Println(".........................")

	fmt.Println("STRUCT")
	myProfile := profile{
		Age:          15,
		Name:         "Adeshina",
		Salary:       300,
		TechInterest: false,
	}
	fmt.Println("Before function call: ", myProfile)         // Go
	fmt.Println("Function call:", modifyStruct(myProfile))   // My favourite language is: Go
	fmt.Println("After function call: ", myProfile)          // Go

	// Primitive/Basic Types ----- ENDS



	fmt.Println(".........................")

	// Concrete Types ----- STARTS
	fmt.Println("SLICE")
	coffeeBox := []string{"egyptian_coffee", "kenyan_coffee", "brazilian_coffee"}
	fmt.Println("Before function call: ", coffeeBox)         // [egyptian_coffee kenyan_coffee brazilian_coffee]
	fmt.Println("Function call:", modifySlice(coffeeBox))    // [egyptian_coffee turkish_coffee brazilian_coffee]
	fmt.Println("After function call: ", coffeeBox)          // [egyptian_coffee turkish_coffee brazilian_coffee]

	fmt.Println(".........................")

	// Concrete Types ----- STARTS
	fmt.Println("MAP")
	expenses := make(map[string]int, 0)
	expenses["transport"] = 30
	expenses["food"] = 300
	expenses["rent"] = 100
	fmt.Println("Before function call: ", expenses)         //  map[food:300 rent:100 transport:30]
	fmt.Println("Function call:", modifyMap(expenses))      //  map[food:4500 rent:100 transport:30]
	fmt.Println("After function call: ", expenses)          //  map[food:4500 rent:100 transport:30]

	fmt.Println(".........................")

	fmt.Println("POINTER")
	myProfile=  profile{
		Age:          0,
		Name:         "",
		Salary:       0,
		TechInterest: false,
	}

	fmt.Println("Before function call: ", message, age, cash, old, country, myProfile)     // {0  0 false [nigeria egypt swed] {0  0 false}}
	ModifyBasicTypes(&message, &age, &cash, &old, &country, &myProfile)
	fmt.Println("After function call: ", message, age, cash, old, country, myProfile)      // {90 Golang 50.45 true [nigerian colombian sudanese] {50 Hassan 45.45 false}}

	fmt.Println(".........................")

	fmt.Println("FUNCTION")
	msg := "adeshina"
	anon := func(n string) string {
		return n
	}

	fmt.Printf("%T", anon)
	fmt.Println("")
	fmt.Println("Before function call: ", anon(msg))
	fmt.Println("Function call: ", modifyFunction(anon, msg))
	fmt.Println("After function call: ", anon(msg))

	fmt.Println(".........................")

	fmt.Println("CHANNEL")
	status := make(chan string)
	go modifyChannel(status)
	fmt.Println("After function call: ", <- status)

	// Composite Types ----- ENDS
}

// Functions for manipulating Primitive/Basic Types
func modifyInt(n int) int {
	return n + 5
}

func modifyFloat(n float64) float64 {
	return n + 5.0
}

func modifyBool(n bool) bool {
	return !n
}

func modifyString(n string) string {
	n = "Golang"
	return n
}

func modifyStruct(p profile) profile {
	p.Age = 85
	p.Name = "Balqees"
	p.Salary = 500.45
	p.TechInterest = true
	return p
}

func modifyArray(coffee [3]string) [3]string {
	coffee[2] = "germany"
	return coffee
}


// Functions for manipulating Composite Types
func modifySlice(coffee []string) []string {
	coffee[1] = "turkish_coffee"
	return coffee
}

func modifyMap(expenses map[string]int) map[string]int {
	expenses["food"] = 4500
	return expenses
}

func ModifyBasicTypes(name *string, age *int, cash *float64, techInterest *bool, countries *[3]string, myProfile *profile) {
	*name = "Golang"
	*age = 90
	*cash = 50.45
	*techInterest = !(*techInterest)
	*countries = [3]string{"sudanese", "belgium", "zambia"}
	*myProfile = profile{
		Age:          100,
		Name:         "GOOGLE",
		Salary:       40.54,
		TechInterest: true,
	}
}

func modifyFunction(f func(string) string, name string) string {
	// return f(name)
	f = func (n string) string{
		return strings.ToUpper(n)
	}

	return f(name)
}

func modifyChannel(s chan string) {
	s <- "INJECTING A NEW MESSAGE"
}
