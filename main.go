package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}

	const workers = 4
	partial := make([]int, workers)

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		start := w * n / workers
		end := (w + 1) * n / workers
		wg.Add(1)
		go func(idx, lo, hi int) {
			defer wg.Done()
			sum := 0
			for i := lo; i < hi; i++ {
				sum += nums[i]
			}
			partial[idx] = sum
		}(w, start, end)
	}
	wg.Wait()

	total := 0
	for _, v := range partial {
		total += v
	}
	fmt.Println(total)
}
