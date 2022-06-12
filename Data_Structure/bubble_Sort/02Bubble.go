package main

import ( //            SORTING SLICE WITH RANDOMLY GENERATED NUMBER USING THE rand.Perm
	"fmt"
	"math/rand"
	"time"
)

//1- math/rand package of GO provides a Perm method that can be used generate the pseudo-random slice of n integers. The array will be pseudo-random permutation of the integers in the range [0,n).
//2- Perm	This function returns, as a slice of n ints, a pseudo-random permutation of the integers [0, n) from the default source.

// 3-Here, we will generate the slice of random numbers using the rand.Perm() number. The rand.Perm() function returns the specified number of random numbers between 0-N. Here, N is the specified number in the rand.Perm() function.

func main() {
	rand.Seed(time.Now().UnixNano())
	t := rand.Perm(10) // generating slice with random number using permutation of rand package
	fmt.Println("Before sort: ", t)
	fmt.Println(demo(t))
}

func demo(t []int) []int {
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t)-1-i; j++ {
			if t[j] > t[j+1] {
				t[j], t[j+1] = t[j+1], t[j]
			}
		}
	}
	return t
}

//rand.perm package will only generate +ve numbers in the slice ,
// but for generating -ve numbers we need to use the rand.intn(999)-rand.intn(999)
