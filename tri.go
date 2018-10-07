package main
import "fmt"

func main() {
	res := getRow(4)
	fmt.Println(res)
}

func getRow(rowIndex int) []int {
    res := []int{}
    ch := 1
    for i:= 0; i<=rowIndex; i++ {
        res = append(res, ch)
		// ch reuse
		// row same same 
		// i inc
		// row -i desc
		// i+1 inc
        ch = ch * (rowIndex - i) / (i+1)
    }
    
    return res;
}
