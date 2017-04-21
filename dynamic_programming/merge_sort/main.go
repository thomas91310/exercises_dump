package main

import "fmt"

func main() {
	a := []int{2, 4, 1, 3, 8, 7, 5, 6, 9}
	fmt.Println("array to sort", a)
	result := mergeSort(a)
	fmt.Println("result : ", result)
}

// [2,4,1,3]
// [2,4] [1,3]
// stop

// [8,7,5,6,9]
// [8,7] [5,6,9]
// [8,7] [5,6] [9]
// actually sort
// [7,8] [5,6] [9]
// [5,7,]

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return []int{a[0]}
	}
	left := a[:len(a)/2]
	right := a[len(a)/2 : len(a)]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	c := []int{}
	i := 0
	j := 0
	for k := 0; k < len(left)+len(right); k++ {
		if i >= len(left) {
			c = append(c, right[j:]...)
			break
		}
		if j >= len(right) {
			c = append(c, left[i:]...)
			break
		}
		fmt.Println("left -> ", left)
		fmt.Println("right -> ", right)

		if left[i] < right[j] {
			c = append(c, left[i])
			i++
		} else {
			c = append(c, right[j])
			j++
		}
	}
	fmt.Printf("Result of left %v and right %v : %v\n", left, right, c)
	return c
}

// Runtime analysis

// 1. How many levels are there based on the input size n ?

// 8 -> 4 levels

//                         0
//                 4               4
//             2       2       2       2
//            1 1     1 1     1 1     1 1

// 16 -> 5 levels

//                         0
//             8                       8
//         4      4               4       4
//       2   2  2   2           2   2   2    2
//      1 1 1 1 1 1 1 1       1  1 1 1 1 1  1  1

// log base2(n) + 1 levels (because first level is level 0)

// 2. How many distinct sub-problems are there for level j ?
// Let's take n = 16
// level 0: subproblem 1 (2^0)
// level 1: subproblem 2 (2^1)
// level 2: subproblem 4 (2^2)
// level 3: subproblem 8 (2^4)

// so number of subproblems = 2^j

// 3. What is the input size at each level j ?
//  input size of each subproblem
// e.g : for n = 16
// level 0 : array is of size 16 (n)
// level 1 : arrays are of size 8 (n/2*1)
// level 2 : arrays are of size 4 (n/2*2)
// level 3: arrays are of size 2 (n/2*2*2)
// so input size is : n/(2^j)

// 4. conclusion
// at each level j :
// subproblem (basically what's happening for a level)
// j = 0,1....log(n) ... there are 2^j subproblems each of size n / 2^j
// input of size n -> 6n operations

// so 2^j * 6(n/2^j) is the work done at each level ==> 6n opeartions at each level
// so number of levels is log2(n) so complexity is O(6nlog(n) + 6n) => O(nlog(n))
