package main

import "fmt"
import "strings"
import "sort"

func Sum(x int, y int) int {
    return x + y
}

func ReturnInt() int {
    return 1
}

func ReturnFloat() float32 {
    return 1.1
}

func ReturnIntArray() [3]int {
    return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
    return []int{1, 2, 3}
}

func IntSliceToString(slice []int) string {
    delim := ""
    return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), delim), "[]")
}

func MergeSlices(sl1 []float32, sl2 []int32) []int {
    newsl1 := make([]int, len(sl1))
    var i int
    var v float32
    for i, v = range sl1 {
        newsl1[i] = int(v)
    }
    newsl2 := make([]int, len(sl2))
    var vv int32
    for i, vv = range sl2 {
        newsl2[i] = int(vv)
    }
    return append(newsl1, newsl2...)
}

func GetMapValuesSortedByKey(mm map[int]string) []string {
    newsl := make([]int, 0, len(mm))
    sortedSl := make([]string, 0, len(mm))
    for k := range mm {
        newsl = append(newsl, k)
    }
    sort.Ints(newsl)

    for _, k := range newsl {
        sortedSl = append(sortedSl, mm[k])
    }

    return sortedSl
}

func main() {

    input := map[int]string{
	    10: "Зима",
	    30: "Лето",
	    20: "Весна",
	    40: "Осень",
    }

    Sum(5, 5)
    ReturnInt()
    ReturnFloat()
    ReturnIntArray()
    ReturnIntSlice()
    IntSliceToString([]int{17, 23, 100500})
    MergeSlices([]float32{1.1, 2.1, 3.1}, []int32{4, 5})
    fmt.Println(GetMapValuesSortedByKey(input))
}
