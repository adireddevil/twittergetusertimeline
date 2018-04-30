package getusertimeline

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"fmt"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestGDriveCreateFile_InvalidUser(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("consumerKey", "NPsc5Cz49sLMkUEwM4gclc6bc")
	tc.SetInput("consumerSecret", "0OelZ71gGz3CA9vuA5HJwX2tO27XHZkUIZV5scZbM0lK1btjqr")
	tc.SetInput("accessToken", "990616654353645569-M5rp4NsPuWavBKcT25BOZG4c5sTwf7a")
	tc.SetInput("accessTokenSecret", "VfmVsLP3ZVp7kQyUU9bC1Vf4FABNT8sQ2MmQYWAHljKZ8")
	tc.SetInput("sinceId", 10)
	tc.SetInput("pageCount", 10)

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")






}
