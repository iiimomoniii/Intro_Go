package main

import (
	"fmt"
);
func main() {
	assignValue();
}

// Struct = Class
type product struct { // Product share to other package
	name  string
	price int
	stock int
}

func assignValue() {
	var p1 product
	p1.name = "Arduino"
	p1.stock = 20
	p1.price = 100
	viewData(p1);
	updateProduct(&p1);
	viewData(p1);

	p1 = p1.discountProduct(5);
	viewData(p1);

	p1 = p1.clearProduct();
	viewData(p1);
}

func viewData(p product) {
	fmt.Println(p);
}

//Update Value
func updateProduct(p *product){
	p.price = p.price + 100;
	p.stock = 10;
	fmt.Println(*p);
}

//Discount Value
func (p product) discountProduct(d int) product{
	p.price = p.price - d;
	return p
}

//Clear Value
func (p product) clearProduct() product{
	p.price = 0
	p.stock = 0
	return p
}