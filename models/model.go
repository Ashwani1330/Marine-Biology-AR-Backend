package models

type ModelMetadata struct {
    ID           int    `json:"id"`
    Name         string `json:"name"`
    Description  string `json:"description"`
    Category     string `json:"category"`
    ModelURL     string `json:"model_url"`
    ThumbnailURL string `json:"thumbnail_url"`
    CreatedAt    string `json:"created_at"`
}
