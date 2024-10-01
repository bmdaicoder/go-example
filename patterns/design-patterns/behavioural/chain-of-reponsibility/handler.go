package chainofreponsibility

type Handler interface {
	SetNext(Handler) Handler
	Handler(string) string
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *BaseHandler) Handler(request string) string {
	if h.next != nil {
		h.next.Handler(request)
	}
	return ""
}
