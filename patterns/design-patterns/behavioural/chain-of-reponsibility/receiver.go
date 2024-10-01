package chainofreponsibility

type FirstReceiver struct {
	BaseHandler
}

func (f *FirstReceiver) Handler(request string) string {
	return f.next.Handler(request)
}

func (f *FirstReceiver) setNext(next Handler) Handler {
	f.next = next
	return next
}

type SecondReceiver struct {
	BaseHandler
}

func (s *SecondReceiver) Handle(request string) string {
	if request == "Hello" {
		return "World"
	}
	return s.next.Handler(request)
}

func (s *SecondReceiver) setNext(next Handler) Handler {
	s.next = next
	return next
}
