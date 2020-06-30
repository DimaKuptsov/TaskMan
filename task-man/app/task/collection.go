package task

import (
	"encoding/json"
	"sort"
)

type TasksCollection struct {
	Tasks []Task
}

func (tc *TasksCollection) SortByPriority() {
	sort.SliceStable(tc.Tasks, func(i, j int) bool {
		return tc.Tasks[i].GetPriority() < tc.Tasks[j].GetPriority()
	})
}

func (tc *TasksCollection) Len() int {
	return len(tc.Tasks)
}

func (tc *TasksCollection) Add(task ...Task) {
	tc.Tasks = append(tc.Tasks, task...)
}

func (tc TasksCollection) IsEmpty() bool {
	return len(tc.Tasks) == 0
}

func (tc TasksCollection) MarshalJSON() ([]byte, error) {
	if tc.IsEmpty() {
		return []byte("{}"), nil
	}
	return json.Marshal(tc.Tasks)
}
