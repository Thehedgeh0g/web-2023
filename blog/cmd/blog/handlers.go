package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	TopBlock           []topBlocks
	Themes             []themesData
	FeaturedPostsTitle string
	FeaturedPosts      []*featuredPostsBlock
	RecentPostsTitle   string
	RecentPosts        []*recentPostsBlock
	DownBlock          []downBlocks
}

type postPage struct {
	Header    []headerData
	Post      []postData
	DownBlock []downBlocks
}

type topBlocks struct {
	Header  []headerData
	Tagline []taglineData
}

type headerData struct {
	Escape    string
	HeaderNav []headerNavData
}

type headerNavData struct {
	One    string
	OneUrl string
	Two    string
	Three  string
	Four   string
}

type taglineData struct {
	TaglineHeader     string
	TaglineText       string
	TaglineButtonText string
}

type themesData struct {
	First  string
	Second string
	Third  string
	Fourth string
	Fiveth string
	Sixth  string
}

type featuredPostsBlock struct {
	PostID      string `db:"post_id"`
	Theme       string `db:"theme"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImage string `db:"author_url"`
	Date        string `db:"publish_date"`
	Background  string `db:"image_url"`
	PostURL     string
}

type recentPostsBlock struct {
	PostID      string `db:"post_id"`
	Theme       string `db:"theme"`
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImage string `db:"author_url"`
	Date        string `db:"publish_date"`
	Background  string `db:"image_url"`
	PostURL     string
}

type downBlocks struct {
	Feedback []feedBackData
	Footer   []footerData
}

type feedBackData struct {
	FeedbackTitle    string
	InputPlaceholder string
	InputButtonText  string
}

type footerData struct {
	Escape    string
	FooterNav []footerNavData
}

type footerNavData struct {
	One   string
	Two   string
	Three string
	Four  string
}

type postData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	PostPicture string `db:"image_url"`
	Content     string `db:"content"`
}

type adminpage struct {
	Header   []headeradmindata
	MainTop  []maintopdata
	MainInfo []maininfodata
	Content  []contentdata
}

type headeradmindata struct {
	Escape string
	Avatar string
	Exit   string
}

type maintopdata struct {
	Title    string
	Subtitle string
	Button   string
}

type maininfodata struct {
	Title   string
	Fields  []fieldsdata
	Preview []previewdata
}

type fieldsdata struct {
	Theme         string
	Title         string
	Description   string
	Author        string
	AuthorPhoto   string
	AuthorImage   string
	Upload        string
	Date          string
	TitleImage    string
	BigImageURL   string
	SmallImageURL string
	BigNote       string
	SmallNote     string
}

type previewdata struct {
	Article  []articledata
	PostCard []postcarddata
}

type articledata struct {
	Label      string
	FrameURL   string
	Title      string
	Subtitle   string
	Background string
}

type postcarddata struct {
	Label       string
	FrameURL    string
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	Author      string
	Date        string
}

type contentdata struct {
	Title   string
	Comment string
}

type loginpage struct {
	Header []headerlogindata
	Main   []mainlogindata
}

type headerlogindata struct {
	Escape string
	Title  string
}

type mainlogindata struct {
	Title  string
	Email  string
	Pass   string
	Button string
}

type createPostRequest struct {
	Theme           string `json:"theme"`
	Title           string `json:"title"`
	SubTitle        string `json:"subtitle"`
	AuthorName      string `json:"authorname"`
	AuthorPhoto     string `json:"authorphoto"`
	AuthorPhotoName string `json:"authorphotoname"`
	Data            string `json:"data"`
	BigImage        string `json:"bigimage"`
	BigImageName    string `json:"bigimagename"`
	SmallImage      string `json:"smallimage"`
	SmallImageName  string `json:"smallimagename"`
	Content         string `json:"content"`
}

type UserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Userdata struct {
	UserId   string
	Email    string
	Password string
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		FeaturedPosts, err := FeaturedPosts(db)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err)
			return
		}

		RecentPosts, err := RecentPosts(db)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := indexPage{
			TopBlock:           TopBlock(),
			Themes:             Themes(),
			FeaturedPostsTitle: "Featured Posts",
			FeaturedPosts:      FeaturedPosts,
			RecentPosts:        RecentPosts,
			RecentPostsTitle:   "Most Recent",
			DownBlock:          DownBlock(),
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Order not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := postPage{
			Header:    Header(),
			Post:      post,
			DownBlock: DownBlock(),
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func admin(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("pages/admin.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
		data := adminpage{
			Header:   headeradmin(),
			MainTop:  maintop(),
			MainInfo: maininfo(),
			Content:  content(),
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

	}
}

func login(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/login.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := loginpage{
		Header: headerlogin(),
		Main:   mainlogin(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func TopBlock() []topBlocks {
	return []topBlocks{
		{
			Header:  Header(),
			Tagline: Tagline(),
		},
	}
}

func Header() []headerData {
	return []headerData{
		{
			Escape:    "Escape.",
			HeaderNav: HeaderNav(),
		},
	}
}

func HeaderNav() []headerNavData {
	return []headerNavData{
		{
			One:    "HOME",
			OneUrl: "/home",
			Two:    "CATEGORIES",
			Three:  "ABOUT",
			Four:   "CONTACT",
		},
	}
}

func Tagline() []taglineData {
	return []taglineData{
		{
			TaglineHeader:     "Let's do it together.",
			TaglineText:       "We travel the world in search of stories. Come along for the ride.",
			TaglineButtonText: "View Latest Posts",
		},
	}
}

func Themes() []themesData {
	return []themesData{
		{
			First:  "Nature",
			Second: "Photography",
			Third:  "Relaxation",
			Fourth: "Vacation",
			Fiveth: "Travel",
			Sixth:  "Adventure",
		},
	}
}

func FeaturedPosts(db *sqlx.DB) ([]*featuredPostsBlock, error) {
	const query = `
		SELECT
		  post_id,
		  theme,
		  title,
		  subtitle,
		  author,
		  author_url,
		  publish_date,
		  image_url
		FROM
		  post
		WHERE featured = 1
	`

	var featuredposts []*featuredPostsBlock

	err := db.Select(&featuredposts, query)
	if err != nil {
		return nil, err
	}

	for _, post := range featuredposts {
		post.PostURL = "/post/" + post.PostID // Формируем исходя из ID ордера в базе
	}

	fmt.Println(featuredposts)

	return featuredposts, nil
}

func RecentPosts(db *sqlx.DB) ([]*recentPostsBlock, error) {
	const query = `
		SELECT
		  post_id,
		  theme,
		  title,
		  subtitle,
		  author,
		  author_url,
		  publish_date,
		  image_url
		FROM
		  post
		WHERE featured = 0
	`

	var mostrecent []*recentPostsBlock

	err := db.Select(&mostrecent, query)
	if err != nil {
		return nil, err
	}

	for _, post := range mostrecent {
		post.PostURL = "/post/" + post.PostID // Формируем исходя из ID ордера в базе
	}

	fmt.Println(mostrecent)

	return mostrecent, nil
}

func postByID(db *sqlx.DB, postID int) ([]postData, error) {
	const query = `
		SELECT
		  title,
		  subtitle,
		  image_url,
		  content
		FROM
		  post
	    WHERE
		  post_id = ?
	`

	var post []postData

	err := db.Select(&post, query, postID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func DownBlock() []downBlocks {
	return []downBlocks{
		{
			Feedback: Feedback(),
			Footer:   Footer(),
		},
	}
}

func Feedback() []feedBackData {
	return []feedBackData{
		{
			FeedbackTitle:    "Stay in Touch",
			InputPlaceholder: "Enter your email address",
			InputButtonText:  "Submit",
		},
	}
}

func Footer() []footerData {
	return []footerData{
		{
			Escape:    "Escape.",
			FooterNav: FooterNav(),
		},
	}
}

func FooterNav() []footerNavData {
	return []footerNavData{
		{
			One:   "HOME",
			Two:   "CATEGORIES",
			Three: "ABOUT",
			Four:  "CONTACT",
		},
	}
}

func headeradmin() []headeradmindata {
	return []headeradmindata{
		{
			Escape: "../static/sources/Logo.svg",
			Avatar: "../static/sources/Avatar.svg",
			Exit:   "../static/sources/log-out.svg",
		},
	}
}

func maintop() []maintopdata {
	return []maintopdata{
		{
			Title:    "New Post",
			Subtitle: "Fill out the form bellow and publish your article",
			Button:   "Publish",
		},
	}
}

func maininfo() []maininfodata {
	return []maininfodata{
		{
			Title:   "Main Information",
			Fields:  fields(),
			Preview: preview(),
		},
	}
}

func fields() []fieldsdata {
	return []fieldsdata{
		{
			Theme:         "Theme",
			Title:         "Title",
			Description:   "Short description",
			Author:        "Author Name",
			AuthorPhoto:   "Author Photo",
			AuthorImage:   "../static/sources/Avatar.svg",
			Upload:        "Upload",
			Date:          "Publish Date",
			TitleImage:    "Hero image",
			BigImageURL:   "../static/sources/hero_image_big.png",
			SmallImageURL: "../static/sources/hero_image_small.png",
			BigNote:       "Size up to 10mb. Format: png, jpeg, gif.",
			SmallNote:     "Size up to 5mb. Format: png, jpeg, gif.",
		},
	}
}

func preview() []previewdata {
	return []previewdata{
		{
			Article:  article(),
			PostCard: postcard(),
		},
	}
}

func article() []articledata {
	return []articledata{
		{
			Label:      "Article preview",
			FrameURL:   "../static/sources/preview-browser-imitation.png",
			Title:      "New Post",
			Subtitle:   "Please, enter any description",
			Background: "../static/sources/preview.png",
		},
	}
}

func postcard() []postcarddata {
	return []postcarddata{
		{
			Label:       "Post card preview",
			FrameURL:    "../static/sources/preview-browser-imitation.png",
			Background:  "../static/sources/preview.png",
			Title:       "New Post",
			Subtitle:    "Please, enter any description",
			AuthorImage: "../static/sources/Avatar.svg",
			Author:      "Enter author name",
			Date:        "4/19/2023",
		},
	}
}

func content() []contentdata {
	return []contentdata{
		{
			Title:   "Content",
			Comment: "Post content (plain text)",
		},
	}
}

func headerlogin() []headerlogindata {
	return []headerlogindata{
		{
			Escape: "../static/sources/login-logo.svg",
			Title:  "Log in to start creating",
		},
	}
}

func mainlogin() []mainlogindata {
	return []mainlogindata{
		{
			Title:  "Log In",
			Email:  "Email",
			Pass:   "Password",
			Button: "Log In",
		},
	}
}
func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "1Error", 500)
			log.Println(err.Error())
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "2Error", 700)
			log.Println(err.Error())
			return
		}

		fileData := req.AuthorPhoto[strings.IndexByte(req.AuthorPhoto, ',')+1:]

		authorImg, err := base64.StdEncoding.DecodeString(fileData)
		if err != nil {
			http.Error(w, "img", 600)
			log.Println(err.Error())
			return
		}

		fileName := "static/sources/" + req.AuthorPhotoName

		fileAuthor, err := os.Create(fileName)
		if err != nil {
			http.Error(w, "img", 400)
			log.Println(err.Error())
			return
		}

		_, err = fileAuthor.Write(authorImg)

		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileData = req.BigImage[strings.IndexByte(req.BigImage, ',')+1:]

		bigImg, err := base64.StdEncoding.DecodeString(fileData)

		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileBig, err := os.Create("static/sources/" + req.BigImageName)

		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		_, err = fileBig.Write(bigImg)

		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileData = req.SmallImage[strings.IndexByte(req.SmallImage, ',')+1:]

		smallImg, err := base64.StdEncoding.DecodeString(fileData)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		fileSmall, err := os.Create("static/sources/" + req.SmallImageName)

		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		_, err = fileSmall.Write(smallImg)
		if err != nil {
			http.Error(w, "img", 500)
			log.Println(err.Error())
			return
		}

		err = saveOrder(db, req)
		if err != nil {
			http.Error(w, "bd", 500)
			log.Println(err.Error())
			return
		}
		return
	}
}

func saveOrder(db *sqlx.DB, req createPostRequest) error {
	const query = `
		INSERT INTO
			post
		(
			theme,
			title,
			subtitle,
			author,
			author_url,
			publish_date,
			image_url,
			content
		)
		VALUES
		(
			?,
			?,
			?,
			?,
			CONCAT('../static/sources/', ?),
			?,
			CONCAT('../static/sources/', ?),
			?
		)
	`

	_, err := db.Exec(query, req.Theme, req.Title, req.SubTitle, req.AuthorName, req.AuthorPhotoName, req.Data, req.BigImageName, req.Content)
	return err
}

func searchUser(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err.Error())
		}

		var req UserRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println(req.Email, ' ', req.Password)
		user, err := getUser(db, req)

		if err != nil {
			http.Error(w, "Incorect email or password", 401)
			log.Println("Incorect email or password")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "authCookieName",
			Value:   fmt.Sprint(user.UserId),
			Path:    "/",
			Expires: time.Now().AddDate(0, 0, 1),
		})

		log.Println("Cookie setted")
	}
}

func AuthByCookie(db *sqlx.DB, w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("authCookieName")

	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "No auth cookie passed", 401)
			log.Println(err)
			return err
		}
		http.Error(w, "Internal Server Error", 500)
		log.Println(err)
		return err
	}

	userIDStr := cookie.Value

	err = search(db, userIDStr)
	log.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func getUser(db *sqlx.DB, req UserRequest) (*Userdata, error) {
	const query = `
	SELECT
	  user_id,
	  email,
	  password
  	FROM
	  user
  	WHERE
	  email = ? AND
	  password = ?
	`
	row := db.QueryRow(query, req.Email, req.Password)
	user := new(Userdata)
	err := row.Scan(&user.UserId, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func search(db *sqlx.DB, UserID string) error {
	const query = `
	SELECT
	  post_id,
	  email,
	  password
	FROM
	  user
	WHERE
	  post_id = ?
	`

	row := db.QueryRow(query, UserID)
	user := new(Userdata)
	err := row.Scan(&user.UserId, &user.Email, &user.Password)
	fmt.Println(user, UserID)
	if err != nil {
		fmt.Println("fdf")
		return err
	}

	fmt.Println(UserID)
	return nil
}

func deleteUser(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:    "authCookieName",
			Path:    "/",
			Expires: time.Now().AddDate(0, 0, -1),
		})

		return
	}
}
