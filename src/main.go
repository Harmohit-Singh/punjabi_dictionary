package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	open_page_number(201)

}

func open_page_number(page_number int) {
	filename := "page" + strconv.Itoa(page_number) + ".txt"
	data, err := os.ReadFile(os.Getenv("HOME") + "/Desktop/punjabi_dictionary/sample_data/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
