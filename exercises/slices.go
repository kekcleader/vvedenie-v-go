
s := []int{10, 20, 30, 40, 50}
s1 := s[:3]
fmt.Println(len(s1), cap(s1))
s2 := s1[:4]
fmt.Println(s2)
s2[3] = 99
fmt.Println("s:", s)
fmt.Println("s1:", s1)
fmt.Println("s2:", s2)


arr := []int{1, 2, 3, 4, 5}
s1 := arr[1:4]
s2 := arr[2:5]
s1[1] = 100
fmt.Println("s1:", s1)
fmt.Println("s2:", s2)
fmt.Println("arr:", arr)


orig := []int{100, 200, 300, 400}
sub := orig[:2]
extended := append(sub, 500)
fmt.Println("orig:", orig)
fmt.Println("sub:", sub)
fmt.Println("extended:", extended)


s := []int{5}
s = append(s, 7)
s = append(s, 9)
x := append(s, 11)
y := append(s, 12)
fmt.Println(s, x, y)


  * * *


func modifyAndAppend(s []int) []int {
    s[0] = 99
    return append(s, 100)
}

func main() {
    arr := make([]int, 2, 3)
    arr[0], arr[1] = 1, 2
    arr = modifyAndAppend(arr)
    fmt.Println(arr, len(arr), cap(arr))
}



var s1 []int
s2 := []int{}
fmt.Println(len(s1), cap(s1), s1 == nil)
fmt.Println(len(s2), cap(s2), s2 == nil)


big := make([]int, 1000)
small := big[10:20]
big = nil
runtime.GC()
fmt.Println(len(small), cap(small))


data := append([]int(nil), small...)




