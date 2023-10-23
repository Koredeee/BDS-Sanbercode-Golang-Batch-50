package category

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()

	router.GET("/categories", GetCategories)
	router.POST("/categories", CreateCategory)
	router.PUT("/categories/:id", UpdateCategory)
	router.DELETE("/categories/:id", DeleteCategory)
	router.GET("/categories/:id/articles", GetArticlesByCategory)

	router.GET("/articles", GetArticles)
	router.POST("/articles", CreateArticle)
	router.PUT("/articles/:id", UpdateArticle)
	router.DELETE("/articles/:id", DeleteArticle)

	http.ListenAndServe(":8080", router)
}

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	ID            int
	Title         string
	Content       string
	ImageURL      string
	ArticleLength string
	CategoryID    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func GetCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("SELECT id, name, created_at, updated_at FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	jsonResponse, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func CreateCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO categories (name, created_at, updated_at) VALUES (?, NOW(), NOW())", name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categoryID, _ := result.LastInsertId()

	response := map[string]int{"category_id": int(categoryID)}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE categories SET name = ?, updated_at = NOW() WHERE id = ?", name, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Category updated successfully"))
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Category deleted successfully"))
}

func GetArticlesByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT id, title, content, image_url, article_length, category_id, created_at, updated_at FROM articles WHERE category_id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	articles := []Article{}
	for rows.Next() {
		article := Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ImageURL, &article.ArticleLength, &article.CategoryID, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	jsonResponse, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func GetArticles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("SELECT id, title, content, image_url, article_length, category_id, created_at, updated_at FROM articles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	articles := []Article{}
	for rows.Next() {
		article := Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ImageURL, &article.ArticleLength, &article.CategoryID, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	title := r.FormValue("title")
	content := r.FormValue("content")
	imageURL := r.FormValue("image_url")
	categoryIDStr := r.FormValue("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)

	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	if imageURL == "" || !ValidateImageURL(imageURL) {
		http.Error(w, "Invalid image URL", http.StatusBadRequest)
		return
	}

	if !ValidateTitle(title) {
		http.Error(w, "Title must have less than 3 words", http.StatusBadRequest)
		return
	}

	articleLength := ConvertContentToArticleLength(content)

	_, err = db.Exec("INSERT INTO articles (title, content, image_url, article_length, category_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())", title, content, imageURL, articleLength, categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	imageURL := r.FormValue("image_url")
	categoryIDStr := r.FormValue("category_id")
	categoryID, err := strconv.Atoi(categoryIDStr)

	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	if imageURL == "" || !ValidateImageURL(imageURL) {
		http.Error(w, "Invalid image URL", http.StatusBadRequest)
		return
	}

	if !ValidateTitle(title) {
		http.Error(w, "Title must have less than 3 words", http.StatusBadRequest)
		return
	}

	articleLength := ConvertContentToArticleLength(content)

	_, err = db.Exec("UPDATE articles SET title = ?, content = ?, image_url = ?, article_length = ?, category_id = ?, updated_at = NOW() WHERE id = ?", title, content, imageURL, articleLength, categoryID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Article updated successfully"))
}

func DeleteArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM articles WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Article deleted successfully"))
}

func ValidateImageURL(url string) bool {
	pattern := `^(https?://)?([a-zA-Z0-9.-]+(\.[a-zA-Z]{2,4})+)(:[0-9]+)?(/.*)?$`

	matched, err := regexp.MatchString(pattern, url)
	if err != nil {
		return false // Return false in case of an error
	}

	return matched
}

func ValidateTitle(title string) bool {
	words := strings.Fields(title)

	return len(words) <= 3
}

func ConvertContentToArticleLength(content string) string {
	wordCount := len(strings.Fields(content))
	if wordCount <= 100 || len(content) <= 400 {
		return "pendek"
	} else if wordCount <= 200 || len(content) <= 800 {
		return "sedang"
	}
	return "panjang"
}
