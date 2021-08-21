package parser

type Bit bool

func (b *Bit) Capture(values []string) error {
	*b = values[0] == "true" || values[0] == "1"
	return nil
}
