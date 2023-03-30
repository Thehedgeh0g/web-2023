package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	TopBlock  []topBlocks
	Themes    []themesData
	Container []containerBlocks
	DownBlock []downBlocks
}

type postPage struct {
	Header      []headerData
	PostTagline []postTaglineData
	Post        []postData
	DownBlock   []downBlocks
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
	One   string
	Two   string
	Three string
	Four  string
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

type containerBlocks struct {
	FeaturedPosts []featuredPostsBlock
	RecentPosts   []recentPostsBlock
}

type featuredPostsBlock struct {
	Title        string
	TheRoadAhead []theRoadAheadData
	FromTopDown  []fromTopDownData
}

type theRoadAheadData struct {
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
}

type fromTopDownData struct {
	Background  string
	Button      string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
}

type recentPostsBlock struct {
	Title             string
	StillStandingTall []stillStandingTallData
	SunnySideUp       []sunnySideUpData
	WaterFalls        []waterFallsData
	ThroughTheMist    []throughTheMistData
	AwakenEarly       []awakenEarlyData
	TryItAlways       []tryItAlwaysData
}

type stillStandingTallData struct {
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
}

type sunnySideUpData struct {
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
}

type waterFallsData struct {
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
}

type throughTheMistData struct {
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
}

type awakenEarlyData struct {
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
}

type tryItAlwaysData struct {
	Background  string
	Title       string
	Subtitle    string
	AuthorImage string
	AuthorName  string
	Data        string
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

type postTaglineData struct {
	Title       string
	Subtitle    string
	PostPicture string
}

type postData struct {
	ParFirst  string
	ParSecond string
	ParThird  string
	ParFourth string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	Data := indexPage{
		TopBlock:  TopBlock(),
		Themes:    Themes(),
		Container: Container(),
		DownBlock: DownBlock(),
	}

	err = ts.Execute(w, Data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/the-road-ahead.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	Data := postPage{
		Header:      Header(),
		PostTagline: PostTagline(),
		Post:        Post(),
		DownBlock:   DownBlock(),
	}

	err = ts.Execute(w, Data) // Заставляем шаблонизатор вывести шаблон в тело ответа
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
			One:   "HOME",
			Two:   "CATEGORIES",
			Three: "ABOUT",
			Four:  "CONTACT",
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

func Container() []containerBlocks {
	return []containerBlocks{
		{
			FeaturedPosts: FeaturedPosts(),
			RecentPosts:   RecentPosts(),
		},
	}
}

func FeaturedPosts() []featuredPostsBlock {
	return []featuredPostsBlock{
		{
			Title:        "Featured Posts",
			TheRoadAhead: TheRoadAhead(),
			FromTopDown:  FromTopDown(),
		},
	}
}

func TheRoadAhead() []theRoadAheadData {
	return []theRoadAheadData{
		{
			Background:  "../static/Sources/the-road-ahead-background.jpg",
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			AuthorImage: "../static/Sources/Mat-Vogles.svg",
			AuthorName:  "Mat Vogels",
			Data:        "September 25, 201",
		},
	}
}

func FromTopDown() []fromTopDownData {
	return []fromTopDownData{
		{
			Background:  "../static/Sources/from-top-down-background.jpg",
			Button:      "ADVENTURE",
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you've never been before.",
			AuthorImage: "../static/Sources/William-Wong.svg",
			AuthorName:  "William Wong",
			Data:        "September 25, 2015",
		},
	}
}

func RecentPosts() []recentPostsBlock {
	return []recentPostsBlock{
		{
			Title:             "Most Recent",
			StillStandingTall: stillStandingTall(),
			SunnySideUp:       SunnySideUp(),
			WaterFalls:        WaterFalls(),
			ThroughTheMist:    ThroughTheMist(),
			AwakenEarly:       AwakenEarly(),
			TryItAlways:       TryItAlways(),
		},
	}
}

func stillStandingTall() []stillStandingTallData {
	return []stillStandingTallData{
		{
			Background:  "../static/Sources/still-standing-tall-background.jpg",
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			AuthorImage: "../static/Sources/William-Wong.svg",
			AuthorName:  "William Wong",
			Data:        "9/25/2015",
		},
	}
}

func SunnySideUp() []sunnySideUpData {
	return []sunnySideUpData{
		{
			Background:  "../static/Sources/sunny-side-up-background.jpg",
			Title:       "Sunny Side Up",
			Subtitle:    "No place is ever as bad as they tell you it's going to be.",
			AuthorImage: "../static/Sources/Mat-Vogles.svg",
			AuthorName:  "Mat Vogels",
			Data:        "9/25/2015",
		},
	}
}

func WaterFalls() []waterFallsData {
	return []waterFallsData{
		{
			Background:  "../static/Sources/water-falls-background.jpg",
			Title:       "Water Falls",
			Subtitle:    "We travel not to Escape life, but for life not to Escape us.",
			AuthorImage: "../static/Sources/Mat-Vogles.svg",
			AuthorName:  "Mat Vogels",
			Data:        "9/25/2015",
		},
	}
}

func ThroughTheMist() []throughTheMistData {
	return []throughTheMistData{
		{
			Background:  "../static/Sources/trough-the-mist-background.jpg",
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			AuthorImage: "../static/Sources/William-Wong.svg",
			AuthorName:  "William Wong",
			Data:        "9/25/2015",
		},
	}
}

func AwakenEarly() []awakenEarlyData {
	return []awakenEarlyData{
		{
			Background:  "../static/Sources/awaken-early-background.jpg",
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			AuthorImage: "../static/Sources/Mat-Vogles.svg",
			AuthorName:  "Mat Vogels",
			Data:        "9/25/2015",
		},
	}
}

func TryItAlways() []tryItAlwaysData {
	return []tryItAlwaysData{
		{
			Background:  "../static/Sources/try-it-always-background.jpg",
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only One page.",
			AuthorImage: "../static/Sources/Mat-Vogles.svg",
			AuthorName:  "Mat Vogels",
			Data:        "9/25/2015",
		},
	}
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

func PostTagline() []postTaglineData {
	return []postTaglineData{
		{
			Title:       "The Road Ahead.",
			Subtitle:    "The road ahead might be paved - it might not be.",
			PostPicture: "../static/Sources/the-road-ahead-background.jpg",
		},
	}
}

func Post() []postData {
	return []postData{
		{
			ParFirst:  "Dark spruce forest frowned on either side the frozen waterway. The trees had been stripped by a recent wind of their white covering of frost, and they seemed to lean towards each other, black and ominous, in the fading light. A vast silence reigned over the land. The land itself was a desolation, lifeless, without movement, so lone and cold that the spirit of it was not even that of sadness. There was a hint in it of laughter, but of a laughter more terrible than any sadness—a laughter that was mirthless as the smile of the sphinx, a laughter cold as the frost and partaking of the grimness of infallibility. It was the masterful and incommunicable wisdom of eternity laughing at the futility of life and the effort of life. It was the Wild, the savage, frozen-hearted Northland Wild.",
			ParSecond: "But there was life, abroad in the land and defiant. Down the frozen waterway toiled a string of wolfish dogs. Their bristly fur was rimed with frost. Their breath froze in the air as it left their mouths, spouting forth in spumes of vapour that settled upon the hair of their bodies and formed into crystals of frost. Leather harness was on the dogs, and leather traces attached them to a sled which dragged along behind. The sled was without runners. It was made of stout birch-bark, and its full surface rested on the snow. The front end of the sled was turned up, like a scroll, in order to force down and under the bore of soft snow that surged like a wave before it. On the sled, securely lashed, was a long and narrow oblong box. There were other things on the sled—blankets, an axe, and a coffee-pot and frying-pan; but prominent, occupying most of the space, was the long and narrow oblong box.",
			ParThird:  "In advance of the dogs, on wide snowshoes, toiled a man. At the rear of the sled toiled a Second man. On the sled, in the box, lay a third man whose toil was over,—a man whom the Wild had conquered and beaten down until he would never move nor struggle again. It is not the way of the Wild to like movement. Life is an offence to it, for life is movement; and the Wild aims always to destroy movement. It freezes the water to prevent it running to the sea; it drives the sap out of the trees till they are frozen to their mighty hearts; and most ferociously and terribly of all does the Wild harry and crush into submission man—man who is the most restless of life, ever in revolt against the dictum that all movement must in the end come to the cessation of movement.",
			ParFourth: "But at front and rear, unawed and indomitable, toiled the two men who were not yet dead. Their bodies were covered with fur and soft-tanned leather. Eyelashes and cheeks and lips were so coated with the crystals from their frozen breath that their faces were not discernible. This gave them the seeming of ghostly masques, undertakers in a spectral world at the funeral of some ghost. But under it all they were men, penetrating the land of desolation and mockery and silence, puny adventurers bent on colossal adventure, pitting themselves against the might of a world as remote and alien and pulseless as the abysses of space.",
		},
	}
}
