package repo

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgURl      string
}

type ProductRepo interface {
	CreateProduct(p Product) (*Product, error)
	GetProducts() (*[]Product, error)
	GetProductById(id int) (*Product, error)
	UpdateProduct(id int, p Product) (*Product, error)
	DeleteProduct(id int) (*Product, error)
}
type productRepo struct {
	productList []Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateProduct(repo)
	return repo
}

func (r *productRepo) CreateProduct(p Product) (*Product, error) {
	r.productList = append(r.productList, p)
	return &p, nil
}
func (r *productRepo) GetProducts() (*[]Product, error) {
	return &r.productList, nil
}

func (r *productRepo) GetProductById(id int) (*Product, error) {
	for _, p := range r.productList {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, nil
}
func (r *productRepo) UpdateProduct(id int, p Product) (*Product, error) {
	for i, v := range r.productList {
		if v.ID == id {
			r.productList[i] = p
			return &p, nil
		}
	}
	return nil, nil
}
func (r *productRepo) DeleteProduct(id int) (*Product, error) {
	var temp []Product
	for _, v := range r.productList {
		if v.ID != id {
			temp = append(temp, v)
		}
	}
	r.productList = temp
	return nil, nil
}

func generateProduct(r *productRepo) {
	r.productList = []Product{
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
