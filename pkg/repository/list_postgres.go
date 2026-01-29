package repository

import (
	"github.com/BudjakovDmitry/go_todo_app"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var id int
	row := tx.QueryRow(
		"INSERT INTO lists (title, description) VALUES ($1, $2) RETURNING id",
		list.Title,
		list.Description,
	)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	_, err = tx.Exec(
		"INSERT INTO users_lists (user_id, list_id) VALUES ($1, $2)",
		userId,
		id,
	)
	if err != nil {
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := `
SELECT l.id, l.title, l.description
FROM lists l
JOIN users_lists ul ON ul.list_id = l.id
WHERE ul.user_id = $1
`
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList
	query := `
SELECT l.id, l.title, l.description
FROM lists l
JOIN users_lists ul ON ul.list_id = l.id
WHERE
	ul.user_id = $1 AND
	l.id = $2
`
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}
