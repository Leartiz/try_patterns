package main

import (
	"fmt"
)

type Document interface {
	Render() string
	Clone() Document
}

type Report struct {
	Title string
	// другие "сложные" поля
}

func (r *Report) Render() string { return r.Title }

func (r *Report) Clone() Document {
	cp := *r
	return &cp
}

func export(template Document, title string) {
	doc := template.Clone()
	if r, ok := doc.(*Report); ok {
		r.Title = title // !
	}
	fmt.Println(doc.Render())
}

func main() {
	fmt.Println("dps_go|generative|prototype|v1")

	template := &Report{Title: "monthly report (template)"}

	export(template, "monthly report - alice")
	export(template, "monthly report - bob")

	fmt.Println("template unchanged:", template.Render())
}
