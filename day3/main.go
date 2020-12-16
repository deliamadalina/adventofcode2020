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

func CoundTrees(data []string) (err error, totalSlope int){
	count := 0
	pozitionSlope1 := 1
	pozitionSlope2 := 1
	pozitionSlope3 := 1
	pozitionSlope4 := 1
	pozitionSlope5 := 1
	total_lines := len(data)
	treesSlope1 := 0
	treesSlope2 := 0
	treesSlope3 := 0
	treesSlope4 := 0
	treesSlope5 := 0
	countSlope5 := 1
	tree := "#"

	for _, line := range data{
		count += 1
		val := strings.Split(line, "")

		for len(val) <= pozitionSlope2 || len(val) <= pozitionSlope1 || len(val) <= pozitionSlope3 || len(val) <= pozitionSlope4{
			val = append(val, val...)
		}
		if val[pozitionSlope1-1] == tree {
			treesSlope1 += 1
		}
		if val[pozitionSlope2-1] == tree {
			treesSlope2 += 1
		}
		if val[pozitionSlope3-1] == tree {
			treesSlope3 += 1
		}
		if val[pozitionSlope4-1] == tree {
			treesSlope4 += 1
		}
		if count == countSlope5{
			if val[pozitionSlope5-1] == tree {
				treesSlope5 += 1
			}
			pozitionSlope5 += 1
			countSlope5 += 2
		}

		if total_lines != count {
			fmt.Errorf("Failed to count all lines")
		}

		pozitionSlope1 += 1
		pozitionSlope2 += 3
		pozitionSlope3 += 5
		pozitionSlope4 += 7


	}
	fmt.Printf("%d, %d, %d, %d, %d\n", treesSlope1, treesSlope2, treesSlope3, treesSlope4, treesSlope5)
	return err, treesSlope1*treesSlope2*treesSlope3*treesSlope4*treesSlope5

}