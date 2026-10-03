package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() (v int, ok bool) {
	defer func() {
		if err := recover(); err != nil {
			ok = false
			v = 0
		} else {
			s.items = s.items[:len(s.items)-1]
		}
	}()
	return s.items[len(s.items)-1], true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	parts := strings.Fields(sc.Text())
	var s Stack
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		_ = n
		s.Push(n)
	}
	for {
		x, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println(x)
	}
	_ = fmt.Print
}
