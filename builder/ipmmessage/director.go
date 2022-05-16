package ipmmessage

// Puede cumplir el rol de context
type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) Build() IPMMessage {
	return d.builder.GetIPMMessage()
}
