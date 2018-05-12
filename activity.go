package twittergetusertimeline

import (
	s "strings"

	"github.com/DipeshTest/allstarsshared/twitter"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-twittergetusertimeline")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	consumerKey := s.TrimSpace(context.GetInput("consumerKey").(string))
	consumerSecret := s.TrimSpace(context.GetInput("consumerSecret").(string))
	accessToken := s.TrimSpace(context.GetInput("accessToken").(string))
	accessTokenSecret := s.TrimSpace(context.GetInput("accessTokenSecret").(string))
	sinceId := (context.GetInput("sinceId").(int))
	pageCount := (context.GetInput("pageCount").(int))

	if len(consumerKey) == 0 {

		context.SetOutput("statusCode", "101")

		context.SetOutput("message", "Consumer Key field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else if len(consumerSecret) == 0 {

		context.SetOutput("statusCode", "102")

		context.SetOutput("message", "Consumer Secret field is blank")

	} else if len(accessToken) == 0 {

		context.SetOutput("statusCode", "103")

		context.SetOutput("message", "Access Token field is blank")

	} else if len(accessTokenSecret) == 0 {

		context.SetOutput("statusCode", "104")

		context.SetOutput("message", "Access Token Secret field is blank")

	} else {
		code, msg := twitter.GetUserTimeline(consumerKey, consumerSecret, accessToken, accessTokenSecret, sinceId, pageCount)
		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)
	}

	return true, err
}
