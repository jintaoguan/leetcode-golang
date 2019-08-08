package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr1("1.1.1.1"))
	fmt.Println(defangIPaddr2("1.1.1.1"))
}

func defangIPaddr1(address string) string {
	splits := strings.Split(address, ".")
	ret := splits[0]
	for index, segment := range splits {
		if index == 0 {
			continue
		}
		ret += "[.]"
		ret += segment
	}
	return ret
}

func defangIPaddr2(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}

