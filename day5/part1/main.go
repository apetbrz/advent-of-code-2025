package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("sequence.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lows := make([]int, 0)
	highs := make([]int, 0)

	fresh_count := 0

	lines := bufio.NewScanner(file)
	for lines.Scan() {
		if lines.Text() == "" {
			break
		}

		line_range := strings.Split(lines.Text(), "-")

		line_lo, err := strconv.Atoi(line_range[0])
		if err != nil {
			panic(err)
		}

		line_hi, err := strconv.Atoi(line_range[1])
		if err != nil {
			panic(err)
		}

		lows = append(lows, line_lo)
		highs = append(highs, line_hi)
	}

	for lines.Scan() {
		id, err := strconv.Atoi(lines.Text())
		if err != nil {
			panic(err)
		}

		// iterate through lows and highs and check if id is in range
		for i := 0; i < len(lows); i++ {
			if id >= lows[i] && id <= highs[i] {
				fresh_count++
				break
			}
		}
	}

	if err := lines.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", fresh_count)
}
