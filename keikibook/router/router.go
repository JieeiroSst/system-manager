package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"keikibook/bus"
	"log"
	"net/http"
)

func InitRouter()*mux.Router{
	router:=mux.NewRouter()

	return router
}

func HandleIndex(w http.ResponseWriter,r *http.Request)  {
	_, _ =w.Write([]byte("Index page"))
}

func HandlePost(w http.ResponseWriter,r *http.Request)  {
	_ = r.ParseForm()
	log.Print(r.PostForm)
	_,_=w.Write([]byte("post\n"))
}

func HandleJson(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	result,_:=json.Marshal(bus.User{"manhquan","123"})
	_, _ = io.WriteString(w, string(result))
}