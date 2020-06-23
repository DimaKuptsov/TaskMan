package handlers

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers/columns"
	"github.com/DimaKuptsov/task-man/handlers/comments"
	"github.com/DimaKuptsov/task-man/handlers/projects"
	"github.com/DimaKuptsov/task-man/handlers/tasks"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	setMiddlewares(r)
	setProjectsRoutes(r)
	setColumnsRoutes(r)
	setTasksRoutes(r)
	setCommentsRoutes(r)
	return r
}

func setMiddlewares(r *chi.Mux) {
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
}

func setProjectsRoutes(r *chi.Mux) {
	r.Route("/projects", func(r chi.Router) {
		r.Get("/all", projects.GetAll)
		r.Get(fmt.Sprintf("/{%s}", projects.ProjectIDField), projects.GetByID)
		r.Post("/create", projects.Create)
		r.Put("/update", projects.Update)
		r.Delete(fmt.Sprintf("/delete/{%s}", projects.ProjectIDField), projects.Delete)
	})
}

func setColumnsRoutes(r *chi.Mux) {
	r.Route("/columns", func(r chi.Router) {
		r.Get(fmt.Sprintf("/project/{%s}", columns.ProjectIDField), columns.GetForProject)
		r.Get(fmt.Sprintf("/{%s}", columns.ColumnIDField), columns.GetByID)
		r.Post("/create", columns.Create)
		r.Put("/update", columns.Update)
		r.Delete(fmt.Sprintf("/delete/{%s}", columns.ColumnIDField), columns.Delete)
	})
}

func setTasksRoutes(r *chi.Mux) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get(fmt.Sprintf("/column/{%s}", tasks.ColumnIDField), tasks.GetForColumn)
		r.Get(fmt.Sprintf("/{%s}", tasks.TaskIDField), tasks.GetByID)
		r.Post("/create", tasks.Create)
		r.Put("/update", tasks.Update)
		r.Delete(fmt.Sprintf("/delete/{%s}", tasks.TaskIDField), tasks.Delete)
	})
}

func setCommentsRoutes(r *chi.Mux) {
	r.Route("/comments", func(r chi.Router) {
		r.Get(fmt.Sprintf("/task/{%s}", comments.TaskIDField), comments.GetForTask)
		r.Get(fmt.Sprintf("/{%s}", comments.CommentIDField), comments.GetByID)
		r.Post("/create", comments.Create)
		r.Put("/update", comments.Update)
		r.Delete(fmt.Sprintf("/delete/{%s}", comments.CommentIDField), comments.Delete)
	})
}
