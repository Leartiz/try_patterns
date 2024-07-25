package main

import (
	"fmt"
	"unicode"
)

// -----------------------------------------------------------------------

/*
То к чему обращается бизнес-логика.
*/
type Client interface {
	method(data string)
}

type Adapter struct {
	service *Service
}

func (a *Adapter) method(data string) {
	inp := ServiceInput{digitCount: 0}
	for _, ch := range data {
		if unicode.IsDigit(ch) {
			inp.digitCount++
		}
	}

	a.service.serviceMethod(inp) // to external useful service!
}

// -----------------------------------------------------------------------

type Service struct{}

type ServiceInput struct {
	digitCount uint64
}

func (s *Service) serviceMethod(inp ServiceInput) {
	fmt.Printf("digit count: %v", inp.digitCount)
}

// -----------------------------------------------------------------------

func main() {
	var service = &Service{}
	var client Client = &Adapter{service: service}

	client.method("abc123abc")
}
