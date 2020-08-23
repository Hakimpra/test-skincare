package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	sessions "github.com/kataras/go-sessions"
	"golang.org/x/crypto/bcrypt"
	// "os"
)

var db *sql.DB
var err error

type user struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
	Email     string
}
type Treatment  struct {
	ID        string
	Title string 
	Harga  uint32 
	Keterangan  string 
}
type PesanTreatment  struct {
	ID        string
	// Username string 
	// Title string 
	User_id string 
	Treatment_id  uint32 
	Total_bayar  string 
	// Keterangan string 
}
type ResponseTreatment struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Treatment 
}
type ResponsePesanTreatment struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []PesanTreatment 
}
func connect_db() {
	db, err = sql.Open("mysql", "root:wakidij@tcp(127.0.0.1:3306)/test-skincare2")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func routes() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/viewtreatment", viewtreatment)
	http.HandleFunc("/pesantreatment", pesantreatment)
	http.HandleFunc("/viewpesantreatment", viewpesantreatment)
}

func main() {
	connect_db()
	routes()

	defer db.Close()

	fmt.Println("Server running on port :8001")
	http.ListenAndServe(":8001", nil)
}

type statusRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
}

func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		password,
		email
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.Password,
			&users.Email,
		)
	return users
}
func viewpesantreatment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	var response ResponsePesanTreatment
	var pesantreatment PesanTreatment
	var arr_pesantreatment []PesanTreatment


	err := r.ParseMultipartForm(4096)

	id := r.FormValue("id")

	/*rows, err := db.Query("SELECT pesantreatment.id, users.username, treatment.title, pesantreatment.total_bayar, treatment.keterangan from pesantreatment left JOIN treatment ON treatment.id = pesantreatment.treatment_id left JOIN users ON users.id = pesantreatment.user_id where pesantreatment.user_id=?",
		id,
	)*/
	rows, err := db.Query("SELECT * from pesantreatment where id=?",
		id,
	)

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		// if err := rows.Scan(&pesantreatment.ID, &pesantreatment.Username, &pesantreatment.Title, &pesantreatment.Total_bayar, &pesantreatment.Keterangan); err != nil {
		if err := rows.Scan(&pesantreatment.ID, &pesantreatment.User_id, &pesantreatment.Treatment_id, &pesantreatment.Total_bayar); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_pesantreatment = append(arr_pesantreatment, pesantreatment)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_pesantreatment

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
func pesantreatment(w http.ResponseWriter, r *http.Request) {

	var response ResponseTreatment
	

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	user_id := r.FormValue("user_id")
	treatment_id := r.FormValue("treatment_id")
	total_bayar := r.FormValue("total_bayar")

	_, err = db.Exec("INSERT INTO pesantreatment (user_id, treatment_id, total_bayar) values (?,?,?)",
		user_id,
		treatment_id,
		total_bayar,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert pesanan treatment success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
func viewtreatment(w http.ResponseWriter, r *http.Request) {

	var response ResponseTreatment
	var treatment Treatment
	var arr_treatmment []Treatment
	if r.Method != "GET" {
		return
	}

	rows, err := db.Query("Select id,title,harga,keterangan from treatment")
	if err != nil {
		log.Print(err)
	}


	for rows.Next() {
		if err := rows.Scan(&treatment.ID, &treatment.Title, &treatment.Harga, &treatment.Keterangan); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_treatmment = append(arr_treatmment, treatment)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_treatmment

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println(username)
	users := QueryUser(username)

	if (user{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, email=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &email)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				res := statusRes{Status: 200, Msg: "berhasil"}
				json.NewEncoder(w).Encode(res)
			}
		}
	} else {
		res := statusRes{Status: 400, Msg: "Method Must be post"}
		json.NewEncoder(w).Encode(res)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 && checkErr(w, r, err) {
		//check session if avaliabel
		res := statusRes{Status: 200, Msg: "berhasil login session avliabe"}
		json.NewEncoder(w).Encode(res)
	}
	if r.Method != "POST" {
		res := statusRes{Status: 400, Msg: "Method Must be post"}
		json.NewEncoder(w).Encode(res)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	//deskripsi dan compare password
	var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if password_tes == nil {
		//login success
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		// session.Set("name", users.FirstName)
		res := statusRes{Status: 200, Msg: "berhasil login"}
		json.NewEncoder(w).Encode(res)
	} else {
		//login failed
		res := statusRes{Status: 400, Msg: "gagal login"}
		json.NewEncoder(w).Encode(res)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	res := statusRes{Status: 200, Msg: "berhasil logout"}
	json.NewEncoder(w).Encode(res)
}
