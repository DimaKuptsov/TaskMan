package project

import (
	"encoding/json"
	"sort"
)

type ProjectsCollection struct {
	Projects []Project
}

func (pc *ProjectsCollection) SortByCreateTime() {
	sort.SliceStable(pc.Projects, func(i, j int) bool {
		return pc.Projects[i].GetCreateTime().After(pc.Projects[j].GetCreateTime())
	})
}

func (pc *ProjectsCollection) Len() int {
	return len(pc.Projects)
}

func (pc *ProjectsCollection) Add(project Project) {
	pc.Projects = append(pc.Projects, project)
}

func (pc ProjectsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(pc.Projects)
}
