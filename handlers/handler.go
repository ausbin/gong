package handlers

type Handler interface {
	Serve(request Request)
}

type Request interface {
	// Attributes
	Path() string
	// If the registered path for this handler is /X/Y/, return /Z/ for
	// Path() == /X/Y/Z/
	Subtree() string

	// Actions
	Redirect(string)
	Write([]byte) (int, error)
	Error(error)
}

type Func func(r Request)

func (f Func) Serve(r Request) { f(r) }
