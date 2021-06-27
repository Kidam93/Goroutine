package interf

import (
	"fmt"
)

type Instrumenter interface {
	Play()
}

type Guitar struct {
	son string
}

type Piano struct {
	son string
}

func (g Guitar) Play() {
	g.son = "tzouing"
	fmt.Printf("je de la guitar %v", g.son)
}	

func (p Piano) Play() {
	p.son = "plip plip"
	fmt.Printf("je joue du piano debout %v", p.son)
}

func SubMainInterf() {
	var p Instrumenter
	p = Guitar{}
	p.Play()
	p = Piano{}
	p.Play()
}