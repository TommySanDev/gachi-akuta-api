package controllers

import (
    "database/sql"
    "net/http"
    "encoding/json"
    "gachi-akuta-api/models"
)

type ChapterController struct {
    DB *sql.DB
}

func NewChapterController(db *sql.DB) *ChapterController {
    return &ChapterController{DB: db}
}

func (cc *ChapterController) Create(w http.ResponseWriter, r *http.Request) {
    var chapter models.Chapter
    err := json.NewDecoder(r.Body).Decode(&chapter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = chapter.Create(cc.DB)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(chapter)
}

func (cc *ChapterController) GetAll(w http.ResponseWriter, r *http.Request) {
    chapters, err := models.GetAllChapters(cc.DB)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(chapters)
}

func (cc *ChapterController) GetByID(w http.ResponseWriter, r *http.Request) {
    var chapter models.Chapter

    id := extractIDFromRequest(r)
    chapter.ID = id

    err := chapter.Get(cc.DB)
    if err != nil {
        http.Error(w, "Chapter not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(chapter)
}

func (cc *ChapterController) Update(w http.ResponseWriter, r *http.Request) {
    var chapter models.Chapter

    id := extractIDFromRequest(r)
    chapter.ID = id

    err := json.NewDecoder(r.Body).Decode(&chapter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = chapter.Update(cc.DB)
    if err != nil {
        http.Error(w, "Failed to update chapter", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(chapter)
}

func (cc *ChapterController) Delete(w http.ResponseWriter, r *http.Request) {
    id := extractIDFromRequest(r)

    err := models.DeleteChapter(cc.DB, id)
    if err != nil {
        http.Error(w, "Failed to delete chapter", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
