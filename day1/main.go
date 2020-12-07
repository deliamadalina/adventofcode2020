//package adventofcode
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)



func main() {

	arg := os.Args[1]
	store, err := GetData(arg)
	if err == nil {
		fmt.Print(store)
	} else {
		os.Exit(1)
	}
	FindNumber(store)


}
func GetData(arg string) (store []int, err error) {
	content, err := ioutil.ReadFile(arg)
	store = []int{}

	if err != nil {
		fmt.Print("ERROR! Can't open file man!")
		os.Exit(1)
	}

	data := strings.Split(string(content), "\n")

	for _, k := range data {
		nr, err := strconv.Atoi(k)
		if err != nil{
			return store, err
		}
		store = append(store, nr)
	}
	return store, err
}

func FindNumber(store []int) {
	for i, firstNum := range store {
		for _, secNum := range store[(i+1):] {
			for _, thirdNum := range store[(i+2):] {
				if (firstNum + secNum + thirdNum) == 2020 {
					fmt.Printf("RESULT: %d, %d, %d sum: %d, product: %d", firstNum, secNum, thirdNum, firstNum+secNum+thirdNum, firstNum*secNum*thirdNum)
				}
			}

		}
	}
}