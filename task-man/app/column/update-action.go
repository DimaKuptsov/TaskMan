package column

type UpdateColumnAction struct {
	DTO        UpdateDTO
	Repository ColumnsRepository
}

func (action UpdateColumnAction) Execute() (updatedColumn Column, err error) {
	updatedColumn, err = action.Repository.FindById(action.DTO.ID, WithoutDeletedColumns)
	if err != nil {
		return
	}
	if action.DTO.Name != updatedColumn.GetName().String() {
		newName := Name{action.DTO.Name}
		err = updatedColumn.ChangeName(newName)
		if err != nil {
			return
		}
	}
	columnsForUpdate := ColumnsCollection{}
	if action.DTO.Priority != 0 && action.DTO.Priority != updatedColumn.GetPriority() {
		projectColumns, err := action.Repository.FindForProject(updatedColumn.GetProjectID(), WithoutDeletedColumns)
		if err != nil {
			return updatedColumn, err
		}
		for _, column := range projectColumns.Columns {
			if column.GetPriority() == action.DTO.Priority {
				column.ChangePriority(updatedColumn.GetPriority())
				columnsForUpdate.Add(column)
				break
			}
		}
		updatedColumn.ChangePriority(action.DTO.Priority)
	}
	columnsForUpdate.Add(updatedColumn)
	err = action.Repository.BatchUpdate(columnsForUpdate)
	return updatedColumn, err
}
