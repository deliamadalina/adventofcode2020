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

	err, totalPart1, totalPart2 := Resolve(data)
	if err != nil{
		fmt.Print("Failed to something")
		os.Exit(1)
	}else {
		fmt.Printf("Total password part 01: %d\n", totalPart1)
		fmt.Printf("Total password part 02: %d\n", totalPart2)
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

func Resolve(data []string) (err error, totalPart1 int, totalPart2 int) {
	count := 0
	totalPart1 = 0
	totalPart2 = 0
	for _, k := range data {

		elem := strings.Split(k, " ")
		pos1, _ := strconv.Atoi(strings.Split(elem[0], "-")[0])
		pos2, _ := strconv.Atoi(strings.Split(elem[0], "-")[1])
		letter := string(elem[1][0])
		password := elem[2]

		//fmt.Printf("check password: %s\n", password)
		for _, l := range strings.Split(password, "") {
			result := letter == l
			if result ==  true {
				count += 1
			}
		}
		if count <= pos2 && count >= pos1 {
			totalPart1 += 1
		}

		if string(password[pos1 -1]) == letter  && string(password[pos2 -1]) != letter {
			totalPart2 += 1
		} else if string(password[pos1 -1]) != letter  && string(password[pos2 -1]) == letter {
			totalPart2 +=1
		}


		count = 0
	}
	return err, totalPart1, totalPart2
}

