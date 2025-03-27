package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := `SELECT * FROM product`
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println((err))
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}
	rows.Close()
	return productList, nil
}

func (p *ProductRepository) CreateProduct(product model.Product) error {
	query := `INSERT INTO product (product_name, price) VALUES ($1, $2)`
	_, err := p.connection.Exec(query, product.Name, product.Price)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (p *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query := `SELECT * FROM product WHERE id=$1`
	row := p.connection.QueryRow(query, id)
	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &product, nil
}
