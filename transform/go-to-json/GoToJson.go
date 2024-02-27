package go_to_json

import "encoding/json"

// GoToJson 补充 go-to-json 方法, https://transform.tools/json-to-go
func GoToJson(d interface{}) string {
	jsonData, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(jsonData)
}

// Model

type EventInstall struct {
	EventName string `json:"event_name"`

	AppsflyerId    string `json:"appsflyer_id"`
	CustomData     string `json:"custom_data"`
	CustomDeviceId string `json:"customDeviceId,omitempty"`

	IdFV          string `json:"idfv"`
	IdFA          string `json:"idfa"`
	AdvertisingId string `json:"advertising_id"`
	// campaign_id
	AfCId     string `json:"af_c_id"`
	AfAdSetId string `json:"af_adset_id"`
	AfAdSet   string `json:"af_adset"`
	AfAdId    string `json:"af_ad_id"`
}
