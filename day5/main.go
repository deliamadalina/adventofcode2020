package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)
var (
	totalChair = 127
	totalRaw = 7
	count = 1
	seatID []int
	arg = os.Args[1]
	data = ReadFile(arg)
	listChairs []int
	listRaws []int
)


func main() {

	seatID := 0
	seatIDs, chairs, rows := GetSeatRecursive(data[len(data)-1])

	for _, v := range seatIDs {
		if seatID < v {
			seatID = v
		}
	}
	fmt.Println("PART I: ", seatID)

	mySeat := FindMySeat(chairs) * 8 + FindMySeat(rows)
	fmt.Printf("PART II: %d", mySeat)

}


func FindMySeat(loc []int)  int{

	rdMap := make(map[int]int)
	result := 0

	sort.Ints(loc)

	for i:=loc[0]; i<= loc[len(loc)-1]; i++ {
		count1 := 0
		for _,c := range loc {
			if i == c {
				count1 += 1
				rdMap[c] = count1
			}
		}
	}

	min := rdMap[len(rdMap)-1]
	for key, val := range rdMap {
		if min > val && key != loc[0] && key != loc[len(loc)-1] {
			min = val
			result = key
		}
	}

	return result

}

func GetSeatRecursive(seat string) ([]int, []int, []int) {
	startChair := 0
	startRaw := 0
	endChair := totalChair
	endRaw   := totalRaw

	letterChair := strings.Split(seat[:7], "")
	letterRaw   := strings.Split(seat[7:], "")

	for _, l := range letterChair {
		if l == "F" {
			endChair = (endChair-startChair)/2 + startChair
		} else if l == "B" {
			startChair = (endChair-startChair)/2 + startChair + totalChair %2
		}
	}

	for _, l := range letterRaw {
		if l == "L" {
			endRaw = (endRaw-startRaw)/2 + startRaw
		} else if l == "R" {
			startRaw = (endRaw-startRaw)/2 + startRaw + totalRaw %2
		}
	}
	seatID = append(seatID, startChair *8 + startRaw)
	listChairs = append(listChairs, startChair)
	listRaws = append(listRaws, startRaw)

	count += 1


	if count <= len(data) {
		GetSeatRecursive(data[len(data)-count])
	}
	return seatID, listChairs, listRaws

}



func ReadFile(arg string) (data []string){
	content, err := ioutil.ReadFile(arg)
	if err != nil {
		fmt.Print("ERROR! Can't open file man!")
		os.Exit(1)
	}
	data = strings.Split(string(content), "\n")
	return data
}

