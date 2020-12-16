package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	arg := os.Args[1]
	data := ReadFile(arg)
	err, trees := CoundTrees(data)
	if err == nil {
		fmt.Printf("Trees encountered: %d", trees)
	} else{
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

}

func ReadFile(arg string) (data []string) {
	content, err := ioutil.ReadFile(arg)

	if err != nil {
		fmt.Print("ERROR! Can't open file man!")
		os.Exit(1)
	}
	data = strings.Split(string(content), "\n")
	return data
}

func CoundTrees(data []string) (err error, trees int){
	count := 0
	pozition := 1
	total_lines := len(data)
	trees = 0

	for _, line := range data{
		count += 1
		pozition = pozition + 3

		val := strings.Split(line, "")

		for len(val) <= pozition {
			val = append(val, val...)
		}
		if val[pozition] == "#" {
			trees += 1
		}

		if total_lines != count {
			fmt.Errorf("Failed to count all lines")
		}

	}
	return err, trees

}