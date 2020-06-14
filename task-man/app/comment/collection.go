package comment

import "sort"

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
