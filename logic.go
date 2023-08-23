package main

import (
	"fmt"
	"regexp"
	"strings"
)

//																	  		  0000000
//       0000										  0000000				 000| |000
//      000000 ----->	 00000000					 000| |000					| |
//      0sun00 --->		0000| |000    000000			| |						| |
//      000000 ------->  	| |		 000| |000			| |		  0000000		| |
//       0000  				| |			| |				| |		 000| |000		| |

//Given are the heights in meters of certain Trees which lie adjacent to each other.
//Sunlight starts falling from the left side of the trees. If there is a tree of a certain Height,
//all the trees to the right side of it having lesser heights or not
//equal cannot see the sun. The task is to find the total number of
//such buildings that receive sunlight.
// Input :
// N = 5
// Arr[] = {7, 4, 8, 2, 9}
// Output : 3
// Exp: Only Trees of height 7, 8, 9 meters can see the sun hence output is 3

// Input :
// N = 5
// Arr[] = {1, 4, 8, 2, 9}
// Output : 4
// Exp : Only Trees of height 1, 4, 8, 9 meters can see the sun

// Input :
// N = 5
// Arr[] {1, 4, 4 ,2, 9}
// Output : 4
// Exp : Only Trees of height 1, 4, 4, 9 meters can see the sun

func Solution1(arr []int, n int) int {
	var tempArr []int
	tempArr = append(tempArr, arr[0])
	index1 := arr[0]
	for i := 0; i < len(arr); i++ {
		if i < len(arr) && index1 < arr[i] {
			tempArr = append(tempArr, arr[i])
			index1++
		}
	}
	return len(tempArr)
}

func Solution2(s string) bool {
	checkStr := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	newStr := strings.ToLower(checkStr)
	fmt.Println(newStr)
	lastIdx := len(newStr) - 1

	for i := 0; i <= lastIdx/2; i++ {
		if i < len(newStr) && newStr[i] != newStr[lastIdx-i] {
			return false
		}
	}
	return true
}

func Solution4(N int, itemList [][]int, M int, query [][]int) {
	for _, q := range query {
		totalItem := 0
		for _, il := range itemList {
			l, r, x, y := q[0], q[1], q[2], q[3]
			vv, qq := il[0], il[1]
			if (vv >= l && vv <= r) && (qq >= x && qq <= y) {
				totalItem++
			}
		}
		fmt.Println("RESULT-3: ", totalItem)
	}
}

func isValidParentheses(s string) bool {
	stack := make([]rune, 0)
	C := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != C[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
