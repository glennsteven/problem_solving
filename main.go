package main

import "fmt"

func main() {
	N := 5
	//=============== No 1 ===============
	Arr1 := []int{1, 4, 4, 2, 9}
	res1 := Solution1(Arr1, N)
	fmt.Println("RESULT-1: ", res1)
	//=============== No 1 ===============

	//=============== No 2 ===============
	res2 := Solution2("A man, a plan, a canal: Panama")
	fmt.Println("RESULT-2: ", res2)
	//=============== No 2 ===============

	//=============== No 4 ===============
	M := 0
	itemList := [][]int{{5, 2}, {4, 2}, {7, 7}, {8, 3}, {4, 5}}
	query := [][]int{{5, 5, 2, 2}, {4, 7, 5, 7}, {4, 8, 2, 7}}
	Solution4(N, itemList, M, query)
	//=============== No 4 ===============

	//=============== Valid Parentheses ===============
	res := isValidParentheses("()[]")
	fmt.Println("Valid Parentheses: ", res)
	//=============== Valid Parentheses ===============

	//=============== Single Number ===============
	arr := []int{1, 1, 2, 3, 3}
	isSingleNumber := singleNumber(arr)
	fmt.Println("isSingleNumber: ", isSingleNumber)
	//=============== Single Number ===============

	//=============== Remove Element ===============
	element := []int{1, 2, 2, 4}
	remove := 2
	resRemoveElement := removeElement(element, remove)
	fmt.Println("RemoveElement: ", resRemoveElement)
	//=============== Remove Element ===============
}
