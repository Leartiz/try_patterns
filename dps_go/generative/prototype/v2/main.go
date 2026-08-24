package main

import (
	"fmt"
	"slices"
)

type Document interface {
	Render() string
	Clone() Document
}

type Report struct {
	Title    string
	Sections []string
}

func (r *Report) Render() string {
	return fmt.Sprintf("%s: %v",
		r.Title, r.Sections)
}

func shallowCopy(r *Report) *Report {
	cp := *r
	return &cp
}

func (r *Report) Clone() Document {
	cp := *r
	cp.Sections = slices.Clone(r.Sections) // !
	return &cp
}

func newTemplate() *Report {
	return &Report{
		Title:    "monthly report (template)",
		Sections: []string{"intro", "metrics"},
	}
}

func main() {
	fmt.Println("dps_go|generative|prototype|v2")

	fmt.Println("--- shallow: cp := *r ---")
	template := newTemplate()
	shallow := shallowCopy(template)
	shallow.Sections[0] = "changed in copy"
	fmt.Println("copy:    ", shallow.Render())
	fmt.Println("template:", template.Render())

	fmt.Println()
	fmt.Println("--- deep: Clone() + slices.Clone ---")
	template = newTemplate()
	deep := template.Clone().(*Report)
	deep.Sections[0] = "changed in copy"
	fmt.Println("copy:    ", deep.Render())
	fmt.Println("template:", template.Render())
}
