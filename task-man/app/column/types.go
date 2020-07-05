package column

import "database/sql/driver"

type Name struct {
	Name string `validate:"required,max=255"`
}

func (n Name) String() string {
	return n.Name
}

func (n *Name) Scan(value interface{}) error {
	name := value.([]byte)
	n.Name = string(name)
	return nil
}

func (n Name) Value() (driver.Value, error) {
	return n.Name, nil
}
