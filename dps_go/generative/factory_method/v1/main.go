package main

import (
	"fmt"
)

type Document interface {
	Render() string
}

type Creator interface {
	Create() Document
}

// products
// -----------------------------------------------------------------------

type PDFDocument struct{}

func (d *PDFDocument) Render() string { return "pdf document" }

type MarkdownDocument struct{}

func (d *MarkdownDocument) Render() string { return "markdown document" }

// creators (factory method per type)
// -----------------------------------------------------------------------

type PDFCreator struct{}

func (c *PDFCreator) Create() Document { return &PDFDocument{} }

type MarkdownCreator struct{}

func (c *MarkdownCreator) Create() Document { return &MarkdownDocument{} }

func export(c Creator) {
	doc := c.Create()
	fmt.Println(doc.Render())
}

// -----------------------------------------------------------------------

func main() {
	fmt.Println("dps_go|generative|factory_method|v1")

	export(&PDFCreator{})
	export(&MarkdownCreator{})
}
