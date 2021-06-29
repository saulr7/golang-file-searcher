package main

import (
	"fmt"
	s "searcher/services"
)

func main() {

	input, err := s.ReadInputs()

	if err != nil {
		fmt.Println(err)
		return
	}

	s.Search(input.Dir, input.FileName)

	s.ShowMatches()

}
