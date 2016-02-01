package main 

import "fmt"

func main() {
	a := []int{7,6,5,4,3,2,1}
	fmt.Println(mergeSort(a))
}

func mergeSort (array []int) ([]int) {
	length := len(array)
	if (len(array) < 2){
		return array
	}
	A := mergeSort(array[0:length/2])
	B := mergeSort(array[length/2:])
	C := make([]int, length);

	// Merge
	j:=0 
	k := 0
	for i:=0; i<length; i++ {
		if(len(A) > j && len(B) > k){
			if A[j] <= B[k] {
				C[i] = A[j]
				j++
			} else {
				C[i] = B[k]
				k++
			}
		}else if(len(A) > j){
			C[i] = A[j]
			j++
		}else if(len(B) > k){
			C[i] = A[k]
			k++
		}

	}
	return C	
}