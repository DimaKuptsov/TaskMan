package task

import "sort"

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

func (tc *TasksCollection) Add(task Task) {
	tc.Tasks = append(tc.Tasks, task)
}
