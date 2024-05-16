package main

import (
	"fmt"
	"math"
)

// 利用接口方法计算矩形和圆的周长和面积
type methods interface {
	// Perimeter 周长
	Perimeter() float64
	// Area 面积
	Area() float64
}

// 矩形
type rectangle struct {
	//长
	length float64
	//宽
	width float64
}

// 圆形
type circle struct {
	//半径
	radius float64
}

// Perimeter 矩形计算周长
func (r *rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Area 矩形计算面积
func (r *rectangle) Area() float64 {
	return r.width * r.length
}

// Perimeter 圆计算周长
func (r *circle) Perimeter() float64 {
	return math.Pi * r.radius * 2
}

// Area 圆计算面积
func (r *circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}

// Calculate 计算函数
func Calculate(r methods) (float64, float64) {
	p := r.Perimeter()
	a := r.Area()
	return p, a
}

func main() {
	r := &rectangle{
		length: 3,
		width:  5,
	}

	c := &circle{
		radius: 2,
	}

	p1, a1 := Calculate(r)
	p2, a2 := Calculate(c)
	fmt.Printf("矩形的周长为%s;矩形的面积为%s\n", fmt.Sprintf("%.2f", p1), fmt.Sprintf("%.2f", a1))
	fmt.Printf("圆的周长为%s;圆的面积为%s", fmt.Sprintf("%.2f", p2), fmt.Sprintf("%.2f", a2))
}
