package ansicode

import (
	"io"
)

var _ io.ByteWriter = (*Decoder)(nil)
var _ io.Writer = (*Decoder)(nil)

// Parser is an interface for an ansi sequence parser.
type Parser interface {
	// Advance advances the parser with the given byte.
	Advance(b byte)
}

// Decoder is a byte writer that decodes ANSI escape sequences.
type Decoder struct {
	parser Parser
}

// NewDecoder creates a new Decoder.
func NewDecoder(parser Parser) *Decoder {
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
