package main

import (
	"fmt"
	"os"
)

func main() {
	const tag = "file_"                                      // Initialization `tag` as "file_" for creating file with same prefix.
	
	var name, surname, dateOfBirth, profession string
	
	enterData("Enter your name: ", &name) 			 // Take data from user
	enterData("Enter your surname: ", &surname)              // ^
	enterData("Enter your date of birth: ", &dateOfBirth)    // ^
	enterData("Enter your profession: ", &profession)        // ^
	
	file1, _ := os.Create(tag + name + surname)              // Creating file named by tag, name & surname of user.
	file1.WriteString("Name: " + name + "\nSurname: " + surname + "\nDate of birth: " + dateOfBirth + "\nProfession" + profession) // Adding to file data of user.
}

func enterData(text, value interface{}) { 			// This function helps program to accept data of user.
	fmt.Println(text)
	fmt.Scan(value)
}
