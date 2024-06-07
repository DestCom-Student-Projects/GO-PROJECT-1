package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/database"
	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/mail"
	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/menu"
	pdfBuilder "github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/pdf"
)


func DisplayProducts() {
    products := database.GetProducts(true)

    menu.DisplayMenuProducts(products)

    fmt.Println("Press any key to continue...")
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
}

func BuyProduct() {
	reader := bufio.NewReader(os.Stdin)
	var isLocked bool = false

	for !isLocked {
		var secondLocked bool = false
		var thirdLocked bool = false
		var fourthLocked bool = false
		var productID string
		var id int
		var existingCustomer database.Customer
		var existingProduct database.Product
		var quantityInt int

		for !secondLocked {
			menu.DisplayMenuProducts(database.GetProducts(true))
			fmt.Println("Enter the ID of the product you want to buy")
			fmt.Println("Press 0 to go back")
			productID, _ = reader.ReadString('\n')
			productID = strings.TrimSpace(productID)
			id , _ = strconv.Atoi(productID)


			//Check if the product exists
			existingProduct = database.GetProductByID(id)
			if existingProduct.ID == 0 {
				fmt.Println("Product not found")
				fmt.Println("Do you want to try again ? (y or n)")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)
				if choice == "n" {
					secondLocked = true
					return
				}
			} else {
				secondLocked = true
			}

			if productID == "0" {
				secondLocked = true
				return
			}
		}

		for !thirdLocked {
			fmt.Println("Enter your email")
			email, _, _ := reader.ReadLine()

			//Check if the email exists
			existingCustomer = database.GetCustomerByEmail(string(email))
			if existingCustomer.Email == "" {
				fmt.Println("Email not found")
				fmt.Println("Do you want to try again ? (y or n)")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)
				if choice == "n" {
					thirdLocked = true
				}
			} else {
				thirdLocked = true
			}
		}

		for !fourthLocked {
			fmt.Println("Enter the quantity")
			quantityBytes, _, _ := reader.ReadLine()
			quantity := string(quantityBytes)
			quantity = strings.TrimSpace(quantity)
			quantityInt, _ = strconv.Atoi(quantity)

			if quantityInt <= 0 || quantityInt > existingProduct.Quantite {
				fmt.Println("The maximum quantity you can buy is", existingProduct.Quantite)
				fmt.Println("Do you want to try again ? (y or n)")
				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)
				if choice == "n" {
					fourthLocked = true
				}
			} else {
				fourthLocked = true
			}
		}


		ProceedToBuy(existingProduct, quantityInt, existingCustomer)

		fmt.Println("Product bought successfully")
		fmt.Println("Press any key to continue...")
		reader.ReadString('\n')
		isLocked = true
	}
}

func ProceedToBuy(product database.Product, quantity int, customer database.Customer) {
    buyed, _ := database.BuyProduct(product, quantity, customer)

	subject := fmt.Sprintf("Order confirmation #%d-%d", customer.ID, buyed.ID)
    body := fmt.Sprintf("Your order has been confirmed. You bought %d %s. For reference, your order ID is %d", quantity, product.Titre, buyed.ID)
	pdfBuilder.GeneratePDF(product, quantity, customer, buyed)
	filepath := fmt.Sprintf("orders_pdf/order-%d-%d.pdf", customer.ID, buyed.ID)

	mail.SendMail(customer.Email, subject, body, filepath)
}

//AZEVEDO-DA-SILVA, A. (2024) GO-PROJECT-1. [Source code] https://github.com/DestCom-Student-Projects/GO-PROJECT-1