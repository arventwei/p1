package main

import "fmt"

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	w1
	w2
	w3
)

type Point struct {
	X int
	Y int
}

type Circle struct {
	radius   int
	position Point
}

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	b1      = "b1"
	b2      = "b2"
	b3      = "b3"
	b4      = 1
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func ArrChange(arr []int) {
	arr[0] = 100
}

type tesetInt func(int2 int) bool

func greatInt(i int) bool {
	return i > 100
}
func main() {

	var t1 tesetInt
	t1 = greatInt
	fmt.Println(t1(3))
	c1 := make([]Circle, 3)
	c2 := new([3]Circle)

	c1[0].position.X = 1
	c2[0].position.X = 1
	fmt.Println(c1)
	fmt.Println(c2)

	//fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
	//fmt.Println(h, i, j)
	//fmt.Println(w1, w2, w3)
	//fmt.Println(b1, b2, b3)

	//am := make(map[int]string)
	//am[1] = "asdfsdf"
	//arr := []int{1, 2, 3, 4, 5, 6, 7}
	//arrb := []int{100, 200, 300}
	//arrc := append(arr, arrb...)
	//ArrChange(arr[:])
	//fmt.Println(arr, arrb, arrc)
	//fmt.Println(am, am[4])
	//fmt.Println(am[5])
	FireMp5()
}
