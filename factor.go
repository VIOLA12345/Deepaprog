package main
import ( 
	"fmt"
)
	
func fact(inputNumber int){
	var divsisor int = int(inputNumber)
	for i:= 1; i <=divsisor; i++ {
		if inputNumber%i == 0 {
			fmt.Println(i, "is a factor of", inputNumber)
		}
	}
}
func main()  {
	
	fact(30000000000)
}