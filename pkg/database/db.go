package database

import (
	"fmt"
	"time"

	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/environnement"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
    ID          int
    Titre       string
    Description string
    Prix        uint
    Quantite    int
    IsActive   bool
}

type Customer struct {
    ID        int
    Nom       string
    Prenom    string
    Telephone string
    Adresse   string
    Email     string
    HasCanceled bool
}

type Order struct {
    ID        int
    IDClient  int
    IDProduit int
    Quantite  int
    Prix      uint
    DateAchat time.Time
}

var dbUser = environnement.GetEnv("DB_USER")
var dbPassword = environnement.GetEnv("DB_PASSWORD")
var dbHost = environnement.GetEnv("DB_HOST")
var dbPort = environnement.GetEnv("DB_PORT")
var dbName = environnement.GetEnv("DB_NAME")
var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

func init() {
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Customer{}, &Order{})
}


func AddProduct(product Product) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    db.Create(&product)
}

func EditProduct(product Product) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    db.Save(&product)

}

func DeleteProduct(product Product) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    db.Save(&product)
}

func GetProductByID(productID int) Product{
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    var product Product
    db.First(&product, productID)

    return product
}

func GetProducts(onlyActive bool) []Product {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    var products []Product

    if onlyActive {
        db.Where("is_active = ?", true).Find(&products)
    } else {
        db.Find(&products)
    }

    return products
}


func AddCustomer(customer Customer) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    db.Create(&customer)
}

func EditCustomer(customer Customer) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    db.Save(&customer)
}

func DeleteCustomer(customer Customer) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    db.Save(&customer)
}

func GetCustomerByID(customerID int) Customer{
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    var customer Customer
    db.First(&customer, customerID)

    return customer
}

func GetCustomers() []Customer {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    var customers []Customer
    db.Find(&customers)

    return customers
}


func GetCustomerByEmail(email string) Customer{
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    var customer Customer
    db.Where("email = ?", email).First(&customer)

    return customer
}

func BuyProduct(product Product, quantity int, customer Customer) (Order, error) {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        return Order{}, fmt.Errorf("failed to connect database: %w", err)
    }

    db.Save(&Product{
        ID: product.ID,
        Titre: product.Titre,
        Description: product.Description,
        Prix: product.Prix,
        Quantite: product.Quantite - quantity,
        IsActive: product.IsActive,
    })

    order := Order{
        IDClient: customer.ID,
        IDProduit: product.ID,
        Quantite: quantity,
        Prix: product.Prix,
        DateAchat: time.Now(),
    }

    result := db.Create(&order)

    if result.Error != nil {
        return Order{}, fmt.Errorf("failed to create order: %w", result.Error)
    }

    return order, nil
}

func GetOrders() []Order{
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }

    var orders []Order
    db.Find(&orders)

    return orders
}

//AZEVEDO-DA-SILVA, A. (2024) GO-PROJECT-1. [Source code] https://github.com/DestCom-Student-Projects/GO-PROJECT-1