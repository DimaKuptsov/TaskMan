package task

type Name struct {
	Name string `validate:"required,max=500"`
}

func (n Name) String() string {
	return n.Name
}

type Description struct {
	Description string `validate:"max=5000"`
}

func (d Description) String() string {
	return d.Description
}
