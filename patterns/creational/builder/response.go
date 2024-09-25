package builder

type ResponseBuilder interface {
	SetEmail(email string) ResponseBuilder
	SetBalance(bal float64) ResponseBuilder
	SetCity(city string) ResponseBuilder
	Build()
}

type responseBuilder struct {
	res *Response
}

type Response struct {
	Email   string
	Balance float64
	City    string
}

func NewResponseBuilder() *responseBuilder {
	return &responseBuilder{
		res: &Response{},
	}
}

func (rb *responseBuilder) SetEmail(email string) {
	rb.res.Email = email

}

func (rb *responseBuilder) SetBalance(bal float64) {
	rb.res.Balance = bal
}

func (rb *responseBuilder) SetCity(city string) {
	rb.res.City = city
}

func (rb *responseBuilder) Build() Response {
	return *rb.res
}
