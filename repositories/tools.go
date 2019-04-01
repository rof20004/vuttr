package repositories

import (
	"database/sql"

	"github.com/lib/pq"
	"github.com/rof20004/vuttr/database"
	"github.com/rof20004/vuttr/models"
)

// FindAllTools - query all tools from database
func FindAllTools() ([]*models.Tool, error) {
	var db = database.GetConnection()

	var tools = make([]*models.Tool, 0)

	rows, err := db.Query("SELECT id, title, link, description, tags FROM tools")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id                       sql.NullInt64
		title, link, description sql.NullString
		tags                     []string
	)

	for rows.Next() {
		rows.Scan(&id, &title, &link, &description, pq.Array(&tags))
		tools = append(tools, &models.Tool{ID: id.Int64, Title: title.String, Link: link.String, Description: description.String, Tags: tags})
	}

	return tools, nil
}

// FindToolsByTag - query all tools from database by a tag identifier
func FindToolsByTag(tag string) ([]*models.Tool, error) {
	var db = database.GetConnection()

	var tools = make([]*models.Tool, 0)

	rows, err := db.Query("SELECT id, title, link, description, tags FROM tools WHERE $1 = any(tags)", tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		id                       sql.NullInt64
		title, link, description sql.NullString
		tags                     []string
	)

	for rows.Next() {
		rows.Scan(&id, &title, &link, &description, pq.Array(&tags))
		tools = append(tools, &models.Tool{ID: id.Int64, Title: title.String, Link: link.String, Description: description.String, Tags: tags})
	}

	return tools, nil
}

// CreateTool - insert a tool in database
func CreateTool(tool models.Tool) (*models.Tool, error) {
	var db = database.GetConnection()

	var id int64
	err := db.QueryRow("INSERT INTO tools(title, link, description, tags) VALUES($1, $2, $3, $4) RETURNING id", tool.Title, tool.Link, tool.Description, pq.Array(tool.Tags)).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &models.Tool{
		ID:          id,
		Title:       tool.Title,
		Link:        tool.Link,
		Description: tool.Description,
		Tags:        tool.Tags,
	}, nil
}

// DeleteTool - remove a tool by its id
func DeleteTool(id string) error {
	var db = database.GetConnection()
	_, err := db.Exec("DELETE FROM tools WHERE id = $1", id)
	return err
}
