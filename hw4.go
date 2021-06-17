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





package main

import (
"bufio"
"fmt"
"os"
"strconv"
)

func main() {
	inputNums := [5]int64{}
	scanner := bufio.NewScanner(os.Stdin)

	// записываем введённые числа в массив в обратном порядке
	for i := len(inputNums) - 1; i >= 0; i-- {
		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			inputNums[i] = num
		} else {
			panic("you must input 5 numbers")
		}
	}
	for _, el := range inputNums {
		fmt.Println(el)
	}
}

