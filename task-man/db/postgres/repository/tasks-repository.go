package repository

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/DimaKuptsov/task-man/db/errors"
	"github.com/DimaKuptsov/task-man/db/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type TasksRepository struct {
	conn *pg.DB
}

func NewTasksRepository() (repository TasksRepository, err error) {
	conn, err := postgres.GetConnection()
	if err != nil {
		return repository, err
	}
	return TasksRepository{conn: conn}, err
}

func (tr TasksRepository) FindById(id uuid.UUID, withDeleted bool) (task task.Task, err error) {
	query := tr.conn.Model(&task).Where(`"task"."id" = ?`, id)
	if !withDeleted {
		query.Where(`"task"."deleted_at" IS NULL`)
	}
	err = query.Select()
	if err != nil && err.Error() == errors.PGNoRowsFoundError {
		errMsg := fmt.Sprintf("task with id %s not exist", id.String())
		return task, errors.NoRowsFoundError{Message: errMsg}
	}
	return task, err
}

func (tr TasksRepository) FindByIds(ids []uuid.UUID, withDeleted bool) (tasks task.TasksCollection, err error) {
	if len(ids) == 0 {
		return tasks, err
	}
	query := tr.conn.Model(&tasks.Tasks).Where(`"task"."id" IN (?)`, pg.In(ids))
	if !withDeleted {
		query.Where(`"task"."deleted_at" IS NULL`)
	}
	err = query.Select()
	return tasks, err
}

func (tr TasksRepository) FindForColumn(columnID uuid.UUID, withDeleted bool) (tasks task.TasksCollection, err error) {
	query := tr.conn.Model(&tasks.Tasks).Where(`"task"."column_id" = ?`, columnID)
	if !withDeleted {
		query.Where(`"task"."deleted_at" IS NULL`)
	}
	err = query.Select()
	return tasks, err
}

func (tr TasksRepository) Save(task task.Task) error {
	return tr.conn.Insert(&task)
}

func (tr TasksRepository) BatchUpdate(tasks task.TasksCollection) error {
	if tasks.IsEmpty() {
		return nil
	}
	_, err := tr.conn.Model(&tasks.Tasks).Update()
	return err
}
