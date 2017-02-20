package main
import "fmt"

func main(){



	array := make([]string, 3)

	array[0] = "go"
	array[1] = "is"
	array[2] = "cool"

	for _, ob := range array{
		
		fmt.Println(ob)
	} 




}
