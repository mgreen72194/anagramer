package main

import (
	"container/list"
	"fmt"

	"anagramer/internal/helper"
)

func main() {
	fmt.Println("Hello world!")

	//wordList := list.New()
	magicWord := "lemon god"
	testWords := list.New()
	testWords.PushBack("Jim")
	testWords.PushBack("a")
	testWords.PushBack("melon")
	testWords.PushBack("dog")
	testWords.PushBack("Thing")
	magicNumber := helper.StringToInt(magicWord)
	fmt.Println(magicNumber)

}
