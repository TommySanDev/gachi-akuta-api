package models

import (
    "database/sql"
    "errors"
)

type Character struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Backstory   string `json:"backstory"`    // Backstory
    Image       string `json:"image"`    // Character image
    Chapters    []Chapter `json:"chapters"`  // Chapters appearances
}

func (c *Character) Create(db *sql.DB) error {
    query := "INSERT INTO characters (name, backstory, image) VALUES ($1, $2, $3) RETURNING id"
    err := db.QueryRow(query, c.Name, c.Backstory, c.Image).Scan(&c.ID)
    return err
}

func (c *Character) Get(db *sql.DB) error {
    query := "SELECT name, backstory, image_url FROM characters WHERE id = $1"
    return db.QueryRow(query, c.ID).Scan(&c.Name, &c.Backstory, &c.Image)
}

func GetAll(db *sql.DB) ([]Character, error) {
    rows, err := db.Query("SELECT id, name, backstory, image FROM characters")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var characters []Character
    for rows.Next() {
        var c Character
        if err := rows.Scan(&c.ID, &c.Name, &c.Backstory, &c.Image); err != nil {
            return nil, err
        }
        characters = append(characters, c)
    }
    return characters, nil
}

func (c *Character) Update(db *sql.DB) error {
    // Verifica si el personaje existe en la base de datos
    var exists bool
    queryCheck := "SELECT EXISTS (SELECT 1 FROM characters WHERE id = $1)"
    err := db.QueryRow(queryCheck, c.ID).Scan(&exists)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("el personaje no existe")
    }

    // Si existe, procede a la actualización
    query := "UPDATE characters SET name = $1, backstory = $2, image = $3 WHERE id = $4"
    _, err = db.Exec(query, c.Name, c.Backstory, c.Image, c.ID)
    return err
}

func DeleteCharacter(db *sql.DB, id int) error {
    // Verifica si el personaje existe en la base de datos
    var exists bool
    queryCheck := "SELECT EXISTS (SELECT 1 FROM characters WHERE id = $1)"
    err := db.QueryRow(queryCheck, id).Scan(&exists)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("el personaje no existe")
    }

    // Si existe, procede a la eliminación
    query := "DELETE FROM characters WHERE id = $1"
    _, err = db.Exec(query, id)
    return err
}

