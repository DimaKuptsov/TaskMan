package task

import "database/sql/driver"

type Name struct {
	Name string `validate:"required,max=500"`
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

type Description struct {
	Description string `validate:"max=5000"`
}

func (d Description) String() string {
	return d.Description
}

func (d *Description) Scan(value interface{}) error {
	description := value.([]byte)
	d.Description = string(description)
	return nil
}

func (d Description) Value() (driver.Value, error) {
	return d.Description, nil
}
