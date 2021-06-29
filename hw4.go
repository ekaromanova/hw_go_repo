package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	counter := map[int64]uint64{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		counter[num]++
	}

	//for number, count := range counter {
		//fmt.Println("number:", number, "count:", count)

		for _, el := range counter {
			fmt.Println(el)
	}
}


