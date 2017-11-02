package api

import (
	"docs/database"
	"encoding/json"
	"net/http"
	"strconv"
	// "fmt"
	"github.com/gorilla/mux"
)

type API struct {
	Database database.DB
}

type Article struct {
	Title string `json:"title"`
	Link  string `json:"link"`
	Text  string `json:"text"`
	Image string `json:"image"`
	Id    int    `json:"id"`
}
type Input struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Text     string `json:"text"`
	Image    string `json:"image"`
	Password string `json:"password"`
}

func (api *API) ShortLink(w http.ResponseWriter, r *http.Request) {
	var a Article
	val := mux.Vars(r)["val"]
	json.Unmarshal([]byte(api.Database.Find(val)), &a)
	http.Redirect(w, r, a.Link, 301)

}

func (api *API) SetHandler(w http.ResponseWriter, r *http.Request) {
	var b Input
	var a Article
	article := json.NewDecoder(r.Body)
	article.Decode(&b)
	count, _ := strconv.Atoi(api.Database.Find("count"))
	count += 1
	val := strconv.Itoa(count)

	if b.Title == "" {
		return
	}

	a.Title = b.Title
	a.Link = b.Link
	a.Text = b.Text
	a.Image = b.Image
	if b.Password == api.Database.Find("password") {
		api.Database.Set(val, a)
		api.Database.Set("count", count)
	}

}

func (api *API) GetHandler(w http.ResponseWriter, r *http.Request) {
	var articles []Article
	count, _ := strconv.Atoi(api.Database.Find("count"))
	for i := 1; i <= count; i++ {
		var a Article
		val := strconv.Itoa(i)
		json.Unmarshal([]byte(api.Database.Find(val)), &a)
		a.Id = i
		articles = append(articles, a)
	}

	WriteJSON(w, articles)
}

func (api *API) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func (api *API) AddIndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "add.html")
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}
