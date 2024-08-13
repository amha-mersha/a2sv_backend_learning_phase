package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateAverage(data ...float64) float64 {
	var total float64
	var count float64
	for _, itm := range data {
		total += itm
		count += 1
	}
	return total / count
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Hello there!")
	fmt.Print("Enter you name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name")
	}
	name = name[:len(name)-1]
	for strings.Trim(name, " ") == "" {
		fmt.Printf("Enter a correct name: ")
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading name")
		}
		name = name[:len(name)-1]
	}

	var subjectCount int
	for {
		var input string
		fmt.Print("Enter the number of courses: ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading number of courses:", err)
			continue
		}
		_, err = fmt.Sscanf(input, "%d", &subjectCount)
		if err != nil {
			fmt.Print("Invalid Value.")
			continue
		}
		break
	}

	subjects := make([]string, subjectCount)
	scores := make([]float64, subjectCount)
	for index := 0; index < subjectCount; index++ {
		var (
			temp float64
			curr string
		)

		for {
			fmt.Printf("Enter the %v subject: ", index+1)
			curr, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error while reading name.")
				continue
			}

			if strings.TrimSpace(curr) == "" {
				fmt.Println("Invalid subject name.")
				continue
			}
			fmt.Println("subject name: ", curr)
			break
		}
		for {
			var input string
			fmt.Printf("Enter the score for the %v subject: ", index+1)
			_, err = fmt.Scanln(&input)
			if err != nil {
				fmt.Println("Error while reading the score of the subject")
				continue
			}
			if strings.TrimSpace(input) == "" {
				fmt.Println("Invalid value.")
				continue
			}

			_, err := fmt.Sscanf(input, "%f", &temp)
			if err != nil {
				fmt.Print("Invalid Value.")
				continue
			}
			break
		}

		subjects[index] = curr
		scores[index] = float64(temp)
	}
	fmt.Printf("Your name is %v.\n", name)
	fmt.Println("Your took the following classes : ")
	for i := range subjects {
		fmt.Printf("%v : %v \n", subjects[i], scores[i])
	}
	fmt.Printf("and your total average is: %v \n", calculateAverage(scores...))
}
