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

	maxValue := []int{2, 1, 7, 6}
	findMax := FindMaxValue(maxValue)
	fmt.Println(findMax)

	quantity := []int32{10, 10, 8, 9, 1}
	getRes := getMaximumAmount(quantity, 6)
	fmt.Println(getRes)

	quantitys := []int32{8, 8, 8, 8}
	ress := getMaximumAmount(quantitys, 4)
	fmt.Println(ress)

	op := []string{
		"1 1",
		"3",
		"1 12",
		"2",
		"3",
		"1 10",
		"3"}
	result := getMax(op)
	fmt.Println(result)

	convert := convertTemperature(36.50)
	fmt.Println(convert)

	reverse := reverseInt(123)
	fmt.Println(reverse)
}
