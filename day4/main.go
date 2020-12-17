package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main(){
	arg := os.Args[1]
	requirments := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	data := ReadFile(arg)

	passportsData := GetPassports(data)

	count1 := GetBadPassportsPartOne(passportsData, requirments)
	fmt.Println("RESULT I: ", count1)

	count2 := GetBadPassportsPartTwo(passportsData, requirments)
	fmt.Println("RESULT II: ", count2)

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

func GetPassports(data []string) (pd map[int][]string) {
	counter := 1
	temp := []string{}
	passportsData := make(map[int][]string)

	for _, line := range data {
		elem := strings.Split(line, " ")
		for _, field := range elem{
			if len(field) != 0 {
				temp = append(temp, field)
			}
		}
		passportsData[counter] = temp

		if elem[0] == "" {
			counter += 1
			temp = nil
		}
	}

	return passportsData
}

func GetBadPassportsPartTwo(passports map[int][]string, requirments []string) (count int) {
	badPass := 0
	totalPass := len(passports)
	checkField := false

	for _, pass := range passports {

		checkField = false

		if len(pass) < (len(requirments) - 1) {
			badPass += 1

		} else if len(pass) == 7 && FindString("cid", pass){

			badPass += 1
			checkField = true

	    } else {
			for _, data := range pass {

				if checkField {
					break
				}

				value := strings.Split(data, ":")[1]

				switch field := strings.Split(data, ":")[0]; field {

				//byr (Birth Year) - four digits; at least 1920 and at most 2002.
				case "byr":
					i, _ := strconv.Atoi(value)
					if len(value) != 4 || i < 1920 || i > 2002 {
						badPass, checkField = CheepCount(checkField, badPass, value, field)
					}
				//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
				case "iyr":
					i, _ := strconv.Atoi(value)
					if len(value) != 4 || i < 2010 || i > 2020 {
						badPass, checkField = CheepCount(checkField, badPass, value, field)
					}
				//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
				case "eyr":
					i, _ := strconv.Atoi(value)
					if len(value) != 4 || i < 2020 || i > 2030 {
						badPass, checkField = CheepCount(checkField, badPass, value, field)
					}

			   //hgt (Height) - a number followed by either cm or in:
			   		//If cm, the number must be at least 150 and at most 193.
			   		//If in, the number must be at least 59 and at most 76.
				case "hgt":
					regex, _ := regexp.Compile("[0-9]{1,}(cm|in)$")
					if regex.MatchString(value) != true {
						badPass, checkField = CheepCount(checkField, badPass, value, field)
					}
					if strings.Contains(value, "cm") {
						nr, _ := strconv.Atoi(strings.Split(value, "cm")[0])
						if nr < 150 || nr > 193 {
							 badPass, checkField = CheepCount(checkField, badPass, value, field)
						 }
					} else if strings.Contains(value, "in") {
						nr, _ := strconv.Atoi(strings.Split(value, "in")[0])
						if (nr < 59) || (nr > 76) {
							badPass, checkField = CheepCount(checkField, badPass, value, field)
						}
					}
				//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
				case "hcl":
					regex, _ := regexp.Compile("(^#)([a-zA-Z0-9]{6,6}$)")
					if regex.MatchString(value) != true {
						badPass, checkField = CheepCount(checkField, badPass, value, field)
					}

				//exactly one of: amb blu brn gry grn hzl oth.
				case "ecl":
					regex, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth){1}?$")
					if regex.MatchString(value) != true {
						badPass, checkField = CheepCount(checkField, badPass, value, field)
					}

				//a nine-digit number, including leading zeroes
				case "pid":
					regex, _ := regexp.Compile("^[0-9]{9}$")
					if regex.MatchString(value) != true {
						badPass, checkField = CheepCount(checkField, badPass, value, field)
					}
				}
			}
		}
	}
	return totalPass-badPass
}

func CheepCount(checkField bool, badPass int, value string, field string) (int, bool){
	badPass += 1
	checkField = true
	//fmt.Printf("%s does not meet chritria for %s\n", value, field)
	return badPass, checkField
}

func GetBadPassportsPartOne(passports map[int][]string, requirments []string) (count int){
	badPassport := 0
	totalPassports := len(passports)

	for _, pass := range passports {
		if len(pass) < (len(requirments)-1) {
			badPassport += 1
		} else if len(pass) == 7 && FindString("cid", pass){
			badPassport += 1
		}
	}

	return totalPassports-badPassport

}
func FindString(required string, passport []string) bool {
	for _, field := range passport{
		d := strings.Split(field, ":")[0]
		if required == d {
			return true
		}
	}
	return false
}

