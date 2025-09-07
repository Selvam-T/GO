package main
import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// mutation will persist outside function scope
func (r *rect) scale1(factor int) {
	r.width *= factor 
}

// mutation will not persist outside function scope
func (r rect) scale2(factor int) {
	r.width *= factor
}

func main() {
	r := rect{width: 10, height: 2}

	// value receiver type
	fmt.Println("area: ",r.area())
	fmt.Println("perim: ",r.perim())

	// pointer receiver type
	rp := &r
	fmt.Println("area: ",rp.area())
        fmt.Println("perim: ",rp.perim())

	//mutation of struct value
	fmt.Println("width before scale1(): ", r.width)
	r.scale1(2)
	fmt.Println("width after scale1() by 2: ", r.width, " <modified>") // modified
	r.scale2(5)
	fmt.Println("width after scale2() by 5: ", r.width, " <not modified>") //not modified

	fmt.Println("width before scale2(): ", rp.width)
        r.scale1(2)
        fmt.Println("width after scale2() by 2: ", rp.width, " <modified>") // modified
        r.scale2(5)
        fmt.Println("width after scale2() by 5: ", rp.width, " <not modified>") // modified

}
