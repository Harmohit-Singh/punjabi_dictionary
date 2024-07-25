package src

import (
	"fmt"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	pages_to_search := []int{201, 202, 203, 204, 205}
	input := take_unicode_input()
	pages_in_order := search(input, pages_to_search)
	to_display := Open_page_number(pages_in_order[0].page_number)
	fmt.Println(to_display)
}

type page_distance struct {
	page_number int
	distance    int
}

func Open_page_number(page_number int) string {
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

func search(word string, pages_to_search []int) []page_distance {
	var pages_in_order []page_distance
	for i := 0; i < len(pages_to_search); i++ {
		page_number := pages_to_search[i]
		txt := Open_page_number(page_number)
		dist := find_min_distance(word, txt)
		pages_in_order = append(pages_in_order, page_distance{page_number, dist})
	}
	// Sort pages_in_order by distance in ascending order
	sort.Slice(pages_in_order, func(i, j int) bool {
		return pages_in_order[i].distance < pages_in_order[j].distance
	})
	return pages_in_order
}

func find_min_distance(word string, txt string) int {
	min_distance := 999
	list_of_words := strings.Fields(txt)

	for _, word_from_list := range list_of_words {
		dist := fuzzy.LevenshteinDistance(word, word_from_list)
		if dist < min_distance {
			min_distance = dist
		}
	}
	return min_distance
}
