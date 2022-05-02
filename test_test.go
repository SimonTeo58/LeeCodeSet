package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome9(t *testing.T) {
	x := 121
	isPalind := isPalindrome(x)
	if isPalind != true {
		t.Errorf("actually 121 is a Palindrome")
	}
}

func TestRomanToInt13(t *testing.T) {
	v := "VII"
	m := "MCMXCIV"
	romanInt := romanToInt(v)
	if romanInt != 7 {
		t.Errorf("value not equal 7")

	}
	romanInt_m := romanToInt(m)
	if romanInt_m != 1994 {
		t.Errorf("value not equal 1994")

	}
}

func TestLongestCommonPrefix14(t *testing.T) {
	flCase := []string{"flower", "flow", "flight"}
	commonLongStrs := longestCommonPrefix(flCase)
	if commonLongStrs != "fl" {
		t.Errorf("flCase's return not equal fl")
	}
}
func TestIsValid20(t *testing.T) {
	s1 := "{[()]}"
	// 123 {
	// 91  [
	// 40  (
	// 41	)
	// 93	]
	// 125	}
	s2 := "([)]"

	s1Checker := isValid(s1)
	if s1Checker != true {
		t.Errorf("s1's return not equal true")
	}

	s2Checker := isValid(s2)
	if s2Checker != false {
		t.Errorf("s2's return not equal false")
	}

}

func TestMergeTwoLists21(t *testing.T) {
	head1 := &ListNode{1, nil}
	head1.Next = &ListNode{2, nil}
	head1.Next.Next = &ListNode{4, nil}
	head2 := &ListNode{1, nil}
	head2.Next = &ListNode{3, nil}
	head2.Next.Next = &ListNode{4, nil}
	// head1 := &ListNode{1, nil}

	// head2 := &ListNode{2, nil}

	newList := mergeTwoLists(head1, head2)
	fmt.Printf("%+v", newList)

}

func TestMaxSubArray(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if maxValue := maxSubArray(arr); maxValue != 6 {
		t.Fatal("the result not equals 6")
	}
}

func TestPulsOne(t *testing.T) {

	digits := []int{9, 9, 9}
	newDigit := plusOne(digits)
	fmt.Println(newDigit)
}

func TestAddBinary(t *testing.T) {
	a := "11100011"
	b := "10101"
	newSum := addBinary(a, b)
	t.Fatal(newSum)
}
