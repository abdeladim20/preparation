package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

// func main() {
// 	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
// 	fmt.Printf("%#v\n", Slice(a, 1))
// 	fmt.Printf("%#v\n", Slice(a, 2, 4))
// 	fmt.Printf("%#v\n", Slice(a, -3))
// 	fmt.Printf("%#v\n", Slice(a, -2, -1))
// 	fmt.Printf("%#v\n", Slice(a, 2, 0))
// }

/*
$ go run .
[]string{"algorithm", "ascii", "package", "golang"}
[]string{"ascii", "package"}
[]string{"ascii", "package", "golang"}
[]string{"package"}
[]string(nil)
*/

func main() {
	elems := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	s := random.StrSlice(chars.Words)

	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = random.StrSlice(chars.Words)
		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range elems {
		challenge.Function("Slice", Slice, solutions.Slice, a...)
	}
}

func Slice(a []string, nbrs ...int) []string {
	n := len(a)

	if len(nbrs) == 0 {
		return nil
	}



	start := nbrs[0]
	end := n

	if len(nbrs) == 1 {
		if start < 0 {
			return a[n+start:]
		}
	}

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	if start > 0 {
		if n < start {
			if n < end {
				return []string{}
			}
		} else {
			if n < end {
				return a[0+start: n]
			}
		}
	}

	if start > 0 && end < 0 || start < 0 && end > 0 {
		return nil
	}

	// Handle negative indices
	m := n * -1
	if start < 0 {
		if m > start {
			if m > end {
				return []string{}
			}else {
				if end != n {
					end = end * -1
					return a[0:n-end]
				}

			}
		}
		start = n + start
	}
	if end < 0 {
		end = n + end
	}

	// If start or end are out of bounds
	if start < 0 || start > n || end < 0 || end > n {
		return []string{}
	}

	// Return nil slice if end is less than or equal to start
	if end <= start {
		return nil
	}

	result := []string{}
	for i := start; i < end; i++ {
		result = append(result, a[i])
	}

	return result
}