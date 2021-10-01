package twit

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/dghubble/gologin/v2/twitter"
	"github.com/dghubble/sessions"
)

const (
	sessionName     = "example-twtter-app"
	sessionSecret   = "example cookie signing secret"
	sessionUserKey  = "twitterID"
	sessionUsername = "twitterUsername"
)

// sessionStore encodes and decodes session data stored in signed cookies
var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

func IssueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		twitterUser, err := twitter.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(twitterUser)
		// 2. Implement a success handler to issue some form of session
		session := sessionStore.New(sessionName)
		session.Values[sessionUserKey] = twitterUser.ID
		session.Values[sessionUsername] = twitterUser.ScreenName
		session.Save(w)
		http.Redirect(w, req, "/", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

// func CreateNewTweet(w http.ResponseWriter, req *http.Request) {
// 	SendTweet(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
// }

func RandomString(n int) string {
	var output string

	ascii := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ascii_arr := strings.Split(ascii, "")
	for i := 1; i < n; i++ {
		randInt := rand.Intn(len(ascii_arr))
		output = output + ascii_arr[randInt]
	}
	return output
}

// func SendTweet(accessToken string, accessSecret string) {
// 	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
// 	token := oauth1.NewToken(accessToken, accessSecret)
// 	httpClient := config.Client(oauth1.NoContext, token)

// 	client := twitt.NewClient(httpClient)
// 	tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		bodyString := string(bodyBytes)
// 		fmt.Println(bodyString)
// 	}
// 	println(tweet, err)

// }
