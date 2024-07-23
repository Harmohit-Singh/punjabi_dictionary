package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := take_unicode_input()
	fmt.Println(input)
	page_txt := open_page_number(201)
	fmt.Println(string(page_txt))
}

func open_page_number(page_number int) string {
	filename := "page" + strconv.Itoa(page_number) + ".txt"
	data, err := os.ReadFile(os.Getenv("HOME") + "/Desktop/punjabi_dictionary/sample_data/" + filename)
	if err != nil {
		fmt.Println(err)
		return "data not found"
	}
	return string(data)
}

func take_unicode_input() string {
	var input string
	fmt.Print("Enter word to search: ")
	fmt.Scan(&input)
	return input
}
