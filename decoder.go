package ansicode

import (
	"io"

	"github.com/danielgatis/go-vte"
)

var _ io.ByteWriter = (*Decoder)(nil)
var _ io.Writer = (*Decoder)(nil)

// Decoder is a byte writer that decodes ANSI escape sequences.
type Decoder struct {
	parser *vte.Parser
}

// NewDecoder creates a new Decoder.
func NewDecoder(handler Handler) *Decoder {
	performer := NewPerformer(handler)
	parser := vte.NewParser(performer)

	return &Decoder{
		parser: parser,
	}
}

// WriteByte writes a byte to the decoder.
func (p *Decoder) WriteByte(c byte) error {
	p.parser.Advance(c)
	return nil
}

// Write writes a byte slice to the decoder.
func (p *Decoder) Write(b []byte) (int, error) {
	for _, c := range b {
		err := p.WriteByte(c)
		if err != nil {
			return 0, err
		}
	}

	return len(b), nil
}
