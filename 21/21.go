package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var arr1 = []int{1, 2, 4}
var arr2 = []int{1, 3, 4}
var nodeList1 = &ListNode{}
var headList1 *ListNode
var nodeList2 = &ListNode{}
var headList2 *ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//fmt.Println("---mergeTwoLists-function---")

	merge := &ListNode{}
	var mergeHeader *ListNode
	if list1 == nil && list2 == nil {
		return mergeHeader
	}
	//fmt.Printf("list1 = %v\n", list1)
	//fmt.Printf("list2 = %v\n", list2)
	//fmt.Printf("merHed = %v\n", mergeHeader)
	mergeHeader = merge
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 != nil && list2 == nil { //001
			//fmt.Printf("001___list1.Val = %d\n", list1.Val)
			merge.Val = list1.Val
			list1 = list1.Next
		} else if list1 == nil && list2 != nil { //002
			//fmt.Printf("002___list2.Val = %d\n", list2.Val)
			merge.Val = list2.Val
			list2 = list2.Next
		} else if list1 != nil && list2 != nil && list1.Val < list2.Val { //004
			//fmt.Printf("004___list1.Val = %d\n", list1.Val)
			merge.Val = list1.Val
			list1 = list1.Next
		} else if list2 != nil { //005
			//fmt.Printf("005___list2.Val = %d\n", list2.Val)
			merge.Val = list2.Val
			list2 = list2.Next
		} else { //006
			//fmt.Printf("006___Break\n")
			break
		}
		//fmt.Printf("Step \n")
		if list1 == nil && list2 == nil {
			break
		}
		merge.Next = &ListNode{}
		merge = merge.Next
	}
	//fmt.Printf("merge addr = %p\n", merge)
	//fmt.Printf("mergeHeader addr = %p\n", mergeHeader)

	return mergeHeader

}

func (nl *ListNode) Addback(addval int, head *ListNode) (*ListNode, *ListNode) {
	ele := &ListNode{
		Val: addval,
	}
	if head == nil {
		head = ele
		nl = ele
	} else {
		nl.Next = ele
	}
	return ele, head
}
func printSingleLinkedList(head *ListNode) {
	for {
		if head != nil {
			fmt.Printf("node = %d\n", head.Val)
			if head.Next == nil {
				break
			}
			head = head.Next
		} else {
			break
		}
	}
}
func main() {
	headList1 = nil
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("array1 = %v\n", arr1)
	for _, val := range arr1 {
		nodeList1, headList1 = nodeList1.Addback(val, headList1)
	}
	printSingleLinkedList(headList1)
	headList2 = nil
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("array2 = %v\n", arr2)
	for _, val := range arr2 {
		nodeList2, headList2 = nodeList2.Addback(val, headList2)
	}
	printSingleLinkedList(headList2)

	// -----------------------------------------------------------------------------------
	printSingleLinkedList(mergeTwoLists(headList1, headList2))
	fmt.Printf("Result = %v", mergeTwoLists(headList1, headList2))
}
