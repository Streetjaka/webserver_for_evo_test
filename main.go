package main

import (
  "fmt"
  "net/http"
  "html/template"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
//  "log"
)

type Article struct {
  Id uint16
  Name string

}

type Article1 struct {
  Id uint16
  Name string

}

var posts = []Article{}
var posts1 = []Article1{}
var showPost = Article{}



func index(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
    }
    db, err := sql.Open("mysql", "mysql:S%treetjaka123@tcp(176.114.4.20:3306)/evotest")
    if err != nil {
      panic(err)
    }
    defer db.Close()

    res, err := db.Query("SELECT * FROM `evotest`")
    if err != nil {
      panic(err)
    }
    posts = []Article{}
    for res.Next(){
      var post Article
      err = res.Scan(&post.Id, &post.Name)
      if err != nil {
        panic(err)


      }
      posts = append(posts, post)
    }

    t.ExecuteTemplate(w, "index", posts)
}

func save_article(w http.ResponseWriter, r *http.Request){
  name := r.FormValue("name")

  //greetings := r.FormValue("anons")

  if name == ""  {
    fmt.Fprintf(w, "Please enter name")
  } else {

    db, err := sql.Open("mysql", "mysql:S%treetjaka123@tcp(176.114.4.20:3306)/evotest")
    if err != nil {
      panic(err)
    }
    defer db.Close()

    res1, err := db.Query(fmt.Sprintf("SELECT * FROM `evotest` WHERE `name` = '%s'", name))
    if err != nil {
      panic(err)
    }

    posts = []Article{}
    for res1.Next(){
      var post Article
      err = res1.Scan(&post.Id, &post.Name)
      if err != nil {
        panic(err)


      }
      posts = append(posts, post)
    }

     if len(posts) == 0 {
       insert, err := db.Query(fmt.Sprintf("INSERT INTO `evotest` (`name`) VALUES('%s')", name))
       if err != nil {
         panic(err)
       }
       defer insert.Close()
     } else {
       fmt.Fprintf(w, "Виделись %s", name)
     }

    // if res1 == nil{
    //   insert, err := db.Query(fmt.Sprintf("INSERT INTO `evotest` (`name`) VALUES ('%s')", name))
    //   if err != nil {
    //     panic(err)
    //   }
    //
    //   defer insert.Close()




    http.Redirect(w, r, "/", http.StatusSeeOther)
}
}


// func create(w http.ResponseWriter, r *http.Request) {
//   t, err := template.ParseFiles("template/create.html", "template/header.html", "template/footer.html")
//
//   if err != nil {
//     fmt.Fprintf(w, err.Error())
//     }
//
//     t.ExecuteTemplate(w, "create", nil)
// }

func show_post(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  t, err := template.ParseFiles("template/show.html", "template/header.html", "template/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
    }
    db, err := sql.Open("mysql", "mysql:S%treetjaka123@tcp(176.114.4.20:3306)/evotest")
    if err != nil {
      panic(err)
    }
    defer db.Close()


  res, err := db.Query(fmt.Sprintf("SELECT * FROM `evotest`"))
  if err != nil {
    panic(err)
  }
  showPost = Article{}
  for res.Next() {
    var post Article
    err = res.Scan(&post.Id, &post.Name)
    if err != nil {
      panic(err)


    }
    showPost = post

  }

  t.ExecuteTemplate(w, "show", showPost)

}

func show_all1(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("template/show_all1.html", "template/header.html", "template/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
    }
    db, err := sql.Open("mysql", "mysql:S%treetjaka123@tcp(176.114.4.20:3306)/evotest")
    if err != nil {
      panic(err)
    }
    defer db.Close()

    res, err := db.Query("SELECT * FROM `evotest`")
    if err != nil {
      panic(err)
    }
    posts = []Article{}
    for res.Next(){
      var post Article
      err = res.Scan(&post.Id, &post.Name)
      if err != nil {
        panic(err)


      }
      posts = append(posts, post)
    }

    t.ExecuteTemplate(w, "show_all1", posts)
}

func handleFunc(){
  rtr := mux.NewRouter()
  rtr.HandleFunc("/", index).Methods("GET")
  //rtr.HandleFunc("/create", create).Methods("GET")
  rtr.HandleFunc("/save_article", save_article).Methods("POST")
  rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")
  rtr.HandleFunc("/show_all1", show_all1).Methods("GET")
  http.Handle("/", rtr)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  http.ListenAndServe(":8080",nil)
}

func main(){
  handleFunc()
}
