package adventutils

import (
	"io/ioutil"
	s "strings"
	"strconv"
)

func ParseInput (path string) ([]int) {
	content, _ := ioutil.ReadFile(path)
	invoice := s.Split(string(content), "\n")
	expenses := make([]int, len(invoice))
	for index, value := range invoice {
		expenses[index], _ = strconv.Atoi(value)
	}

	return expenses
}