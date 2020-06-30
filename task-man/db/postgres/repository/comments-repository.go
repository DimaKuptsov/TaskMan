package repository

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/app/comment"
	"github.com/DimaKuptsov/task-man/db/errors"
	"github.com/DimaKuptsov/task-man/db/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type CommentsRepository struct {
	conn *pg.DB
}

func NewCommentsRepository() (repository CommentsRepository, err error) {
	conn, err := postgres.GetConnection()
	if err != nil {
		return repository, err
	}
	return CommentsRepository{conn: conn}, err
}

func (cr CommentsRepository) FindById(id uuid.UUID, withDeleted bool) (comment comment.Comment, err error) {
	query := cr.conn.Model(&comment).Where(`"comment"."id" = ?`, id)
	if !withDeleted {
		query.Where(`"comment"."deleted_at" IS NULL`)
	}
	err = query.Select()
	if err != nil && err.Error() == errors.PGNoRowsFoundError {
		errMsg := fmt.Sprintf("comment with id %s not exist", id.String())
		return comment, errors.NoRowsFoundError{Message: errMsg}
	}
	return comment, err
}

func (cr CommentsRepository) FindForTasks(tasksIDs []uuid.UUID, withDeleted bool) (comments comment.CommentsCollection, err error) {
	if len(tasksIDs) == 0 {
		return comments, err
	}
	query := cr.conn.Model(&comments.Comments).Where(`"comment"."task_id" IN (?)`, pg.In(tasksIDs))
	if !withDeleted {
		query.Where(`"comment"."deleted_at" IS NULL`)
	}
	err = query.Select()
	return comments, err
}

func (cr CommentsRepository) Save(comment comment.Comment) error {
	return cr.conn.Insert(&comment)
}

func (cr CommentsRepository) BatchUpdate(comments comment.CommentsCollection) error {
	if comments.IsEmpty() {
		return nil
	}
	_, err := cr.conn.Model(&comments.Comments).Update()
	return err
}
