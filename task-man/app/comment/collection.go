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

func (cc *CommentsCollection) Add(comment Comment) {
	cc.Comments = append(cc.Comments, comment)
}

func (cc CommentsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(cc.Comments)
}
