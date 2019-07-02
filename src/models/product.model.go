package models

import (
	"config"
	"entites"
)

type ProductModel struct {
}

func (productModel ProductModel) FindAll() ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Condition1(status bool) ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Where("status = ?", status).Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Condition2(min, max float64) ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Where("price >= ? and price <= ?", min, max).Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Like(keyword string) ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Where("name like ?", "%"+keyword+"%").Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Sort1() ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Order("price desc").Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Sort2(status bool) ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Where("status = ?", status).Order("price desc").Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Limit1(n int) ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Limit(n).Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Limit2(n int) ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Order("id desc").Limit(n).Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) Limit3(status bool, n int) ([]entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var product []entites.Product
		db.Where("status = ?", status).Order("id desc").Limit(n).Find(&product)
		return product, nil
	}
}

func (productModel ProductModel) FindById(id int) (entites.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return entites.Product{}, err
	} else {
		var product entites.Product
		db.Where("id = ?", id).First(&product)
		return product, nil
	}
}

func (productModel ProductModel) FindAllAPI() ([]entites.ProductAPI, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var productAPIs []entites.ProductAPI
		db.Table("product").Select("id, name, price").Scan(&productAPIs)
		return productAPIs, nil
	}
}

// func (productModel ProductModel) Count1() int64 {
// 	db, err := config.GetDB()
// 	var result int64
// 	db.Table("product").Count(&result)
// 	return result
// }

// func (productModel ProductModel) Count2(status bool) int64 {
// 	db, err := config.GetDB()
// 	var result int64
// 	db.Table("product").Where("status = ?", status)Count(&result)
// 	return result
// }

func (productModel ProductModel) Sum1() int64 {
	db, _ := config.GetDB()
	var total int64
	rows, _ := db.Table("product").Select("sum(quantity) as total").Rows()
	rows.Next()
	rows.Scan(&total)
	return total
}

// func (productModel ProductModel) MaxAndMinPrice() (float64 float64) {
// 	db, _ := config.GetDB()
// 	var min, max float64
// 	rows, _ := db.Table("product").Select("min(price) as minPrice, max(price) as maxPrice").Rows()
// 	rows.Next()
// 	rows.Scan(&min, &max)
// 	return min, max
// }

func (productModel ProductModel) Create(product *entites.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&product)
		return nil
	}
}

func (productModel ProductModel) Update(product *entites.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		db.Update(&product)
		return nil
	}
}

func (productModel ProductModel) Delete(product *entites.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(&product)
		return nil
	}
}
