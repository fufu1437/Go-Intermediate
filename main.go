package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	total := 0

	const parts = 4
	// 计算每段的大小，向上取整，保证覆盖所有元素
	chunkSize := (n + parts - 1) / parts

	for i := 0; i < parts; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if start >= n {
			break
		}
		if end > n {
			end = n
		}

		wg.Add(1)
		go func(lo, hi int) {
			defer wg.Done()
			sum := 0
			for _, v := range nums[lo:hi] {
				sum += v
			}
			mu.Lock()
			total += sum
			mu.Unlock()
		}(start, end)
	}

	wg.Wait()
	fmt.Println(total)
}
