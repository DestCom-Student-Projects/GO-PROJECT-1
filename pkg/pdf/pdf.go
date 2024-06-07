package pdfBuilder

import (
	"fmt"
	"os"

	"github.com/DestCom-Student-Projects/GO-PROJECT-1/pkg/database"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(product database.Product, quantity int, customer database.Customer, order database.Order) error {
	var err error

	if _ , err := os.Stat("orders_pdf"); os.IsNotExist(err) {
        err := os.Mkdir("orders_pdf", 0755)
        if err != nil {
            return err
        }
    }

    // Génération du PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Order Details")
    pdf.Ln(20)
    pdf.SetFont("Arial", "", 12)
    pdf.Cell(40, 10, fmt.Sprintf("Client Name: %s %s", customer.Prenom, customer.Nom))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Client ID: %d", customer.ID))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Product: %s", product.Titre))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Product ID: %d", product.ID))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Quantity: %d", quantity))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Total Price: %.2f", float64(quantity) * float64(product.Prix)))
    pdf.Ln(10)
    pdf.Cell(40, 10, fmt.Sprintf("Order ID: %d", order.ID))

    filePath := fmt.Sprintf("orders_pdf/order-%d-%d.pdf", customer.ID ,order.ID)
    err = pdf.OutputFileAndClose(filePath)
    if err != nil {
        return err
    }

    fmt.Println("PDF generated successfully")
    
    return nil
}