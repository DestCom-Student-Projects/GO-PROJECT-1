package menu

import (
	"fmt"

	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/database"
)


func DisplayBase(){
	ClearScreen()	
	fmt.Println("Welcome to the grocery-store.online")
	fmt.Println("Choose an action below:")
}

func DisplayMenuOne(doesClearScreen bool){
	if doesClearScreen {
		ClearScreen()
	}

	fmt.Println("1. For customers")
	fmt.Println("2. For admins")
	fmt.Println("3. Exit")
}

func DisplayMenuLogin(){
	ClearScreen()
	fmt.Println("Please enter your email linked to your account")
}

func DisplayMenuCustomers(){
	ClearScreen()
	fmt.Println("1. Display products")
	fmt.Println("2. Buy product")
	fmt.Println("3. Exit")
}

func DisplayMenuProducts(products []database.Product){
	ClearScreen()
	fmt.Println("List of products:")
	for _, product := range products {
		fmt.Printf("%d. %s\n", product.ID, product.Titre)
	}
}

func DisplayMenuAdmins(){
	ClearScreen()
	fmt.Println("Welcome admin, please choose an action below:")
	fmt.Println("1. Add product")
	fmt.Println("2. Edit product")
	fmt.Println("3. Delete product")
	fmt.Println("4. Add customer")
	fmt.Println("5. Edit customer")
	fmt.Println("6. Delete customer")
	fmt.Println("7. Export menu")
	fmt.Println("8. Listings menu")
	fmt.Println("9. Exit")
}

func DisplayMenuListings(){
	ClearScreen()
	fmt.Println("1. List products")
	fmt.Println("2. List customers")
	fmt.Println("3. List orders")
	fmt.Println("4. Go back")
}

func DisplayMenuExport(){
	ClearScreen()
	fmt.Println("1. Export products")
	fmt.Println("2. Export customers")
	fmt.Println("3. Export orders")
	fmt.Println("4. Go back")
}

func DisplayErrors(err string){
	fmt.Println(err)
}

func DisplayGoodbye(){	
	fmt.Println("Goodbye")
}

func ClearScreen(){
	fmt.Println("\033[H\033[2J")
}

//AZEVEDO-DA-SILVA, A. (2024) GO-PROJECT-1. [Source code] https://github.com/DestCom-Student-Projects/GO-PROJECT-1