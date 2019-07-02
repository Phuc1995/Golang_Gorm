package main

import (
	"entites"
	"fmt"
	"models"
	"time"
)

func main() {
	Demo4()
}

// func Demo1() {
// 	var productModel models.ProductModel
// 	products, _ := productModel.FindAllAPI()
// 	for _, product := range products {
// 		fmt.Println(product)
// 		fmt.Println("*************************")
// 	}
// }

// func Demo2() {
// 	var productModel models.ProductModel
// 	products, err := productModel.FindAllAPI()
// 	if err != nil {
// 		fmt.Println("not find")
// 	} else {
// 		fmt.Println(products.ToString)
// 	}
// }

func Demo4() {
	var productModel models.ProductModel
	result := productModel.Sum1()
	fmt.Println("result: ", result)
}

func Demo5() {
	var productModel models.ProductModel
	product := entites.ProductAPI{
		"Phuc",
		5,
		2,
		true,
		time.Now,
	}
	productModel.Create(&product)
	fmt.Println(product.ToString)
}

func Demo6() {
	var productModel models.ProductModel
	product := productModel.FindById(2)
	product.Name = "def"
	productModel.Update(&product)
	fmt.Println(product.ToString)
}
