package main
import "fmt"


func addem(one, two, three int) (int, int){

	a := one + two + three
	b := 200

	return a,b

}



func main(){

	a, b := addem(1, 2, 4)

	fmt.Println(a)
	fmt.Println(b)

}
