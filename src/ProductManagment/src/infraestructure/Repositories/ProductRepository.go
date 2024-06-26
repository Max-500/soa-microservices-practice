package repositories

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"order-managment/src/domain/entities"
	requests "order-managment/src/infraestructure/Controllers/DTO/Requests"
	configMongo "order-managment/src/infraestructure/Database/MongoDB/Config"
	configMySQL "order-managment/src/infraestructure/Database/MySQL/Config"
)

type ProductRepository struct {
	dbType  string
	dbMySQL *configMySQL.MySQL
	dbMongo *configMongo.MongoDB
	dbName  string
}

func NewProductRepository(dbType string, mysql *configMySQL.MySQL, mongo *configMongo.MongoDB, dbName string) *ProductRepository {
	return &ProductRepository{
		dbType:  dbType,
		dbMySQL: mysql,
		dbMongo: mongo,
		dbName:  dbName,
	}
}

func (repo *ProductRepository) Create(data []entities.Product) ([]entities.Product, error) {
	var insertedProducts []entities.Product

	if repo.dbType == "MongoDB" {
		collection := repo.dbMongo.Client.Database(repo.dbName).Collection("products")
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

func (repo *ProductRepository) Delete(uuids []requests.DeleteProductRequest) (string, error) {
	var notDeleted []string

	if repo.dbType == "MongoDB" {
		collection := repo.dbMongo.Client.Database(repo.dbName).Collection("products")
		for _, uuid := range uuids {
			filter := bson.M{"uuid": uuid.Product}
			res, err := collection.DeleteOne(context.TODO(), filter)
			if err != nil {
				return "", err
			}
			if res.DeletedCount == 0 {
				notDeleted = append(notDeleted, uuid.Product)
			}
		}
	}

	if repo.dbType == "MySQL" {
		db := repo.dbMySQL.DB
		stmt, err := db.Prepare("DELETE FROM products WHERE uuid = ?")
		if err != nil {
			return "", err
		}
		defer stmt.Close()

		for _, uuid := range uuids {
			res, err := stmt.Exec(uuid.Product)
			if err != nil {
				return "", err
			}
			rowCnt, err := res.RowsAffected()
			if err != nil {
				return "", err
			}
			if rowCnt == 0 {
				notDeleted = append(notDeleted, uuid.Product)
			}
		}
	}

	if len(notDeleted) > 0 {
		return fmt.Sprintf("No se pudieron eliminar los productos con los siguientes UUIDs: %v", notDeleted), nil
	}

	return "Todos los productos se eliminaron correctamente", nil
}

func (repo *ProductRepository) UpdateTracking(uuid string, amount int) (string, error) {
	if repo.dbType == "MongoDB" {
		collection := repo.dbMongo.Client.Database(repo.dbName).Collection("products")
		filter := bson.M{"uuid": uuid, "stock": bson.M{"$gte": amount}}
		update := bson.M{"$inc": bson.M{"stock": -amount}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return "", err
		}
	}

	if repo.dbType == "MySQL" {
		db := repo.dbMySQL.DB
		query := `UPDATE products SET stock = stock - ? WHERE uuid = ? AND stock >= 1`
		_, err := db.Exec(query, amount, uuid)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

func (repo *ProductRepository) GetAllProducts() ([]entities.Product, error) {
	var products []entities.Product

	if repo.dbType == "MySQL" {
		rows, err := repo.dbMySQL.DB.Query("SELECT * FROM products")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var product entities.Product
			err := rows.Scan(&product.Uuid, &product.Name, &product.Price, &product.Stock) // Asegúrate de que estos campos coincidan con los de tu estructura y tabla de productos.
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	if repo.dbType == "MongoDB" {
		collection := repo.dbMongo.Client.Database(repo.dbName).Collection("products")
		cursor, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			log.Fatalf("Error al buscar en MongoDB: %v", err)
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var product entities.Product
			err := cursor.Decode(&product)
			if err != nil {
				log.Fatalf("Error al decodificar el producto: %v", err)
			}
			products = append(products, product)
		}

		if err := cursor.Err(); err != nil {
			log.Fatalf("Error en el cursor: %v", err)
		}
	}

	return products, nil
}

func (repo *ProductRepository) GetProduct(uuid string) ([]entities.Product, error) {
	var products []entities.Product

	if repo.dbType == "MongoDB" {
		collection := repo.dbMongo.Client.Database(repo.dbName).Collection("products")
		filter := bson.M{"uuid": uuid}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {
			var product entities.Product
			err := cursor.Decode(&product)
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}

	if repo.dbType == "MySQL" {
		db := repo.dbMySQL.DB
		query := `SELECT * FROM products WHERE uuid = ?`
		rows, err := db.Query(query, uuid)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var product entities.Product
			err := rows.Scan(&product.Uuid, &product.Name, &product.Price, &product.Stock) // Asegúrate de que estos campos coincidan con los de tu estructura y tabla de productos.
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}
	}

	return products, nil
}

