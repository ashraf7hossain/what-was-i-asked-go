package utils

import (
	"fmt"
)

// numbers := []int{1, 2, 3, 4, 5}
// 	squared := Map(numbers, func(x int) int {
// 		return x * x
// 	})
func Map[T any, U any](list []T, f func(T) U) []U {
	result := make([]U, len(list))
	for i, v := range list {
		result[i] = f(v)
	}
	return result
}
// numbers := []int{1, 2, 3, 4, 6}
// odds    := Filter(numbers, func(x int) int {
//            if x % 2 == 1 {// return x}
//})
func Filter[T any](list []T, f func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, v := range list {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// ints := []int{1, 2, 3, 4, 5}
//	sum := Reduce(ints, func(acc int, item int) int {
//		return acc + item
//	}, 0)
//	fmt.Println("Sum:", sum)
func Reduce[T any, R any](list []T, reducer func(acc R, item T) R, initial R) R {
	acc := initial
	for _, item := range list {
		acc = reducer(acc, item)
	}
	return acc
}

func ForEach[T any](list []T, f func(T)) {
	for _, v := range list {
		f(v)
	}
}

func Contains[T comparable](list []T, target T) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](list []T, target T) (int, error) {
	for i, v := range list {
		if v == target {
			return i, nil
		}
	}
	return -1, fmt.Errorf("element not found")
}
