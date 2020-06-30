package repository

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/app/column"
	"github.com/DimaKuptsov/task-man/db/errors"
	"github.com/DimaKuptsov/task-man/db/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type ColumnsRepository struct {
	conn *pg.DB
}

func NewColumnsRepository() (repository ColumnsRepository, err error) {
	conn, err := postgres.GetConnection()
	if err != nil {
		return repository, err
	}
	return ColumnsRepository{conn: conn}, err
}

func (cr ColumnsRepository) FindById(id uuid.UUID, withDeleted bool) (column column.Column, err error) {
	query := cr.conn.Model(&column).Where(`"column"."id" = ?`, id)
	if !withDeleted {
		query.Where(`"column"."deleted_at" IS NULL`)
	}
	err = query.Select()
	if err != nil && err.Error() == errors.PGNoRowsFoundError {
		errMsg := fmt.Sprintf("column with id %s not exist", id.String())
		return column, errors.NoRowsFoundError{Message: errMsg}
	}
	return column, err
}

func (cr ColumnsRepository) FindForProject(projectID uuid.UUID, withDeleted bool) (columns column.ColumnsCollection, err error) {
	query := cr.conn.Model(&columns.Columns).Where(`"column"."project_id" = ?`, projectID)
	if !withDeleted {
		query.Where(`"column"."deleted_at" IS NULL`)
	}
	err = query.Select()
	return columns, err
}

func (cr ColumnsRepository) Save(column column.Column) error {
	return cr.conn.Insert(&column)
}

func (cr ColumnsRepository) Update(column column.Column) error {
	return cr.conn.Update(&column)
}

func (cr ColumnsRepository) BatchUpdate(columns column.ColumnsCollection) error {
	if columns.IsEmpty() {
		return nil
	}
	_, err := cr.conn.Model(&columns.Columns).Update()
	return err
}
