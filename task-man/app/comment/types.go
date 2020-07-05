package comment

import "database/sql/driver"

type Text struct {
	Text string `validate:"required,max=5000"`
}

func (t Text) String() string {
	return t.Text
}

func (t *Text) Scan(value interface{}) error {
	text := value.([]byte)
	t.Text = string(text)
	return nil
}

func (t Text) Value() (driver.Value, error) {
	return t.Text, nil
}
