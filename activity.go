package testing

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
)

const (
	ConsumerKey       = "T2pxJxJIaPHue1QnmNZ8XkYQq"
	ConsumerSecret    = "p1covApVTzWMJSnX7DeAxsunpa26jbeSNRo3uV6dEN708aET64"
	AccessToken       = "855707869773086720-gLzwKhsunfZ1pahZQMN5e52zX1Z1S0U"
	AccessTokenSecret = "GOGSE6BVmWxbR2yCtIRUfBBNLnAP8v5xmdFc4KmWVwbjH"
)

var logt = logger.GetLogger("activity-testing")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	consumer := oauth.NewConsumer(ConsumerKey,
		ConsumerSecret,
		oauth.ServiceProvider{})
	//NOTE: remove this line or turn off Debug if you don't
	//want to see what the headers look like
	//consumer.Debug(true)
	//Roll your own AccessToken struct
	accessToken := &oauth.AccessToken{Token: AccessToken,
		Secret: AccessTokenSecret}
	twitterEndPoint := "https://api.twitter.com/1.1/statuses/mentions_timeline.json"
	response, err := consumer.Get(twitterEndPoint, nil, accessToken)
	if err != nil {
		log.Fatal(err, response)
	}
	defer response.Body.Close()
	fmt.Println("Response:", response.StatusCode, response.Status)
	logt.Debugf("The Flogo engine produces response ....", response.StatusCode, response.Status)

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	logt.Debugf(string(respBody))
	context.SetOutput("response", "The flogo gives response"+response.Status)
	return true, nil
}
