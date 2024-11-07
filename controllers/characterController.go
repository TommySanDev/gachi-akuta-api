package controllers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "gachi-akuta-api/models"
)

type CharacterController struct {
    DB *sql.DB
}

func NewCharacterController(db *sql.DB) *CharacterController {
    return &CharacterController{DB: db}
}

func (cc *CharacterController) Create(w http.ResponseWriter, r *http.Request) {
    var character models.Character
    err := json.NewDecoder(r.Body).Decode(&character)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = character.Create(cc.DB)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(character)
}

func (cc *CharacterController) GetAll(w http.ResponseWriter, r *http.Request) {
    characters, err := models.GetAllCharacters(cc.DB) // Corrección aquí
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(characters)
}

