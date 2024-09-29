package adapter

type PaymentGateway interface {
	ProcessPayment(amount float32) bool
}

type PaypalAdapter struct {
	Paypal *Paypal
}

func NewPaypalAdapter() *PaypalAdapter {
	return &PaypalAdapter{
		Paypal: &Paypal{},
	}
}

func (p *PaypalAdapter) ProcessPayment(amount float32) bool {
	return p.Paypal.MakePayment(amount)
}

type AmazonAdapter struct {
	Amazon *Amazon
}

func NewAmazonAdapter() *AmazonAdapter {
	return &AmazonAdapter{
		Amazon: &Amazon{},
	}
}

func (a *AmazonAdapter) ProcessPayment(amount float32) bool {
	return a.Amazon.PayAmazon(amount)
}

type Paypal struct{}

func (p *Paypal) MakePayment(amount float32) bool {
	return true
}

type Amazon struct{}

func (a *Amazon) PayAmazon(amount float32) bool {
	return true
}
