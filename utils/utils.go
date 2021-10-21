package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UnmarshalInStruct(r *http.Request, model interface{}) (interface{}, error) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(data, &model)
	return model, nil
}
