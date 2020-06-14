package column

type Name struct {
	Name string `validate:"required,max=255"`
}

func (n Name) String() string {
	return n.Name
}
