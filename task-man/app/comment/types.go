package comment

type Text struct {
	Text string `validate:"required,max=5000"`
}

func (t Text) String() string {
	return t.Text
}
