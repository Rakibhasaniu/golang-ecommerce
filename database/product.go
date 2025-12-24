package database

var productList []Product

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgURl      string
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {

	return productList
}

func ListById(id int) *Product {
	for _, p := range productList {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func Update(id int, p Product) *Product {
	for i, v := range productList {
		if v.ID == id {
			productList[i] = p
			return &p
		}
	}
	return nil
}

func Delete(id int) {
	var temp []Product
	for _, v := range productList {
		if v.ID != id {
			temp = append(temp, v)
		}
	}
	productList = temp
}

func init() {
	productList = []Product{
		{
			ID:          1,
			Title:       "Orange",
			Description: "Fresh organic oranges from local farms",
			Price:       2.99,
			ImgURl:      "https://example.com/orange.jpg",
		},
		{
			ID:          2,
			Title:       "Apple",
			Description: "Crisp red apples, perfect for snacking",
			Price:       3.49,
			ImgURl:      "https://example.com/apple.jpg",
		},
		{
			ID:          3,
			Title:       "Banana",
			Description: "Ripe yellow bananas, rich in potassium",
			Price:       1.99,
			ImgURl:      "https://example.com/banana.jpg",
		},
		{
			ID:          4,
			Title:       "Grapes",
			Description: "Sweet seedless grapes",
			Price:       4.99,
			ImgURl:      "https://example.com/grapes.jpg",
		},
		{
			ID:          5,
			Title:       "Mango",
			Description: "Tropical mangoes, sweet and juicy",
			Price:       5.99,
			ImgURl:      "https://example.com/mango.jpg",
		},
		{
			ID:          6,
			Title:       "Watermelon",
			Description: "Large fresh watermelon, perfect for summer",
			Price:       7.99,
			ImgURl:      "https://example.com/watermelon.jpg",
		},
		{
			ID:          7,
			Title:       "Strawberry",
			Description: "Fresh sweet strawberries",
			Price:       6.49,
			ImgURl:      "https://example.com/strawberry.jpg",
		},
	}
}
