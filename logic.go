package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
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

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		// ^= used XoR bitwise
		result ^= nums[i]
	}
	return result
}

func removeElement(nums []int, val int) int {
	//first, check length array of int
	//if length array of int less than zero
	//return 0
	if len(nums) <= 0 {
		return 0
	}

	//we initiate value i = 0
	i := 0
	for _, num := range nums {
		if num != val {
			//iteration i when num(index of array) != val (Values to be removed from the array elements)
			nums[i] = num
			i++
		}
	}

	return i
}

// FindMaxValue Interview User, Live Coding Bank Aladdin
func FindMaxValue(arr []int) int {
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}

func FindSecondMaxValue(arr []int, sMax int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			sMax = max
			max = arr[i]
		}

		if sMax < arr[i] && arr[i] < max {

		}
	}

	return sMax
}

// Pembeli = 6 orang
// [10, 10, 8, 9, 1] -> Desc
// Pembeli ->1
// Barang
// 	A , B , C, D, E
// [10, 10, 9, 8, 1] -> total jenis barang 5
//		Ambil index paling mahal -> [0]
//		Ambil barang A (10-1), sisa barang A -> 9
// Revenue = 10
// Pembeli ->2
// Barang
// 	A, B , C, D, E
// [9, 10, 9, 8, 1] -> total jenis barang 5
//		 B , A, C, D, E
// Sort [10, 9, 9, 8, 1]
//		Ambil index paling mahal -> [0]
//		Ambil Barang B (10-1), sisa barang B -> 9
// Revenue = 20
// Pembeli ->3
// Barang
// 	A, B, C, D, E
// [9, 9, 9, 8, 1] -> total jenis barang 5
//		 A, B, C, D, E
// Sort [9, 9, 9, 8, 1]
//		Ambil index paling mahal -> [0]
//		Ambil Barang A (9-1), sisa barang A -> 8
// Revenue = 29
// Pembeli ->4
// Barang
// 	A, B, C, D, E
// [8, 9, 9, 8, 1] -> total jenis barang 5
//		 B, C, A, D, E
// Sort [9, 9, 8, 8, 1]
//		Ambil index paling mahal -> [0]
//		Ambil Barang B (9-1), sisa barang B -> 8
// Revenue = 38
// Pembeli ->5
// Barang
// 	A, B, C, D, E
// [8, 8, 9, 8, 1] -> total jenis barang 5
//		 C, B, A, D, E
// Sort [9, 8, 8, 8, 1]
//		Ambil index paling mahal -> [0]
//		Ambil Barang C (9-1), sisa barang C -> 8
// Revenue = 47
// Pembeli ->6
// Barang
// 	A, B, C, D, E
// [8, 8, 8, 8, 1] -> total jenis barang 5
//		 A, B, C, D, E
// Sort [8, 8, 8, 8, 1]
//		Ambil index paling mahal -> [0]
//		Ambil Barang A (9-1), sisa barang A -> 7
// Revenue = 55

func getMaximumAmount(quantity []int32, m int32) int64 {
	sort.Slice(quantity, func(i, j int) bool {
		return quantity[i] > quantity[j]
	})

	var revenue int64
	for i := 0; i < int(m); i++ {
		revenue += int64(quantity[0])
		quantity[0]--
		sort.Slice(quantity, func(i, j int) bool {
			return quantity[i] > quantity[j]
		})
	}

	return revenue
}

// 1 x  -Push the element x into the stack.
// 2    -Delete the element present at the top of the stack.
// 3    -Print the maximum element in the stack.

// stack []
// n = 10
// 1 97 -> 1 push data ke stack
// stack -> [97]
// 2 -> delete data di stack
// stack -> []
// 1 20 -> push data ke stack
// stack -> [20]
// 2 -> delete data di stack
// stack -> []
// 1 26 -> push data ke stack
// stack -> [26]
// 1 20 -> push data ke stack
// stack -> [20,26]
// 2 -> delete data di stack
// stack -> [26]
// 3 -> print nilai max di stack = 26
// 1 91 -> push data ke stack
// stack -> [91,26]
// 3 -> print nilai max di stack = 91
func getMax(operations []string) []int32 {
	var (
		stack  []int32
		result []int32
	)

	for i := 0; i < len(operations); i++ {
		op := strings.Fields(operations[i])
		fmt.Println("FIELD", op)
		if len(op) > 0 {
			switch op[0] {
			case "1":
				if len(op) > 1 {
					max, _ := strconv.Atoi(op[1])
					stack = append(stack, int32(max))
				}
			case "2":
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			case "3":
				if len(stack) > 0 {
					max := stack[0]
					for m := 0; m < len(stack); m++ {
						if max < stack[m] {
							max = stack[m]
						}
					}
					result = append(result, max)
				}
			default:
				fmt.Printf("Invalid operation: %s\n", operations[i])
			}
		}
	}

	return result
}

func convertTemperature(celsius float64) []float64 {
	var output []float64

	kelvin := celsius + 273.15
	output = append(output, kelvin)
	fahrenheit := (celsius * 1.80) + 32.00
	output = append(output, fahrenheit)

	return output
}

func reverseInt(x int) int {
	res := 0
	for x > 0 || x < 0 {
		mod := x % 10
		res = (res * 10) + mod
		x /= 10
	}
	return res
}

func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	count := 1
	flagging := 1
	mayoritas := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count > flagging {
			flagging = count
			mayoritas = nums[i]
		}
	}

	return mayoritas
}

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) Add(value int) {
	newMode := &Node{
		Val:  value,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = newMode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newMode
}

func (l *List) Remove(val int) {
	if l.Head == nil {
		return
	}

	if l.Head.Val == val {
		l.Head = l.Head.Next
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Val != val {
		curr = curr.Next
	}

	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func printList(l *List) {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

// Print Traverse Linked List
func (l *List) Print() {
	if l.Head == nil {
		fmt.Println("No nodes in list")
	}
	ptr := l.Head
	for ptr != nil {
		fmt.Println("Node: ", ptr.Val)
		ptr = ptr.Next
	}
}

func (l *List) Search(val int) int {
	curr := l.Head
	for curr != nil {
		if curr.Val == val {
			return curr.Val
		}
		curr = curr.Next
	}
	return -1
}

func romanToInt(s string) int {
	result := 0
	newS := strings.ToUpper(s)
	length := len(newS)
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < length; i++ {
		abjad := newS[i : i+1]
		value := roman[abjad]
		nextValue := 0
		if i+1 < length {
			next := newS[i+1 : i+2]
			nextValue = roman[next]
		}

		if nextValue > value {
			result -= value
		} else {
			result += value
		}
	}
	return result
}
