package main

type BookPDF struct {
}

func (b BookPDF) NextPage() {
	println("Turn to next page in book.")
}

type BookPDFAdapter struct {
	book BookPDF
}

func (b BookPDFAdapter) Forward() {
	b.book.NextPage()
}

func main() {
	a := BookPDFAdapter{}
	a.Forward()
}
