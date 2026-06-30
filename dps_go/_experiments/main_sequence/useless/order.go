package useless

import "fmt"

// Interfaces on top of interfaces - zone of uselessness: abstract, awkward API.

type OrderGateway interface {
	Dispatch(cmd Command) error
}

type Command interface {
	Execute(ctx Context) error
}

type Context interface {
	Logger() Logger
}

type Logger interface {
	Info(msg string)
}

type Factory interface {
	Build() OrderGateway
}

type factory struct{}

func NewFactory() Factory {
	return &factory{}
}

func (f *factory) Build() OrderGateway {
	return &gateway{}
}

type gateway struct{}

func (g *gateway) Dispatch(cmd Command) error {
	return cmd.Execute(&ctx{})
}

type ctx struct{}

func (c *ctx) Logger() Logger {
	return &logger{}
}

type logger struct{}

func (l *logger) Info(msg string) {
	fmt.Println("useless:", msg)
}

type placeOrder struct {
	productID string
}

func NewPlaceOrderCommand(productID string) Command {
	return &placeOrder{productID: productID}
}

func (c *placeOrder) Execute(ctx Context) error {
	ctx.Logger().Info(fmt.Sprintf("place order product=%s", c.productID))
	return nil
}

// PlaceOrder - caller must wire factory -> gateway -> command manually.
func PlaceOrder(productID string) error {
	return NewFactory().Build().Dispatch(NewPlaceOrderCommand(productID))
}
