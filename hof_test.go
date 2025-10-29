package hof_test

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/suryanshu-09/hof"
)

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

func TestForEach(t *testing.T) {
	t.Run("test array transformation", func(t *testing.T) {
		var buf bytes.Buffer

		inpArr := []int{1, 2, 3, 4, 5}
		hof.ForEach(inpArr, func(x int) {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		})

		got := buf.String()
		want := "1\n2\n3\n4\n5\n"

		if got != want {
			t.Errorf("ForEach() = %q, want %q", got, want)
		}
	})
}

func TestFind(t *testing.T) {
	t.Run("find cherry", func(t *testing.T) {
		inputArr := []string{"banana", "apple", "cherry", "orange"}
		cherry, found := hof.Find(inputArr, func(fruit string) bool {
			return fruit == "cherry"
		})

		if !found {
			t.Error("not found")
		}

		if cherry != "cherry" {
			t.Errorf("got:%s\nwant:%s", cherry, "cherry")
		}
	})
	t.Run("find cherry not", func(t *testing.T) {
		inputArr := []string{"banana", "apple", "orange"}
		cherry, found := hof.Find(inputArr, func(fruit string) bool {
			return fruit == "cherry"
		})

		if found {
			t.Error("incorrect found")
		}

		if cherry != "" {
			t.Errorf("got:%s\nwant:%s", cherry, "")
		}
	})
}

func TestSome(t *testing.T) {
	t.Run("greater than 5", func(t *testing.T) {
		inputArr := []int{1, 2, 3, 6, 6, 7, 8}
		want := true
		if out := !hof.Some(inputArr, func(x int) bool {
			return x > 5
		}); out {
			t.Errorf("got:%v\nwant:%v", out, want)
		}
	})

	t.Run("greater than 5 not", func(t *testing.T) {
		inputArr := []int{1, 2, 3, 4}
		want := false
		if out := hof.Some(inputArr, func(x int) bool {
			return x > 5
		}); out {
			t.Errorf("got:%v\nwant:%v", out, want)
		}
	})
}

func TestEvery(t *testing.T) {
	t.Run("greater than 5", func(t *testing.T) {
		inputArr := []int{6, 6, 7, 8}
		want := true
		if out := !hof.Every(inputArr, func(x int) bool {
			return x > 5
		}); out {
			t.Errorf("got:%v\nwant:%v", out, want)
		}
	})

	t.Run("greater than 5 not", func(t *testing.T) {
		inputArr := []int{5, 7, 8, 1, 3}
		want := false
		if out := hof.Every(inputArr, func(x int) bool {
			return x > 5
		}); out {
			t.Errorf("got:%v\nwant:%v", out, want)
		}
	})
}

func TestSquare(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		inputArr := [][]int{
			{1, 2, 3, 4, 5},
			{-4, 0, 69, 12},
			{0},
			{},
		}

		outputArr := [][]int{
			{1, 4, 9, 16, 25},
			{16, 0, 4761, 144},
			{0},
			{},
		}

		for idx, val := range inputArr {
			req := []int{}
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
			{},
		}

		outputArr := [][]float64{
			{3.61, 4.84, 10.24, 16, 26.01},
			{24.01, 0.01, 4819.1364, 161.29},
			{0},
			{},
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

func TestCube(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		inputArr := [][]int{
			{1, 2, 3, 4, 5},
			{-2, 0, 3, -4},
			{0},
			{},
		}

		outputArr := [][]int{
			{1, 8, 27, 64, 125},
			{-8, 0, 27, -64},
			{0},
			{},
		}

		for idx, val := range inputArr {
			req := []int{}
			for cube := range hof.Cube(val) {
				req = append(req, cube)
			}

			if !reflect.DeepEqual(req, outputArr[idx]) {
				t.Errorf("got:%v\nwant:%v", req, outputArr[idx])
			}
		}
	})
	t.Run("floats", func(t *testing.T) {
		inputArr := [][]float64{
			{1.0, 2.0, 3.0},
			{-2.0, 0.5, 1.5},
			{0.0},
			{},
		}

		outputArr := [][]float64{
			{1.0, 8.0, 27.0},
			{-8.0, 0.125, 3.375},
			{0.0},
			{},
		}

		for idx, val := range inputArr {
			var req []float64
			for cube := range hof.Cube(val) {
				req = append(req, cube)
			}

			if !slicesAlmostEqual(req, outputArr[idx], 1e-9) {
				t.Errorf("case %d:\ngot:  %v\nwant: %v", idx, req, outputArr[idx])
			}
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		testCases := []struct {
			input []int
			want  int
		}{
			{[]int{1, 2, 3, 4, 5}, 15},
			{[]int{-1, -2, -3}, -6},
			{[]int{0, 0, 0}, 0},
			{[]int{100}, 100},
			{[]int{}, 0},
			{[]int{-5, 5, -10, 10}, 0},
		}

		for _, tc := range testCases {
			got := hof.Sum(tc.input)
			if got != tc.want {
				t.Errorf("Sum(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})

	t.Run("floats", func(t *testing.T) {
		testCases := []struct {
			input []float64
			want  float64
		}{
			{[]float64{1.5, 2.5, 3.0}, 7.0},
			{[]float64{-1.1, -2.2, -3.3}, -6.6},
			{[]float64{0.0, 0.0}, 0.0},
			{[]float64{3.14}, 3.14},
			{[]float64{}, 0.0},
		}

		for _, tc := range testCases {
			got := hof.Sum(tc.input)
			if !floatAlmostEqual(got, tc.want, 1e-9) {
				t.Errorf("Sum(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})
}

func TestAverage(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		testCases := []struct {
			input []int
			want  float64
		}{
			{[]int{1, 2, 3, 4, 5}, 3.0},
			{[]int{-2, -4, -6}, -4.0},
			{[]int{0, 0, 0}, 0.0},
			{[]int{100}, 100.0},
			{[]int{10, 20}, 15.0},
			{[]int{1, 3, 5, 7, 9}, 5.0},
		}

		for _, tc := range testCases {
			got := hof.Average(tc.input)
			if !floatAlmostEqual(got, tc.want, 1e-9) {
				t.Errorf("Average(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})

	t.Run("floats", func(t *testing.T) {
		testCases := []struct {
			input []float64
			want  float64
		}{
			{[]float64{1.0, 2.0, 3.0}, 2.0},
			{[]float64{-1.5, -2.5}, -2.0},
			{[]float64{5.5}, 5.5},
			{[]float64{0.1, 0.2, 0.3}, 0.2},
			{[]float64{10.5, 20.5, 30.0}, 20.333333333333332},
		}

		for _, tc := range testCases {
			got := hof.Average(tc.input)
			if !floatAlmostEqual(got, tc.want, 1e-9) {
				t.Errorf("Average(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		testCases := []struct {
			input []int
			want  int
		}{
			{[]int{1, 2, 3, 4, 5}, 1},
			{[]int{5, 4, 3, 2, 1}, 1},
			{[]int{-1, -2, -3}, -3},
			{[]int{0, -5, 10}, -5},
			{[]int{100}, 100},
			{[]int{7, 7, 7}, 7},
			{[]int{}, 0}, // empty slice returns zero value
		}

		for _, tc := range testCases {
			got := hof.Min(tc.input)
			if got != tc.want {
				t.Errorf("Min(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})

	t.Run("floats", func(t *testing.T) {
		testCases := []struct {
			input []float64
			want  float64
		}{
			{[]float64{1.5, 2.5, 0.5}, 0.5},
			{[]float64{-1.1, -2.2, -0.5}, -2.2},
			{[]float64{3.14}, 3.14},
			{[]float64{5.0, 5.0, 5.0}, 5.0},
			{[]float64{}, 0.0}, // empty slice returns zero value
		}

		for _, tc := range testCases {
			got := hof.Min(tc.input)
			if !floatAlmostEqual(got, tc.want, 1e-9) {
				t.Errorf("Min(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		testCases := []struct {
			input []int
			want  int
		}{
			{[]int{1, 2, 3, 4, 5}, 5},
			{[]int{5, 4, 3, 2, 1}, 5},
			{[]int{-1, -2, -3}, -1},
			{[]int{0, -5, 10}, 10},
			{[]int{100}, 100},
			{[]int{7, 7, 7}, 7},
		}

		for _, tc := range testCases {
			got := hof.Max(tc.input)
			if got != tc.want {
				t.Errorf("Max(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})

	t.Run("floats", func(t *testing.T) {
		testCases := []struct {
			input []float64
			want  float64
		}{
			{[]float64{1.5, 2.5, 0.5}, 2.5},
			{[]float64{-1.1, -2.2, -0.5}, -0.5},
			{[]float64{3.14}, 3.14},
			{[]float64{5.0, 5.0, 5.0}, 5.0},
		}

		for _, tc := range testCases {
			got := hof.Max(tc.input)
			if !floatAlmostEqual(got, tc.want, 1e-9) {
				t.Errorf("Max(%v) = %v, want %v", tc.input, got, tc.want)
			}
		}
	})
}
