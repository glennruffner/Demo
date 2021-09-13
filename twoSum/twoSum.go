package twoSum

import (
	"fmt"
	"strconv"
)

func TwoSum(fld1, fld2 string) string {
	answer := ""
	index1 := len(fld1) - 1
	index2 := len(fld2) - 1
	if len(fld1) == 0 {
		return fld2 // nothing to add to fld2
	}
	if len(fld2) == 0 {
		return fld1 // nothing to add to fld1
	}
	carry := 0
	for index1 >= 0 && index2 >= 0 {
		s1 := fld1[index1 : index1+1]
		s2 := fld2[index2 : index2+1]
		d1, err := strconv.Atoi(s1)
		if err != nil {
			fmt.Println("Error")
		}
		d2, err := strconv.Atoi(s2)
		if err != nil {
			fmt.Println("Error")
		}
		a1 := d1 + d2 + carry
		carry = 0
		if a1 > 10 {
			carry = carry + 1
			a1 = a1 - 10
		}
		answer = strconv.Itoa(a1) + answer
		index1--
		index2--
	}
	fmt.Println(index1, index2, answer, " ", carry)
	if index1 >= 0 {
		s1 := fld1[index1 : index1+1]
		fmt.Println("s1:", s1)
		fmt.Println("carry:", carry)
		d1, err := strconv.Atoi(s1)
		if err != nil {
			fmt.Println("Error")
		}
		a1 := d1 + carry
		carry = 0
		if a1 > 10 {
			carry = carry + 1
			a1 = a1 - 10
		}
		fmt.Println("a1:", a1)
		answer = strconv.Itoa(a1) + answer
		fmt.Println("temp answer:", answer)
		index1--
	}
	if index2 >= 0 {
		s2 := fld2[index2 : index2+1]
		d2, err := strconv.Atoi(s2)
		if err != nil {
			fmt.Println("Error")
		}
		a2 := d2 + carry
		carry = 0
		if a2 > 10 {
			carry = carry + 1
			a2 = a2 - 10
		}
		answer = strconv.Itoa(a2) + answer
		index2--
	}
	if index1 >= 0 {
		fmt.Println("remaining added:", fld1[0:index1+1])
		answer = fld1[0:index1+1] + answer
	}
	if index2 >= 0 {
		answer = fld2[0:index2+1] + answer
	}
	fmt.Println("returning:", answer)
	return answer
}
