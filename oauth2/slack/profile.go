package slack

import (
	"encoding/json"
	"net/http"
)

// GetProfileData returns profile data for slack
func GetProfileData(token string) (map[string]interface{}, error) {
	c := http.Client{}
	res, err := c.Get(profileEndpoint + "?token=" + token)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	r, err := c.Get("https://slack.com/api/users.info?token=" + token + "&user=" + data["user_id"].(string))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var data2 map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data2)
	if err != nil {
		return nil, err
	}
	return data2, nil
}
