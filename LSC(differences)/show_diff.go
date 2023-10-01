package main

import "fmt"

func Show(data []changedCell) {
	lastType := None
	fmt.Println("-------------------------")

	for _, elem := range data {

		if elem.changeType != lastType {

			if lastType != None {
				fmt.Print("\n")
			}

			lastType = elem.changeType

			switch elem.changeType {
			case Added:
				fmt.Print("+| ")
			case Removed:
				fmt.Print("-| ")
			case Equal:
				fmt.Print("=| ")
			default:
				fmt.Print("err !!!")
			}
		}

		fmt.Printf("%c", elem.value)

	}
	fmt.Println("\n-------------------------")
}
