package main

import (
	"fmt"
	"os"
)

func main() {
	const tag = "file_"
	var name, surname, dateOfBirth, profession string
	enterData("Enter your name: ", &name)
	enterData("Enter your surname: ", &surname)
	enterData("Enter your date of birth: ", &dateOfBirth)
	enterData("Enter your profession: ", &profession)
	file1, _ := os.Create(tag + name + surname)
	file1.WriteString("Name: " + name + "\nSurname: " + surname + "\nDate of birth: " + dateOfBirth + "\nProfession" + profession)
}

func enterData(text, value interface{}) {
	fmt.Println(text)
	fmt.Scan(value)
}
