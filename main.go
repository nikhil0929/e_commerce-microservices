package main

func main() {
	product_dao = products.NewProductDao()
	product_service = products.Service(product_dao)
	
}