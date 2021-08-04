package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("****************************")
	fmt.Println(" welcome to simon says! ")
	fmt.Println("****************************")
	reader := bufio.NewReader(os.Stdin)
	
	for x := 1; x <= 5; x++ {
		fmt.Println("Enter the statement you want me to print : ")
		
		input, error := reader.ReadString('\n')

		if error != nil {
			fmt.Println("there is an issue reading the input", error)
			return 
		}
		inputtolower := strings.ToLower(input)
		checksimon :="simon says"
		switch  {
		case strings.Contains(inputtolower, checksimon):
			command:=strings.SplitN(inputtolower,"simon says", -1)
			fmt.Println("Round: ",x)
			fmt.Println(command)
		default: 
			//fmt.Println("Round: ",x)
			fmt.Println("Simon says not detected...try again =( =p")
			x--
			fmt.Println("Round: ",x)
		}
		

	}
	fmt.Println("Yayy...you finished and congratulations...hope you had fun =) =P")
} 