//package main
//import "fmt"
//func main()  {
//marvelcharacters:= [7]int{"Hulk", "Iron Man", "Loki", "Vision", "Magneto", "Black Widow", "Wolverine", "Mistique"}
//fmt.Println(oddnumbers)
//fmt.Println(marvelcharacters[4])
//fmt.Println([5])
//fmt.Println(oddnumbers[6])

//}
package main

import "fmt"

func main() {
	blackpanthercharacters := [5]string{"T'Challa", "Nakia", "Okoye", "Shuri", "Killmonger"}
	for index, place := range blackpanthercharacters {
		fmt.Println(index, place)
	}
	fmt.Println("Array printed as a whole", blackpanthercharacters)
}
