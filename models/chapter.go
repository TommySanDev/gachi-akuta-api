package models

import (
    "database/sql"
    "errors"
)

type Chapter struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Number      int    `json:"number"`
    Characters  []Character `json:"characters"`
}

func (c *Chapter) Create(db *sql.DB) error {
    query := "INSERT INTO chapters (title, number) VALUES ($1, $2) RETURNING id"
    err := db.QueryRow(query, c.Title, c.Number).Scan(&c.ID)
    return err
}

func (c *Chapter) Get(db *sql.DB) error {
    query := "SELECT title, number FROM chapters WHERE id = $1"
    return db.QueryRow(query, c.ID).Scan(&c.Title, &c.Number)
}

func GetAllChapters(db *sql.DB) ([]Chapter, error) {
    rows, err := db.Query("SELECT id, title, number FROM chapters")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var chapters []Chapter
    for rows.Next() {
        var c Chapter
        if err := rows.Scan(&c.ID, &c.Title, &c.Number); err != nil {
            return nil, err
        }
        chapters = append(chapters, c)
    }
    return chapters, nil
}

func (c *Chapter) Update(db *sql.DB) error {
    //Verificacion chapter
    var exists bool
    queryCheck := "SELECT EXISTS (SELECT 1 FROM chapters WHERE id = $1)"
    err := db.QueryRow(queryCheck, c.ID).Scan(&exists)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("el capítulo no existe")
    }

    // Update
    query := "UPDATE chapters SET title = $1, number = $2 WHERE id = $3"
    _, err = db.Exec(query, c.Title, c.Number, c.ID)
    return err
}

func DeleteChapter(db *sql.DB, id int) error {
    //Verificacion
    var exists bool
    queryCheck := "SELECT EXISTS (SELECT 1 FROM chapters WHERE id = $1)"
    err := db.QueryRow(queryCheck, id).Scan(&exists)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("el capítulo no existe")
    }

    //Delete
    query := "DELETE FROM chapters WHERE id = $1"
    _, err = db.Exec(query, id)
    return err
}
