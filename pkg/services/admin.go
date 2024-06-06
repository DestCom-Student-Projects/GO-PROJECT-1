package services

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/database"
	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/menu"
)

func AddProduct(){
	reader := bufio.NewReader(os.Stdin)

	var productName string
	var productDescription string
	var productPrice string
	var productQuantity string


	menu.ClearScreen()
	fmt.Println("Please enter the name of the product :")
	productName, _ = reader.ReadString('\n')
	productName = strings.TrimSpace(productName)

	fmt.Println("Please enter the description of the product :")
	productDescription, _ = reader.ReadString('\n')
	productDescription = strings.TrimSpace(productDescription)

	fmt.Println("Please enter the price of the product :")
	productPrice, _ = reader.ReadString('\n')
	productPrice = strings.TrimSpace(productPrice)

	fmt.Println("Please enter the quantity of the product :")
	productQuantity, _ = reader.ReadString('\n')
	productQuantity = strings.TrimSpace(productQuantity)
	price, _ := strconv.Atoi(productPrice)
	quantity, _ := strconv.Atoi(productQuantity)

	
	// Add product to database
	database.AddProduct(database.Product{
		Titre: productName,
		Description: productDescription,
		Prix: uint(price),
		Quantite: quantity,
		IsActive: true,
	})

	fmt.Println("Product added successfully")
}

func EditProduct(){
	reader := bufio.NewReader(os.Stdin)

	var loopLock bool = false
	var productID string
	var id int
	var productName string
	var productDescription string
	var productPrice string
	var productQuantity string
	var productAvailability string
	var isAvailable bool

	for !loopLock {
		menu.ClearScreen()
		fmt.Println("Please enter the ID of the product you want to edit :")
		productID, _ = reader.ReadString('\n')
		productID = strings.TrimSpace(productID)
		id , _ = strconv.Atoi(productID)


		//Check if the product exists
		existingProduct := database.GetProductByID(id)
		if existingProduct.ID == 0 {
			fmt.Println("Product not found")
			fmt.Println("Do you want to try again ? (y or n)")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			if choice == "n" {
				loopLock = true
				return
			}
		} else {
			loopLock = true
		}
	}


	fmt.Println("Please enter the name of the product :")
	productName, _ = reader.ReadString('\n')
	productName = strings.TrimSpace(productName)

	fmt.Println("Please enter the description of the product :")
	productDescription, _ = reader.ReadString('\n')
	productDescription = strings.TrimSpace(productDescription)

	fmt.Println("Please enter the price of the product :")
	productPrice, _ = reader.ReadString('\n')
	productPrice = strings.TrimSpace(productPrice)

	fmt.Println("Please enter the quantity of the product :")
	productQuantity, _ = reader.ReadString('\n')
	productQuantity = strings.TrimSpace(productQuantity)

	fmt.Println("Is product available ? (y or n)")
	productAvailability, _ = reader.ReadString('\n')
	productAvailability = strings.TrimSpace(productAvailability)
	if productAvailability == "y" {
		isAvailable = true
	} else {
		isAvailable = false
	}

	price, _ := strconv.Atoi(productPrice)
	quantity, _ := strconv.Atoi(productQuantity)

	// Edit product in database
	database.EditProduct(database.Product{
		ID: id,
		Titre: productName,
		Description: productDescription,
		Prix: uint(price),
		Quantite: quantity,
		IsActive: isAvailable,
	})

	fmt.Println("Product edited successfully")
}

func DeleteProduct(){
	reader := bufio.NewReader(os.Stdin)

	var loopLock bool = false
	var productID string
	var id int
	var existingProduct database.Product

	for !loopLock {
		menu.ClearScreen()
		fmt.Println("Please enter the ID of the product you want to delete :")
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
				loopLock = true
				return
			}
		} else {
			loopLock = true
		}
	}
	
	database.DeleteProduct(database.Product{
		ID: id,
		Titre: existingProduct.Titre,
		Description: existingProduct.Description,
		Prix: existingProduct.Prix,
		Quantite: existingProduct.Quantite,
		IsActive: false,
	})

	fmt.Println("Product deleted successfully")
}

func AddCustomer(){
	reader := bufio.NewReader(os.Stdin)

	var customerName string
	var customerFirstName string
	var customerPhone string
	var customerAddress string
	var customerEmail string

	menu.ClearScreen()
	fmt.Println("Please enter the last name of the customer :")
	customerName, _ = reader.ReadString('\n')
	customerName = strings.TrimSpace(customerName)

	fmt.Println("Please enter the first name of the customer :")
	customerFirstName, _ = reader.ReadString('\n')
	customerFirstName = strings.TrimSpace(customerFirstName)

	fmt.Println("Please enter the phone number of the customer :")
	customerPhone, _ = reader.ReadString('\n')
	customerPhone = strings.TrimSpace(customerPhone)

	fmt.Println("Please enter the address of the customer :")
	customerAddress, _ = reader.ReadString('\n')
	customerAddress = strings.TrimSpace(customerAddress)

	fmt.Println("Please enter the email of the customer :")
	customerEmail, _ = reader.ReadString('\n')
	customerEmail = strings.TrimSpace(customerEmail)


	// Add customer to database
	database.AddCustomer(database.Customer{
		Nom: customerName,
		Prenom: customerFirstName,
		Telephone: customerPhone,
		Adresse: customerAddress,
		Email: customerEmail,
		HasCanceled: false,
	})

	fmt.Println("Customer added successfully")
}

func EditCustomer(){
	reader := bufio.NewReader(os.Stdin)

	var loopLock bool = false
	var customerID string
	var id int
	var customerName string
	var customerFirstName string
	var customerPhone string
	var customerAddress string
	var customerEmail string
	var hasCanceled string
	var hasCanceledBool bool

	for !loopLock {
		menu.ClearScreen()
		fmt.Println("Please enter the ID of the customer you want to edit :")
		customerID, _ = reader.ReadString('\n')
		customerID = strings.TrimSpace(customerID)
		id , _ = strconv.Atoi(customerID)

		//Check if the customer exists
		existingCustomer := database.GetCustomerByID(id)
		if existingCustomer.ID == 0 {
			fmt.Println("Customer not found")
			fmt.Println("Do you want to try again ? (y or n)")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			if choice == "n" {
				loopLock = true
				return
			}
		} else {
			loopLock = true
		}
	}

	fmt.Println("Please enter the last name of the customer :")
	customerName, _ = reader.ReadString('\n')
	customerName = strings.TrimSpace(customerName)

	fmt.Println("Please enter the first name of the customer :")
	customerFirstName, _ = reader.ReadString('\n')
	customerFirstName = strings.TrimSpace(customerFirstName)

	fmt.Println("Please enter the phone number of the customer :")
	customerPhone, _ = reader.ReadString('\n')
	customerPhone = strings.TrimSpace(customerPhone)

	fmt.Println("Please enter the address of the customer :")
	customerAddress, _ = reader.ReadString('\n')
	customerAddress = strings.TrimSpace(customerAddress)

	fmt.Println("Please enter the email of the customer :")
	customerEmail, _ = reader.ReadString('\n')
	customerEmail = strings.TrimSpace(customerEmail)

	fmt.Println("Has the customer canceled his account ? (y or n)")
	hasCanceled, _ = reader.ReadString('\n')
	hasCanceled = strings.TrimSpace(hasCanceled)
	if hasCanceled == "y" {
		hasCanceledBool = true
	} else {
		hasCanceledBool = false
	}

	database.EditCustomer(database.Customer{
		ID: id,
		Nom: customerName,
		Prenom: customerFirstName,
		Telephone: customerPhone,
		Adresse: customerAddress,
		Email: customerEmail,
		HasCanceled: hasCanceledBool,
	})

	fmt.Println("Customer edited successfully")
}


func DeleteCustomer(){
	reader := bufio.NewReader(os.Stdin)

	var loopLock bool = false
	var customerID string
	var id int
	var existingCustomer database.Customer

	for !loopLock {
		menu.ClearScreen()
		fmt.Println("Please enter the ID of the customer you want to delete :")
		customerID, _ = reader.ReadString('\n')
		customerID = strings.TrimSpace(customerID)
		id , _ = strconv.Atoi(customerID)

		//Check if the customer exists
		existingCustomer = database.GetCustomerByID(id)
		if existingCustomer.ID == 0 {
			fmt.Println("Customer not found")
			fmt.Println("Do you want to try again ? (y or n)")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			if choice == "n" {
				loopLock = true
				return
			}
		} else {
			loopLock = true
		}
	}

	database.DeleteCustomer(database.Customer{
		ID: id,
		Nom: existingCustomer.Nom,
		Prenom: existingCustomer.Prenom,
		Telephone: existingCustomer.Telephone,
		Adresse: existingCustomer.Adresse,
		Email: existingCustomer.Email,
		HasCanceled: true,
	})

	fmt.Println("Customer deleted successfully")
}


func ExportServices(){
	menu.DisplayMenuExport()

	inpts := menu.Inputs([]string{"1", "2", "3", "4"})
	var locked bool = false

	for !locked {
		switch inpts {
		case 1:
			ExportProducts()
			locked = true
		case 2:
			ExportCustomers()
			locked = true
		case 3:
			ExportOrders()
			locked = true
		case 4:
			locked = true
		default:
			menu.DisplayErrors("Invalid choice, exiting...")
			return
		}
	}
}

func ExportProducts(){
	products := database.GetProducts(false)

	file, err := os.Create("products.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Titre", "Description", "Prix", "Quantite", "IsActive"})
	for _, product := range products {
		err := writer.Write([]string{strconv.Itoa(product.ID), product.Titre, product.Description, strconv.Itoa(int(product.Prix)), strconv.Itoa(product.Quantite), strconv.FormatBool(product.IsActive)})

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ExportCustomers(){
	customers := database.GetCustomers()

	file, err := os.Create("customers.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Nom", "Prenom", "Telephone", "Adresse", "Email", "HasCanceled"})
	for _, customer := range customers {
		err := writer.Write([]string{strconv.Itoa(customer.ID), customer.Nom, customer.Prenom, customer.Telephone, customer.Adresse, customer.Email, strconv.FormatBool(customer.HasCanceled)})

		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func ExportOrders(){
	orders := database.GetOrders()

	file, err := os.Create("orders.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "ClientID", "ProduitID", "Quantite", "Prix", "DateAchat"})
	for _, order := range orders {
		err := writer.Write([]string{strconv.Itoa(order.ID), strconv.Itoa(order.IDClient), strconv.Itoa(order.IDProduit), strconv.Itoa(order.Quantite), strconv.Itoa(int(order.Prix)), order.DateAchat.String()})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ListingsServices(){
	menu.DisplayMenuListings()

	inpts := menu.Inputs([]string{"1", "2", "3", "4"})
	var locked bool = false

	for !locked {
		switch inpts {
		case 1:
			ListProducts()

			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

			locked = true
		case 2:
			ListCustomers()

			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

			locked = true
		case 3:
			ListOrders()

			fmt.Println("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			
			locked = true
		case 4:
			locked = true
		default:
			menu.DisplayErrors("Invalid choice, exiting...")
			return
		}
	}
}

func ListProducts(){
	products := database.GetProducts(true)

	menu.DisplayMenuProducts(products)
}

func ListCustomers(){
	customers := database.GetCustomers()

	menu.ClearScreen()
	fmt.Println("List of customers:")
	for _, customer := range customers {
		fmt.Printf("%d. %s %s\n", customer.ID, customer.Prenom, customer.Nom)
	}
}

func ListOrders(){
	orders := database.GetOrders()

	menu.ClearScreen()
	fmt.Println("List of orders:")
	for _, order := range orders {
		fmt.Printf("%d. %d %d\n", order.ID, order.IDClient, order.IDProduit)
	}
}
