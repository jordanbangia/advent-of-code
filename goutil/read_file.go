package goutil

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, nil
}
