package hof_test

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/suryanshu-09/hof"
)

func TestSquare(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		inputArr := [][]int{
			{1, 2, 3, 4, 5},
			{-4, 0, 69, 12},
			{0},
		}

		outputArr := [][]int{
			{1, 4, 9, 16, 25},
			{16, 0, 4761, 144},
			{0},
		}

		for idx, val := range inputArr {
			var req []int
			for sq := range hof.Square(val) {
				req = append(req, sq)
			}
			if !reflect.DeepEqual(req, outputArr[idx]) {
				t.Errorf("got:%v\nwant:%v", req, outputArr[idx])
			}
		}
	})
	t.Run("floats", func(t *testing.T) {
		inputArr := [][]float64{
			{1.9, 2.2, 3.2, 4.0, 5.1},
			{-4.9, 0.1, 69.42, 12.7},
			{0.0},
		}

		outputArr := [][]float64{
			{3.61, 4.84, 10.24, 16, 26.01},
			{24.01, 0.01, 4819.1364, 161.29},
			{0},
		}

		for idx, val := range inputArr {
			var req []float64
			for sq := range hof.Square(val) {
				req = append(req, sq)
			}

			if !slicesAlmostEqual(req, outputArr[idx], 1e-9) {
				t.Errorf("case %d:\ngot:  %v\nwant: %v", idx, req, outputArr[idx])
			}
		}
	})
}

func slicesAlmostEqual(a, b []float64, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > eps {
			return false
		}
	}
	return true
}

func TestMap(t *testing.T) {
	t.Run("int to main", func(t *testing.T) {
		inputArr := [][]int{
			{1, 2, 3, 4, 5},
			{-4, 0, 69, 12},
			{0},
		}

		outputArr := [][]string{
			{"this is 1", "this is 2", "this is 3", "this is 4", "this is 5"},
			{"this is a neg 4", "this is empty", "this is 69", "this is 12"},
			{"this is empty"},
		}

		for idx, val := range inputArr {
			var req []string
			for i := range hof.Map(val, func(j int) string {
				switch {
				case j > 0:
					return fmt.Sprintf("this is %d", j)
				case j == 0:
					return "this is empty"
				case j < 0:
					return fmt.Sprintf("this is a neg %d", int(math.Abs(float64(j))))
				default:
					t.Errorf("incorrect input got %d", j)
				}
				return ""
			}) {
				req = append(req, i)
			}

			if !reflect.DeepEqual(req, outputArr[idx]) {
				t.Errorf("got:%v\nwant:%v", req, outputArr[idx])
			}
		}
	})

	t.Run("square", func(t *testing.T) {
		inputArr := [][]int{
			{1, 2, 3, 4, 5},
			{-4, 0, 69, 12},
			{0},
		}

		outputArr := [][]int{
			{1, 4, 9, 16, 25},
			{16, 0, 4761, 144},
			{0},
		}

		for idx, val := range inputArr {
			var req []int
			for sq := range hof.Map(val, func(x int) int { return x * x }) {
				req = append(req, sq)
			}
			if !reflect.DeepEqual(req, outputArr[idx]) {
				t.Errorf("got:%v\nwant:%v", req, outputArr[idx])
			}
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter coffee", func(t *testing.T) {
		inputArr := [][]string{
			{"tea", "soup", "coffee", "coke"},
			{"pan", "ghee", "mug", "spatula"},
			{"mustard", "egg", "meat"},
		}

		outputArr := [][]string{
			{"soup", "coke"},
			{"ghee"},
			{},
		}
		filterArr := []func(string) bool{
			func(eat string) bool {
				if eat == "soup" || eat == "coke" {
					return true
				}
				return false
			},
			func(utensil string) bool {
				return utensil == "ghee"
			},
			func(food string) bool {
				return food == "chicken"
			},
		}
		for idx, val := range inputArr {
			req := []string{}
			for v := range hof.Filter(val, filterArr[idx]) {
				req = append(req, v)
			}

			if !reflect.DeepEqual(req, outputArr[idx]) {
				t.Errorf("got:%v\nwant:%v", req, outputArr[idx])
			}
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("TestReduceSumInts", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := hof.Reduce(nums, func(acc, v int) int { return acc + v }, 0)
		want := 15
		if got != want {
			t.Errorf("Sum of ints: got %v, want %v", got, want)
		}
	})

	t.Run("TestReduceSumFloats", func(t *testing.T) {
		nums := []float64{1.2, 2.3, 3.4}
		got := hof.Reduce(nums, func(acc, v float64) float64 { return acc + v }, 0)
		want := 6.9
		if !floatAlmostEqual(got, want, 1e-9) {
			t.Errorf("Sum of floats: got %v, want %v", got, want)
		}
	})

	t.Run("TestReduceConcatStrings", func(t *testing.T) {
		words := []string{"Go", " ", "is", " ", "fun"}
		got := hof.Reduce(words, func(acc, v string) string { return acc + v }, "")
		want := "Go is fun"
		if got != want {
			t.Errorf("Concat strings: got %v, want %v", got, want)
		}
	})

	t.Run("TestReduceProductInts", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		got := hof.Reduce(nums, func(acc, v int) int { return acc * v }, 1)
		want := 24
		if got != want {
			t.Errorf("Product of ints: got %v, want %v", got, want)
		}
	})

	t.Run("TestReduceCountCondition", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		got := hof.Reduce(nums, func(acc, v int) int {
			if v%2 == 0 {
				return acc + 1
			}
			return acc
		}, 0)
		want := 3 // even numbers: 2, 4, 6
		if got != want {
			t.Errorf("Count evens: got %v, want %v", got, want)
		}
	})

	t.Run("TestReduceStructs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{"A", 20},
			{"B", 25},
			{"C", 30},
		}

		got := hof.Reduce(people, func(acc int, p Person) int { return acc + p.Age }, 0)
		want := 75
		if got != want {
			t.Errorf("Sum ages: got %v, want %v", got, want)
		}
	})

	t.Run("TestReduceEmptySlice", func(t *testing.T) {
		nums := []int{}
		got := hof.Reduce(nums, func(acc, v int) int { return acc + v }, 10)
		want := 10
		if got != want {
			t.Errorf("Empty slice: got %v, want %v", got, want)
		}
	})

	t.Run("TestReduce_StringCount", func(t *testing.T) {
		fruits := []string{"apple", "banana", "apple", "orange"}

		got := hof.Reduce(fruits, func(acc map[string]int, val string) map[string]int {
			acc[val]++
			return acc
		}, map[string]int{})

		want := map[string]int{"apple": 2, "banana": 1, "orange": 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Reduce() mismatch\n got:  %v\n want: %v", got, want)
		}
	})
}

func floatAlmostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
