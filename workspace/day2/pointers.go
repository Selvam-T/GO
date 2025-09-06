package main

import ( "fmt")

func swap(a *int, b *int) {

	fmt.Println("pointer a,b is", a, b)
	fmt.Println("value a,b is", *a, *b)
	temp := *b
	*b = *a
	*a = temp
}

/* parameter is full copy of array, changes don't persist */
func incArray(arr [3]int) {
	for i:= range arr {
		arr[i]++
	}
}

/* parameter is a pointer, changes persist */
func incArray2(arr *[3]int) {
	for i:= range arr {
		(*arr)[i]++
	}
}

/* parameter points to same arrayy, changes persist */
func incSlice(arr []int) {
	for i := range arr {
		arr[i]++	
	}

}

func main() {
	x := 10;
	y := 20;
	
	/* swap 2 numbers */
	fmt.Println("Before swap x, y is",x,y)
	swap(&x, &y)
	fmt.Println("After swap x, y is", x,y)

	/* Increment all elements of an array */
	arr := []int{10, 20, 30}
	brr := [3]int{100, 200, 300}

	fmt.Println(arr);
	incSlice(arr)
	fmt.Println(arr);

	incArray(brr);
	fmt.Println(brr);

	incArray2(&brr);
        fmt.Println(brr);

}	
