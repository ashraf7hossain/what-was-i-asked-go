package utils

import(
	"fmt"
)

func Map[T any, U any](list []T, f func(T) U) []U {
    result := make([]U, len(list))
    for i, v := range list {
        result[i] = f(v)
				fmt.Println("modified v =", f(v))
				fmt.Println()

    }
    return result
}

func Filter[T any](list []T, f func(T) bool) []T {
		result := make([]T, 0, len(list))
		for _, v := range list {
				if f(v) {
						result = append(result, v)
				}
		}
		return result
}

func Reduce[T any](list []T, f func(T, T) T) T {
		result := list[0]
		for _, v := range list[1:] {
				result = f(result, v)
		}
		return result
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

func IndexOf[T comparable](list []T, target T) (int , error) {
		for i, v := range list {
				if v == target {
						return i, nil
				}
		}
		return -1, fmt.Errorf("element not found")
}

