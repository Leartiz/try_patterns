package balanced

import "fmt"

// Repository - one abstraction at the boundary.
type Repository interface {
	Save(productID string) error
}

type memoryRepo struct{}

func NewMemoryRepo() Repository {
	return &memoryRepo{}
}

func (r *memoryRepo) Save(productID string) error {
	fmt.Printf("balanced: save order product=%s\n", productID)
	return nil
}

// Service - depends on interface, not on DB details.
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) PlaceOrder(productID string) error {
	if productID == "" {
		return fmt.Errorf("product id required")
	}
	return s.repo.Save(productID)
}
