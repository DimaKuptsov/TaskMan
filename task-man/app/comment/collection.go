package comment

import (
	"encoding/json"
	"sort"
)

type CommentsCollection struct {
	Comments []Comment
}

func (cc *CommentsCollection) SortByCreateTime() {
	sort.SliceStable(cc.Comments, func(i, j int) bool {
		return cc.Comments[i].GetCreateTime().After(cc.Comments[j].GetCreateTime())
	})
}

func (cc *CommentsCollection) Len() int {
	return len(cc.Comments)
}

func (cc *CommentsCollection) Add(comment ...Comment) {
	cc.Comments = append(cc.Comments, comment...)
}

func (cc CommentsCollection) IsEmpty() bool {
	return len(cc.Comments) == 0
}

func (cc CommentsCollection) MarshalJSON() ([]byte, error) {
	if cc.IsEmpty() {
		return []byte("{}"), nil
	}
	return json.Marshal(cc.Comments)
}
