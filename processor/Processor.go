package processor

import "fmt"

func GenerateMsg(msg string)  {
	fmt.Println(msg)
}


/*pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	pdf.Cell(50, 110, str)
	pdf.Cell(40, 10, "Hello, world")
	pdf.Cell(40, 10, "Hello, world")
	pdf.Cell(40, 10, "Hello, world")
	pdf.Cell(40, 10, "Hello, world")

	err := pdf.OutputFileAndClose(conf.PdfFilePath+"hello2.pdf")

	if err != nil{
		panic(err)
	}*/

//https://github.com/jung-kurt/gofpdf