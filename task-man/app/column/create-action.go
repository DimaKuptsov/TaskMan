package column

type CreateColumnAction struct {
	DTO        CreateDTO
	Repository ColumnsRepository
	Factory    ColumnsFactory
}

func (action CreateColumnAction) Execute() (newColumn Column, err error) {
	newColumn, err = action.Factory.Create(action.DTO)
	if err != nil {
		return
	}
	err = action.Repository.Save(newColumn)
	return newColumn, err
}
