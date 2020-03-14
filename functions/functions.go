package functions


type Profile struct {
	Age int
	Name string
	Salary float64
	TechInterest bool
}

// Functions for manipulating Primitive/Basic Types
func ModifyInt(n int) int {
	return n + 5
}

func ModifyFloat(n float64) float64 {
	return n + 5.0
}

func ModifyBool(n bool) bool {
	return !n
}

func ModifyString(n string) string {
	return "My favourite language is: " + n
}

func ModifyStruct(p Profile) Profile {
	p.Age = 85
	p.Name = "Balqees"
	p.Salary = 500.45
	p.TechInterest = true
	return p
}

func ModifyArray(coffee [3]string) [3]string {
	coffee[2] = "germany"
	return coffee
}



// Functions for manipulating Composite Types
func ModifySlice(coffee []string) []string {
	coffee[1] = "turkish_coffee"
	return coffee
}

func ModifyMap(expenses map[string]int) map[string]int {
	expenses["food"] = 4500
	return expenses
}

func ModifyBasicTypes(name *string, age *int, cash *float64, techInterest *bool, country *[3]string, p *Profile) (*string, *int, *float64, *bool, *[3]string, *Profile)  {
	newName := "Golang"
	newAge := 90
	newCash := 50.45
	newTechInterest := true
	newCountries := [3]string{"nigerian", "colombian", "sudanese"}
	newProfile := Profile {
		Age: 50,
		Name: "Hassan",
		Salary: 45.45,
		TechInterest: false,
	}

	name = &newName
	age = &newAge
	cash = &newCash
	techInterest = &newTechInterest
	country = &newCountries
	p = &newProfile

	return name, age, cash, techInterest, country, p
}


