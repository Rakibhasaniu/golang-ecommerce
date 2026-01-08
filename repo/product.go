package repo

import (
	"database/sql"
	"main/domain"
	"main/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}
type productRepo struct {
	dbCon *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) ProductRepo {
	repo := &productRepo{
		dbCon: dbCon,
	}
	return repo
}

func (r *productRepo) CreateProduct(p domain.Product) (*domain.Product, error) {
	query := `INSERT INTO products (title, description, price, img_url) VALUES ($1, $2, $3, $4) RETURNING id`
	row := r.dbCon.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *productRepo) GetProducts(page int, limit int) ([]domain.Product, error) {
	var productList []domain.Product
	query := `SELECT * FROM products LIMIT $1 OFFSET $2`
	err := r.dbCon.Select(&productList, query, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	return productList, nil
}

func (r *productRepo) CountProducts() (int, error) {
	query := `SELECT COUNT(*) FROM products`
	var count int
	err := r.dbCon.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *productRepo) GetProductById(id int) (*domain.Product, error) {
	query := `SELECT * FROM products WHERE id = $1`
	var product domain.Product
	err := r.dbCon.Get(&product, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
func (r *productRepo) UpdateProduct(p domain.Product) (*domain.Product, error) {
	query := `UPDATE products SET title = $1, description = $2, price = $3, img_url = $4 WHERE id = $5 RETURNING id`
	row := r.dbCon.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, p.ID)
	err := row.Scan(&p.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}
func (r *productRepo) DeleteProduct(id int) (*domain.Product, error) {
	query := `DELETE FROM products WHERE id = $1 RETURNING id`
	_, err := r.dbCon.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// func generateProduct(r *productRepo) {
// 	r.productList = []Product{
// 		{
// 			ID:          1,
// 			Title:       "Orange",
// 			Description: "Fresh organic oranges from local farms",
// 			Price:       2.99,
// 			ImgURl:      "https://example.com/orange.jpg",
// 		},
// 		{
// 			ID:          2,
// 			Title:       "Apple",
// 			Description: "Crisp red apples, perfect for snacking",
// 			Price:       3.49,
// 			ImgURl:      "https://example.com/apple.jpg",
// 		},
// 		{
// 			ID:          3,
// 			Title:       "Banana",
// 			Description: "Ripe yellow bananas, rich in potassium",
// 			Price:       1.99,
// 			ImgURl:      "https://example.com/banana.jpg",
// 		},
// 		{
// 			ID:          4,
// 			Title:       "Grapes",
// 			Description: "Sweet seedless grapes",
// 			Price:       4.99,
// 			ImgURl:      "https://example.com/grapes.jpg",
// 		},
// 		{
// 			ID:          5,
// 			Title:       "Mango",
// 			Description: "Tropical mangoes, sweet and juicy",
// 			Price:       5.99,
// 			ImgURl:      "https://example.com/mango.jpg",
// 		},
// 		{
// 			ID:          6,
// 			Title:       "Watermelon",
// 			Description: "Large fresh watermelon, perfect for summer",
// 			Price:       7.99,
// 			ImgURl:      "https://example.com/watermelon.jpg",
// 		},
// 		{
// 			ID:          7,
// 			Title:       "Strawberry",
// 			Description: "Fresh sweet strawberries",
// 			Price:       6.49,
// 			ImgURl:      "https://example.com/strawberry.jpg",
// 		},
// 	}
// }
