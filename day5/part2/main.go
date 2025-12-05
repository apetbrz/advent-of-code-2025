package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("sequence.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	range_order := make([]int, 0)
	ranges := make(map[int]int)

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

		//if line_lo already exists in map
		//replace its hi value with the greater of:
		//(the existing one, the new one)
		//avoids having two ranges with equal low bounds
		if ranges[line_lo] == 0 {
			ranges[line_lo] = line_hi
			range_order = append(range_order, line_lo)
		} else {
			ranges[line_lo] = max(line_hi, ranges[line_lo])
		}
	}

	slices.Sort(range_order)

	for i := 1; i < len(range_order); i++ {
		xlo := range_order[i-1]
		xhi := ranges[xlo]
		ylo := range_order[i]
		yhi := ranges[ylo]
		// list is sorted, therefor,
		// we know xlo < ylo
		//
		// check for if xhi >= ylo
		if ylo <= xhi {
			// as then we can "consume" y
			// and join it into x
			ranges[xlo] = max(xhi, yhi)
			delete(ranges, ylo)
			range_order = append(range_order[:i], range_order[i+1:]...)
			// and decrement as to re-check the same spot
			i--
		}
	}

	sum := 0

	for lo, hi := range ranges {
		sum += hi - lo + 1
	}

	fmt.Println(sum)

	if err := lines.Err(); err != nil {
		panic(err)
	}
}
