package repository

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/app/project"
	"github.com/DimaKuptsov/task-man/db/errors"
	"github.com/DimaKuptsov/task-man/db/postgres"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type ProjectsRepository struct {
	conn *pg.DB
}

func NewProjectsRepository() (repository ProjectsRepository, err error) {
	conn, err := postgres.GetConnection()
	if err != nil {
		return repository, err
	}
	return ProjectsRepository{conn: conn}, err
}

func (pr ProjectsRepository) FindNotDeleted() (projects project.ProjectsCollection, err error) {
	err = pr.conn.Model(&projects.Projects).
		Where(`"project"."deleted_at" IS NULL`).
		Select()
	return projects, err
}

func (pr ProjectsRepository) FindById(id uuid.UUID, withDeleted bool) (project project.Project, err error) {
	query := pr.conn.Model(&project).Where(`"project"."id" = ?`, id)
	if !withDeleted {
		query.Where(`"project"."deleted_at" IS NULL`)
	}
	err = query.Select()
	if err != nil && err.Error() == errors.PGNoRowsFoundError {
		errMsg := fmt.Sprintf("project with id %s not exist", id.String())
		return project, errors.NoRowsFoundError{Message: errMsg}
	}
	return project, err
}

func (pr ProjectsRepository) Save(project project.Project) error {
	return pr.conn.Insert(&project)
}

func (pr ProjectsRepository) Update(project project.Project) error {
	return pr.conn.Update(&project)
}
