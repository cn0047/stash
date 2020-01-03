package model

type Product struct {
	Name string
	CategoryID int
}

var products []Product = []Product{
	Product{
		Name: "BMW X5",
		CategoryID: 1,
	},
}

func GetProducts() []Product {
	return products
}

func GetProductsForCategory(categoryID int) []Product {
	result := []Product{}
	for _, p := range products {
		if p.CategoryID == categoryID {
			result = append(result, p)
		}
	}
	return result
}
