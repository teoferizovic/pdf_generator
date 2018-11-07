package processor

import (
	"encoding/json"
	"github.com/jung-kurt/gofpdf"
	"pdf_generator/model"
	"strconv"
)

func GenerateMsg(msg string,pdfFilePath string)  error{

	var order model.Order

	err := json.Unmarshal([]byte(msg), &order)

	if err != nil{
		return err
	}

	orderId := strconv.FormatInt(order.Id, 10)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "OrderId:"+orderId)
	pdf.Ln(10);
	pdf.Cell(40, 10, "UserId:"+strconv.FormatInt(order.User_id, 10))
	pdf.Ln(10);
	pdf.Cell(40, 10, "Status:"+order.Status)
	pdf.Ln(10);
	pdf.Cell(40, 10, "Final price:"+strconv.FormatFloat(order.Final_price, 'f', 6, 64))
	pdf.Ln(10);
	pdf.Cell(40, 10, "Created:"+order.Created_at)

	err = pdf.OutputFileAndClose(pdfFilePath+"order"+orderId+".pdf")

	if err != nil{
		return err
	}

	return nil
}

