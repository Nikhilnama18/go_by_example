package main

import "fmt"

func main() {
	fmt.Println("Go" + "lang")

	fmt.Println("What " + "a won" + "derful day")

	fmt.Print(1234)

	fmt.Println(111.253)

	name := "Nick"
	fmt.Printf("My name is %s \n", name)
	/**
	* For more verbs like %s take a look at https://pkg.go.dev/fmt
	* %v can be used in most cases since it uses default format of the value
	 */
	petName := "Mac"
	fmt.Printf("My pet name is %v \n", petName)

	age := 4
	fmt.Printf("%v is %v old \n", petName, age)

	fmt.Println(true && false)
}
