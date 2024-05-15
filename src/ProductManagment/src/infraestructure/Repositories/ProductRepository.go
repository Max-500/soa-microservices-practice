package repositories

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"order-managment/src/domain/entities"
	configMongo "order-managment/src/infraestructure/Database/MongoDB/Config"
	configMySQL "order-managment/src/infraestructure/Database/MySQL/Config"
)

type ProductRepository struct {
	dbType  string
	dbMySQL *configMySQL.MySQL
	dbMongo *configMongo.MongoDB
}

func NewProductRepository(dbType string, mysql *configMySQL.MySQL, mongo *configMongo.MongoDB) *ProductRepository {
	return &ProductRepository{
		dbType:  dbType,
		dbMySQL: mysql,
		dbMongo: mongo,
	}
}

func (repo *ProductRepository) Create(data []entities.Product) ([]entities.Product, error) {
	var insertedProducts []entities.Product

	if repo.dbType == "MongoDB" {
		collection := repo.dbMongo.Client.Database("mydatabase").Collection("products")
		for _, product := range data {
			product.Uuid = uuid.New().String()
			_, err := collection.InsertOne(context.TODO(), bson.M{
				"uuid":  product.Uuid,
				"name":  product.Name,
				"price": product.Price,
				"stock": product.Stock,
			})
			if err != nil {
				return nil, err
			}
			insertedProducts = append(insertedProducts, product)
		}
	}

	if repo.dbType == "MySQL" {
		db := repo.dbMySQL.DB
		for _, product := range data {
			product.Uuid = uuid.New().String()
			query := `INSERT INTO products (uuid, name, price, stock) VALUES (?, ?, ?, ?)`
			_, err := db.Exec(query, product.Uuid, product.Name, product.Price, product.Stock)
			if err != nil {
				return nil, err
			}
			insertedProducts = append(insertedProducts, product)
		}
	}
	return insertedProducts, nil
}

func (repo *ProductRepository) Delete(id uuid.UUID) (string, error) {
	return "", nil
}

func (repo *ProductRepository) UpdateTracking(id uuid.UUID) (string, error) {
	return "", nil
}

func (repo *ProductRepository) GetAllProducts() ([]entities.Product, error) {
	return []entities.Product{}, nil
}
