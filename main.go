package main

import (
	"fmt"
	"regexp"
)

func center(s string, w int) string {
	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%*s", (w+len(s))/2, s))
}


func main()  {
	fmt.Println(">>>>>>>", center("QuantumForge", 30), "<<<<<<<")
	fmt.Println()
	fmt.Println("Building tools for Post-Quantum Cryptography migration")
	fmt.Println()
	
	for{
		
		var Name string
		fmt.Println()
		fmt.Print("Please enter your name: ")
		fmt.Scanln(&Name)
		fmt.Println()
		
		if len(Name) == 0 {
			fmt.Println("✗ Error: Name cannot be empty")
			continue
		}
	
		var valid_name = regexp.MustCompile(`^[a-zA-Z]+$`)
		if !valid_name.MatchString(Name) {
			fmt.Printf("✗ Invalid Name: %v are not allowed.\n", Name)
			continue
		}
	
		fmt.Printf("Hello %s Welcome to QuantumForge\n", Name)
		fmt.Println("Let's build quantum-safe tools together.")
		fmt.Println()
		fmt.Println("Project 1 completed. Ready for the next step.")
		break
	}
}