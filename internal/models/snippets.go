package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
// 	stmt := `INSERT INTO snippets (title, content, created, expires)
//         VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '$3 DAYS')
//         RETURNING id`
// 	var id int
// 	// Use QueryRow to execute the statement and scan the returned id directly.
// 	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes).
	stmt := `INSERT INTO snippets (title, content, created, expires)
        VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL $3)
        RETURNING id`
	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// title, content, and expiry values for the placeholder parameters. This
	// method returns a sql.Result type, which contains some basic
	// information about what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, title, content, fmt.Sprintf("%d DAYS", expires))
	if err != nil {
		return 0, err
	}
	// Use the LastInsertId() method on the result to get the ID of our
	// newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}


func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
        WHERE expires > CURRENT_TIMESTAMP AND id = $1`
	row := m.DB.QueryRow(stmt, id)
	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
        WHERE expires > CURRENT_TIMESTAMP ORDER BY id DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets := []*Snippet{}
	for rows.Next() {
		s := &Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
