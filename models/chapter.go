package models

type Chapter struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Number      int    `json:"number"`
    Characters  []Character `json:"characters"`
}
