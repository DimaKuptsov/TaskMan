package column

import (
	"encoding/json"
	"sort"
)

type ColumnsCollection struct {
	Columns []Column
}

func (cc *ColumnsCollection) SortByPriority() {
	sort.SliceStable(cc.Columns, func(i, j int) bool {
		return cc.Columns[i].GetPriority() < cc.Columns[j].GetPriority()
	})
}

func (cc *ColumnsCollection) Len() int {
	return len(cc.Columns)
}

func (cc *ColumnsCollection) Add(column ...Column) {
	cc.Columns = append(cc.Columns, column...)
}

func (cc ColumnsCollection) IsEmpty() bool {
	return len(cc.Columns) == 0
}

func (cc ColumnsCollection) MarshalJSON() ([]byte, error) {
	if cc.IsEmpty() {
		return []byte("{}"), nil
	}
	return json.Marshal(cc.Columns)
}
