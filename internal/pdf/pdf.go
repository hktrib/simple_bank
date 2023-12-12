package pdf

import (
	"time"

	"github.com/jung-kurt/gofpdf"
)

type PDFGenerator struct {
	PDF *gofpdf.Fpdf
}

func (p *PDFGenerator) NewPDFFormat() {
	p.PDF = gofpdf.New("L", "mm", "Letter", "")
	p.PDF.AddPage()
	p.PDF.SetFont("Times", "B", 28)
	p.PDF.Cell(40, 10, "Transaction Results")
	p.PDF.Ln(12)
	p.PDF.SetFont("Times", "", 20)
	p.PDF.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
}

func (p *PDFGenerator) CreateHeader(header []string) {
	p.PDF.SetFont("Times", "B", 16)
	p.PDF.SetFillColor(240, 240, 240)
	for _, str := range header {
		p.PDF.CellFormat(40, 7, str, "1", 0, "", true, 0, "")
	}
	p.PDF.Ln(-1)
}

func (p *PDFGenerator) CreateTable(table [][]string) {
	p.PDF.SetFont("Times", "", 16)
	p.PDF.SetFillColor(255, 255, 255)
	alignmentFormat := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range table {
		for i, str := range line {
			p.PDF.CellFormat(40, 7, str, "1", 0, alignmentFormat[i], false, 0, "")
		}
		p.PDF.Ln(-1)
	}
}

func (p *PDFGenerator) SavePDF() error {
	return p.PDF.OutputFileAndClose("transactions.pdf")
}
