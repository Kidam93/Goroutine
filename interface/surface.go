package interf

import (
    "fmt"
)

type Forme interface {
    Air() float64
    Perimetre() float64
}

type Rectangle struct {
    largeur  float64
    longueur float64
}

func (r Rectangle) Air() float64 {
	return r.largeur * r.longueur
}

func (r Rectangle) Perimetre() float64 {
	return 2 * (r.largeur * r.longueur)
}


func SubMainSurface() {
    var f Forme
    f = Rectangle{5.0, 4.0}
    r := Rectangle{5.0, 4.0} 
    fmt.Println("Type de f :", f)
    fmt.Printf("Valeur de f : %v\n", f)
    fmt.Println("Air du rectangle r :", f.Air())
    fmt.Println("f == r ? ", f == r)
}