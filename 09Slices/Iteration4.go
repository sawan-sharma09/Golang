package main //         ******SLICE ITERATION USING FOR LOOP******
import "fmt"

func main() {
	slc := []int{22, 44, 56, 67}
	for i := 0; i < len(slc); i++ {
		fmt.Println(slc[i])
	}
}
