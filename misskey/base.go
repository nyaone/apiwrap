package misskey

import (
	"apiwrap/config"
	"apiwrap/global"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//type Error_Response struct {
//	Error struct {
//		Message string `json:"message"`
//		Code    string `json:"code"`
//		ID      string `json:"id"`
//		Kind    string `json:"kind"`
//	} `json:"error"`
//}

func PostAPIRequest(
	apiEndpointPath string, reqBody any,
) (*any, int, error) {
	// Prepare request
	apiEndpoint := fmt.Sprintf("%s/api%s", config.Config.Misskey.Instance, apiEndpointPath)

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		global.Logger.Errorf("Failed to marshall request body with error: %v", err)
		return nil, http.StatusBadRequest, err
	}

	global.Logger.Debugf("Request URL: %s , body: %s", apiEndpoint, reqBodyBytes)

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewReader(reqBodyBytes))
	if err != nil {
		global.Logger.Errorf("Failed to prepare request with error: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Do request
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		global.Logger.Errorf("Failed to finish request with error: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	// Parse response
	var resBody any
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		global.Logger.Errorf("Failed to decode response body with error: %v", err)
		return nil, res.StatusCode, err
	}

	return &resBody, res.StatusCode, nil

}
