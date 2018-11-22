package processor

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jung-kurt/gofpdf"
	"net/http"
	"pdf_generator/model"
	"strconv"
)

func PdfGenerate(msg string,conf model.Config)  error{

	var order model.Order

	err := json.Unmarshal([]byte(msg), &order)

	if err != nil{
		return err
	}

	err = PostRequest(order,conf.ExternalPostPath)

	if err != nil {
		return err
	}

	err = GeneratePdfFile(order,conf.PdfFilePath)

	if err != nil{
		return err
	}

	return nil
}

func GeneratePdfFile(order model.Order, pdfFilePath string) error {

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

	err := pdf.OutputFileAndClose(pdfFilePath+"order"+orderId+".pdf")

	if err != nil{
		return err
	}

	return nil
}

func PostRequest(order model.Order,externalPostPath string) error {

	var jsonStr = []byte(`{"user_id":`+strconv.FormatInt(order.User_id, 10)+`,"order_id":`+strconv.FormatInt(order.Id, 10)+`,"payment_id":1,"final_price":`+strconv.FormatFloat(order.Final_price, 'f', 2, 64)+`,"created_at":"`+order.Created_at+`"}`)
	req, err := http.NewRequest("POST", externalPostPath, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Status != "201 Created" {
		return errors.New("Server error")
	}

	return nil

}

