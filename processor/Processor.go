package processor

import (
	"bytes"
	"encoding/json"
	"github.com/jung-kurt/gofpdf"
	"net/http"
	"pdf_generator/model"
	"strconv"
)

func GenerateMsg(msg string,pdfFilePath string)  error{

	var order model.Order

	err := json.Unmarshal([]byte(msg), &order)

	if err != nil{
		return err
	}

	err = PostRequest(order)

	if err != nil {
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

func PostRequest(order model.Order) error {

	url := "http://127.0.0.1:8000/forder/create"

	var jsonStr = []byte(`{"user_id":`+strconv.FormatInt(order.User_id, 10)+`,"order_id":`+strconv.FormatInt(order.Id, 10)+`,"payment_id":1,"final_price":`+strconv.FormatFloat(order.Final_price, 'f', 2, 64)+`,"created_at":"`+order.Created_at+`"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Status == "201" {
		return nil
	}

	return err

	//fmt.Println("response Status:", resp.Status)
	/*fmt.Println("response Headers:", resp.Header)*/
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

