package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ahabune"
	dbname   = "golangtestdb"
)

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float32
}

var db *sql.DB

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO product(title, description, price) VALUES($1, $2, $3)", data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}

	lastInsertedId, _ := result.LastInsertId()
	fmt.Printf("Eklenen kayıt id'si (%d)", lastInsertedId)
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE product SET title=$2, description=$3, price=$4 WHERE id=$1", data.ID, data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}

	rowAffected, err := result.RowsAffected()
	fmt.Printf("Etkilenen Kayıt Sayısı (%d)", rowAffected)
}

func GetProducts() {
	rows, err := db.Query("SELECT * FROM product")
	defer rows.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Reocords found!")
			return
		}
		log.Fatal(err)
	}

	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.ID, &prd.Title, &prd.Description, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range products {
		fmt.Printf("%d - %s, %s, $%.2f\n", v.ID, v.Title, v.Description, v.Price)
	}
}

func GetProductByID(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM product WHERE ID = $1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.\n")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product is %s\n", product)
	}
}
