package main

import "fmt"

type Document interface {
	Render() string
}

type PDFDocument struct{}

func (d *PDFDocument) Render() string { return "pdf document" }

type MarkdownDocument struct{}

func (d *MarkdownDocument) Render() string { return "markdown document" }

func NewDocument(format string) Document {
	switch format {
	case "pdf":
		return &PDFDocument{}
	case "md":
		return &MarkdownDocument{}
	default:
		panic("unknown format: " + format)
	}
}

func export(format string) {
	doc := NewDocument(format)
	fmt.Println(doc.Render())
}

func main() {
	fmt.Println("dps_go|generative|simple_factory|v1")

	export("pdf")
	export("md")
}
