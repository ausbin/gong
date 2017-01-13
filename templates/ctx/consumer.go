package ctx

import (
	"html/template"
	"io"
)

// Consumes an context object, possibly writing output to the Writer
// Normally, this is an abstraction for a template, but in tests, we
// use this to collect and validate the context object generated.
type Consumer interface {
	Consume(io.Writer, Global) error
}

func NewConsumer(templ *template.Template) Consumer {
	return &consumer{templ}
}

type consumer struct {
	templ *template.Template
}

func (c *consumer) Consume(writer io.Writer, ctx Global) error {
	return c.templ.Execute(writer, ctx)
}
