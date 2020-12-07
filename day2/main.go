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

	data := GetData(arg)

	err, totalPartOne := PartOne(data)
	if err != nil{
		fmt.Print("Failed to something")
		os.Exit(1)
	}else {
		fmt.Printf("Total password part 01: %d\n", totalPartOne)
	}

	err, totalPart2 := PartTwo(data)
	if err != nil{
		fmt.Print("Failed to something")
		os.Exit(1)
	}else {
		fmt.Printf("Total password part 02: %d", totalPart2)
	}

}
func GetData(arg string) ( data []string) {
	content, err := ioutil.ReadFile(arg)

	if err != nil {
		fmt.Print("ERROR! Can't open file man!")
		os.Exit(1)
	}
	data = strings.Split(string(content), "\n")

	return data
}

func PartOne(data []string) (err error, total int) {
	count := 0
	total = 0
	for _, k := range data {

		elem := strings.Split(k, " ")
		min, _ := strconv.Atoi(strings.Split(elem[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(elem[0], "-")[1])
		letter := string(elem[1][0])
		password := elem[2]

		//fmt.Printf("check password: %s\n", password)
		for _, l := range strings.Split(password, "") {
			result := letter == l
			if result ==  true {
				count += 1
			}
		}
		//fmt.Printf("count: %d, min: %d, max: %d, letter: %s\n", count, min, max, letter)
		if count <= max && count >= min {
			total += 1
		}
		count = 0
	}
	return err, total
}

func PartTwo(data []string) (err error, total int){
	total = 0
	for _, k := range data {

		elem := strings.Split(k, " ")
		pos1, _ := strconv.Atoi(strings.Split(elem[0], "-")[0])
		pos2, _ := strconv.Atoi(strings.Split(elem[0], "-")[1])
		letter := string(elem[1][0])
		password := elem[2]

		//fmt.Printf("check password: %s\n", password)

		if string(password[pos1 -1]) == letter  && string(password[pos2 -1]) != letter {
			total += 1
		} else if string(password[pos1 -1]) != letter  && string(password[pos2 -1]) == letter {
			total +=1
		}

		//fmt.Printf("total: %d, pos1: %d, pos2: %d, letter: %s, password: %s\n", total, pos1, pos2, letter, password)

	}
	return err, total
}

