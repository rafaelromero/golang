package main

import "fmt"

func main(){



	s:=make([]string, 4)
	s[0] = "first string"
	s[1] = "second string"
	s[2] = "3"
	s[3] = "4"

	fmt.Println("array " , s)
	slice1 := s[1:3]
	fmt.Println(slice1)


	i:=make([]int, 2)
	i[0] = 5200
	fmt.Println("int arry ", i)
	

}
