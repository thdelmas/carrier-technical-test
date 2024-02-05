package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-memdb" // Import the correct memdb package
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var db *memdb.MemDB // Declare the database variable

func fillDB() {
	// Create a write transaction
	txn := db.Txn(true)

	// Insert some products
	products := []*Product{
		&Product{1, "Apple", 3},
		&Product{2, "Orange", 5},
		&Product{3, "Banana", 1},
		&Product{4, "Smoothie", 351},
	}
	for _, p := range products {
		if err := txn.Insert("product", p); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()
}

func getAllProducts(c *gin.Context) {
	// Create read-only transaction
	txn := db.Txn(false)
	defer txn.Abort()

	// List all products
	it, err := txn.Get("product", "id")
	if err != nil {
		panic(err)
	}

	var products []*Product
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Product)
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products) // Return products as JSON
}

func main() {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"product": &memdb.TableSchema{
				Name: "product",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	// Create a new data base
	newDB, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	db = newDB // Assign the newDB to the global db variable

	fillDB()
	router := gin.Default()

	router.GET("/api/products", getAllProducts)

	err = router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}

