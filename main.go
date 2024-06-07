package main

import (
	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/menu"
	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/services"
)

func main() {

	var locked bool = false

 	if !locked {
		menu.DisplayBase()
		menu.DisplayMenuOne(false)

		inpts := menu.Inputs([]string{"1", "2", "3"})

		if inpts == 1 {
			var lockedLoopOne bool = false
			var choice int

			for !lockedLoopOne {
				menu.DisplayMenuCustomers()
				choice = menu.Inputs([]string{"1", "2", "3"})
				switch choice {
				case 1:
					services.DisplayProducts()
				case 2:
					services.BuyProduct()
				case 3:
					lockedLoopOne = true
				default:
					menu.DisplayErrors("Invalid choice, exiting...")
					return
				}
			}
		} else if inpts == 2 {
			var lockedLoopTwo bool = false
			var choice int

			for !lockedLoopTwo {
				menu.DisplayMenuAdmins()
				choice = menu.Inputs([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})
				switch choice {
				case 1:
					services.AddProduct()
				case 2:
					services.EditProduct()
				case 3:
					services.DeleteProduct()
				case 4:
					services.AddCustomer()
				case 5:
					services.EditCustomer()
				case 6:
					services.DeleteCustomer()
				case 7:
					services.ExportServices()
				case 8:
					services.ListingsServices()
				case 9:
					lockedLoopTwo = true
				default:
					menu.DisplayErrors("Invalid choice, exiting...")
					return
				}
			}
		} else if inpts == 3 {
			menu.DisplayGoodbye()
			locked = true
		} else {
			menu.DisplayErrors("Invalid choice, exiting...")
			return
		}
	}
}


//AZEVEDO-DA-SILVA, A. (2024) GO-PROJECT-1. [Source code] https://github.com/DestCom-Student-Projects/GO-PROJECT-1