package main
//go run HellowWorld.go
//cmd direction of mongo 4.0\bin -> 
//mongo conect{mongod --dbpath "C:\Program Files\MongoDB\Server\4.0\data\db"}
//go get github.com/denisenkom/go-mssqldb
//go get github.com/gorilla/mux
//* go get github.com/mongodb/mongo-go-driver
// go get gopkg.in/mgo.v2
import(
	"fmt"
	"time"
	"encoding/json"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"database/sql"
	"html/template"
	//_ "github.com/denisenkom/go-mssqldb"
	//_"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/gorilla/mux"
)
var tpl *template.Template
var people []Person
type Person struct
{
	Id bson.ObjectId `bson:"_id,omitempty"`
	_id string `json:"id"`
	Name string `json:"name"`
	LastName string `json:"lastName"`
}
func OpenConect()*mgo.Session{
	session, err := mgo.Dial("localhost:27017")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Connected!!")
	return session
}

//close database db.close();
func GetPeopleHandler(w http.ResponseWriter, r *http.Request){
	db := OpenConect()
	c := db.DB("User").C("user")
	var Result []Person
	err := c.Find(nil).All(&Result)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(Result)
	db.Close();
	answer, err := json.Marshal(Result)
	if err!=nil{
		log.Fatal(err)
	}
	w.Write(answer);
}
func print10000numbers(s string){
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func InsertPersonHandler(w http.ResponseWriter, r *http.Request){
	db := OpenConect()
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	c := db.DB("User").C("user")
	err := c.Insert(person)
	if err!=nil{
		log.Fatal(err)
	}
	log.Printf("Exito")
	db.Close();
}
func DeletePersonHandler(w http.ResponseWriter, r *http.Request){
	db := OpenConect()
	params:= mux.Vars(r)
	var person Person
	person._id = params["id"]
	c := db.DB("User").C("user")
	err := c.RemoveId(bson.ObjectIdHex(person._id))
	if err!=nil{
		log.Fatal(err)
	}
	log.Printf("Exito")
	db.Close();
}
func UpdatePersonHandler(w http.ResponseWriter, r *http.Request){
	db := OpenConect()
	params:= mux.Vars(r)
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	person._id = params["id"]
	c := db.DB("User").C("user")
	err := c.UpdateId(bson.ObjectIdHex(person._id),person)
	if err!=nil{
		log.Printf("Sin cambios")
		log.Fatal(err)
	}
	log.Printf("Exito")
	db.Close();
}
func init(){
}
func main(){
	//fs := http.FileServer(http.Dir("Resources"))
	go print10000numbers("world")
	print10000numbers("hello")
	/*router := mux.NewRouter()
	router.HandleFunc("/People",GetPeopleHandler).Methods("GET")
	router.HandleFunc("/InsertPeople",InsertPersonHandler).Methods("POST")
	router.HandleFunc("/UpdatePeople/{id}",UpdatePersonHandler).Methods("POST")
	router.HandleFunc("/DeletePeople/{id}",DeletePersonHandler).Methods("DELETE")
	router.Handle("/Resources/Angular/angular.js",http.StripPrefix("/Resources", fs))
	router.Handle("/Resources/JS/AppController.js",http.StripPrefix("/Resources", fs))
	tpl = template.Must(template.ParseGlob("Views/*"))
	fmt.Println("Listening")
	router.HandleFunc("/",chargeHtml).Methods("GET")
	http.ListenAndServe(":8081",router)*/
}

func chargeHtml(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"index.html",nil)	
}