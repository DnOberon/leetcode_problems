package twonumbers

import (
	"math/big"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// no need to reverse through list, simple get the array of ints from each
	list1 := getVals(l1)
	list2 := getVals(l2)

	// once you have the ints in hand, you'll need to convert them into strings DO NOT FORGET BIG INT
	return buildList(addLists(list1, list2))
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	result := ListNode{}
	result.Val = vals[0]
	result.Next = buildList(vals[1:])

	return &result
}

func getVals(list *ListNode) []int {
	if list == nil {
		return []int{}
	}
	results := []int{}

	results = append(results, list.Val)

	results = append(results, getVals(list.Next)...)

	return results
}

func addLists(l1, l2 []int) []int {
	l1String := []string{}
	l2String := []string{}

	// they want us to loop backwards through the lists
	for i := len(l1) - 1; i >= 0; i-- {
		newInt := big.NewInt(int64(l1[i]))
		l1String = append(l1String, newInt.String())
	}

	for i := len(l2) - 1; i >= 0; i-- {
		newInt := big.NewInt(int64(l2[i]))
		l2String = append(l2String, newInt.String())
	}

	a := new(big.Int)
	b := new(big.Int)
	a.SetString(strings.Join(l1String, ""), 10)
	b.SetString(strings.Join(l2String, ""), 10)

	result := strings.Split(a.Add(a, b).String(), "")
	revResult := []int{}

	for i := len(result) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(result[i])
		revResult = append(revResult, n)
	}

	return revResult
}
