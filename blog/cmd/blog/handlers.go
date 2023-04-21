package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"strconv"

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

func Post(db *sqlx.DB) ([]postData, error) {
	const query = `
		SELECT
		  title,
		  subtitle,
		  image_url,
		  text
		FROM
		  post
		WHERE featured = 0
	`

	var post []postData

	err := db.Select(&post, query)
	if err != nil {
		return nil, err
	}

	return post, nil
}
