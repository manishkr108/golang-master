package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

//Struct with Methods

type Rectangele struct {
	width  float64
	height float64
}

func (r Rectangele) Area() float64 {
	return r.width * r.height
}

func main() {
	res := Rectangele{width: 5, height: 5}
	fmt.Println("Area is: ", res.Area())

	emp := []Student{{Name: "Manish Kumar", Age: 30, Salary: 167000}}
	// Struct to JSON (Marshal)
	//Marshal = Convert Go struct → JSON (send to API, save to file, store in DB)
	jsonData, _ := json.Marshal(emp)

	fmt.Println("JsonData", string(jsonData))
	for _, val := range emp {
		fmt.Println("Name:", val.Name, "Age: ", val.Age, "Salary: ", val.Salary)
	}
	// JSON to Struct (Unmarshal)
	//Unmarshal = Convert JSON → Go struct (read from API response, load from file)
	// Unmarshal → Convert JSON string → Go struct/map/slice

	// Create a variable of type Student (empty struct)
	var s2 []Student
	jsonStr := `{"name":"Rahul","age":25,"Salary":163000}`
	_ = json.Unmarshal([]byte(jsonStr), &s2)
	for i, s := range s2 {
		fmt.Printf("Record %d => Name: %s | Age: %d | salary: %.2f\n",
			i+1, s.Name, s.Age, s.Salary)
	}

}
