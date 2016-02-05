package main 

import "fmt"

func main() {
	a := []int{7,6,5,4,3,2,1}
	fmt.Println(countInv(a))
}

func countInv(array []int) (int, []int) {
	if len(array) < 2 {
		return 0,array
	}	

	invA, A := countInv(array[0:len(array)/2])
	invB, B := countInv(array[len(array)/2:])
	invSpl, splArray := countSplInv(A,B)

	return invA + invB + invSpl, splArray
}
func countSplInv(a []int, b []int) (int, []int) {
	i, j, invCount := 0 , 0 , 0
 	c := make([]int,len(a) + len(b))

	for k := 0; k < len(c); k++ {
		if len(b) > j && len(a) > i {
			if a[i] > b[j] {
					invCount += len(a) - i
					c[k] = b[j]
					j++
			}	else {
				c[k] = a[i]
				i++
			}
		} else if len(b) > j {
			c[k] = b[j]
			j++
		} else if len(a) > i {
			c[k] = a[i]
			i++
		}
	}	

	return invCount, c
}

