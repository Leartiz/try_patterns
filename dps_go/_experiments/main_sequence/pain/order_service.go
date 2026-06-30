package pain

import "fmt"

// OrderService - zone of pain: concrete, does everything, hard to test or extend.
type OrderService struct{}

func (s *OrderService) PlaceOrder(productID string, email string) error {
	if productID == "" {
		return fmt.Errorf("product id required")
	}
	if email == "" {
		return fmt.Errorf("email required")
	}

	// pretend DB
	fmt.Printf("pain: INSERT order product=%s\n", productID)

	// pretend SMTP
	fmt.Printf("pain: SEND email to %s\n", email)

	return nil
}
