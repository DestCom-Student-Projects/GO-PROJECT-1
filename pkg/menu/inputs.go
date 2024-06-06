package menu

import "fmt"


func Inputs(choices []string) (int) {
	var choice int
	var isCorrect bool = false

	for !isCorrect {
		fmt.Scan(&choice)
		if choice < 1 || choice > len(choices) {
			fmt.Println("Invalid choice, try again")
			continue
		}
		isCorrect = true
	}
	return choice
}