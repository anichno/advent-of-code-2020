package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/set"
)

func main() {
	inputFileName := os.Args[1]
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validFields := set.New("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid")
	optionalFields := set.New("cid")
	checkFields := validFields.Difference(optionalFields)

	passports := make([]map[string]string, 0)
	passport := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
			continue
		}
		for _, entryPair := range strings.Split(text, " ") {
			parts := strings.Split(entryPair, ":")
			passport[parts[0]] = parts[1]
		}
	}
	passports = append(passports, passport)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	reByr, _ := regexp.Compile("^[0-9]{4}$")
	reIyr, _ := regexp.Compile("^[0-9]{4}$")
	reEyr, _ := regexp.Compile("^[0-9]{4}$")
	reHgt, _ := regexp.Compile("^([0-9]+)(cm|in)$")
	reHcl, _ := regexp.Compile("^#[0-9a-f]{6}$")
	reEcl, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	rePid, _ := regexp.Compile("^[0-9]{9}$")

	validPassportsPart1 := 0
	validPassportsPart2 := 0
	for _, passport := range passports {
		passportSet := set.New()
		for k := range passport {
			passportSet.Insert(k)
		}
		if checkFields.Difference(passportSet).Len() == 0 {
			validPassportsPart1++

			// validate fields for part 2

			// byr (Birth Year) - four digits; at least 1920 and at most 2002.
			result := reByr.FindString(passport["byr"])
			val, err := strconv.Atoi(result)
			if err != nil || val < 1920 || val > 2002 {
				continue
			}

			// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
			result = reIyr.FindString(passport["iyr"])
			val, err = strconv.Atoi(result)
			if err != nil || val < 2010 || val > 2020 {
				continue
			}

			// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
			result = reEyr.FindString(passport["eyr"])
			val, err = strconv.Atoi(result)
			if err != nil || val < 2020 || val > 2030 {
				continue
			}

			// hgt (Height) - a number followed by either cm or in:
			// If cm, the number must be at least 150 and at most 193.
			// If in, the number must be at least 59 and at most 76.
			resultHgt := reHgt.FindStringSubmatch(passport["hgt"])
			if len(resultHgt) == 3 {
				val, err = strconv.Atoi(resultHgt[1])
				if err != nil || (resultHgt[2] == "cm" && (val < 150 || val > 193)) || (resultHgt[2] == "in" && (val < 59 || val > 76)) {
					continue
				}
			} else {
				continue
			}

			// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
			if !reHcl.MatchString(passport["hcl"]) {
				continue
			}

			// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
			if !reEcl.MatchString(passport["ecl"]) {
				continue
			}

			// pid (Passport ID) - a nine-digit number, including leading zeroes.
			if !rePid.MatchString(passport["pid"]) {
				continue
			}

			// cid (Country ID) - ignored, missing or not.

			validPassportsPart2++
		}
	}

	fmt.Println("Part 1:", validPassportsPart1)
	fmt.Println("Part 2:", validPassportsPart2)
}
