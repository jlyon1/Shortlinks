package api

import (
	"docs/database"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"strings"
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
	Id    string    `json:"id"`
	Count int    `json:"count"`
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
	fmt.Println(val)
	ar := api.Database.Find(val)
	if(val[len(val)-1:] == "+"){
		ar = api.Database.Find(val[:len(val)-1])
	}
	json.Unmarshal([]byte(ar), &a)
	a.Count = a.Count + 1;
	api.Database.Set(val, a)
	if(val[len(val)-1:] == "+"){
		w.Header().Set("Content-Type", "text/html")
		WriteJSON(w,a)

	}else{
		http.Redirect(w, r, a.Link, 302)

	}

}

func (api *API) SetHandler(w http.ResponseWriter, r *http.Request) {
	var b Input
	var a Article
	article := json.NewDecoder(r.Body)
	article.Decode(&b)
	count := 0
	oth := ""
	if(b.Text != ""){
		oth = b.Text
	}else{
		count, _ = strconv.Atoi(api.Database.Find("count"))
		count += 1
	}
	if(api.Database.Find(oth) != ""){
		WriteJSON(w,"bad");
		return;
	}
	fmt.Printf("Count %v", count)
	val := strconv.Itoa(count)

	if b.Title == "" {
		return
	}

	a.Title = b.Title
	a.Link = b.Link
	a.Text = b.Text
	a.Image = b.Image

	if(oth == ""){
		api.Database.Set(val, a)
		api.Database.Set("count", count)
	}else{
		api.Database.Set(oth, a)
		emoval := api.Database.Find("emoj")
		emoval = strings.Replace(emoval,"\\\"","",-1)
		emoval += ";"
		emoval += oth
		api.Database.Set("emoj",emoval)
	}
	WriteJSON(w, count)


}

func (api *API) GetHandler(w http.ResponseWriter, r *http.Request) {
	var articles []Article
	count, _ := strconv.Atoi(api.Database.Find("count"))
	emojval := api.Database.Find("emoj")
	emojval = strings.Replace(emojval,"\"","",-1)
	emojval = strings.Replace(emojval,"\\","",-1)

	a := strings.Split(emojval,";")
	for _,s := range a{
		var a Article
		fmt.Printf(s)
		val := s
		json.Unmarshal([]byte(api.Database.Find(val)), &a)
		a.Id = val
		articles = append(articles, a)
	}
	for i := 1; i <= count; i++ {
		var a Article
		val := strconv.Itoa(i)
		json.Unmarshal([]byte(api.Database.Find(val)), &a)
		a.Id = strconv.Itoa(i)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Write(b)
	return nil
}
