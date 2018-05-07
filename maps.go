package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Ibrahim"] = 42
	grades["Ahmed"] = 92
	grades["Khaled"] = 67
	grades["Fahad"] = 83
	grades["Hamad"] = 79
	grades["Ali"] = 91
	fmt.Println(grades)

	Tims_grade := grades["Fahad"]
	fmt.Println(Tims_grade)

	delete(grades, "Fahad")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}
