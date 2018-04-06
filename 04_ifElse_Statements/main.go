package main

import "fmt"

func main() {
	
	name := "Frank"
	age := 17

	if name == "Frank" || age == 17 {
		
		if name == "Frank" {
			if age == 17 {

				fmt.Println("Access Granted")
				
			} else {
				fmt.Println("Wrong Age")
			}
		} else { 
			fmt.Println("Wrong Name");
		}
		
	} else {

		fmt.Println("Access Denied")

	}

}