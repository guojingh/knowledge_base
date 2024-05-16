package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {

	r := &rectangle{
		length: 5,
		width:  2,
	}

	c := &circle{
		radius: 4,
	}

	p1, a1 := Calculate(r)
	p2, a2 := Calculate(c)

	t.Logf("矩形的周长为%s;矩形的面积为%s\n", fmt.Sprintf("%.2f", p1), fmt.Sprintf("%.2f", a1))
	t.Logf("圆的周长为%s;圆的面积为%s", fmt.Sprintf("%.2f", p2), fmt.Sprintf("%.2f", a2))
}
