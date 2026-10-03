package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func validateAge(s string) (int, error) {
	// implement
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse: %w", err)
	}
	if v < 0 {
		return v, errors.New("negative")
	}
	return v, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	age, err := validateAge(sc.Text())
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		fmt.Printf("age: %d\n", age)
	}
}
