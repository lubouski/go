package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

/* We could just return a copy of str 

func changeValue2(str string) string {
	str = "changed!"
	return str
} */

func main() {
	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)
}

/* in main func use -> changeValue2(toChange)

