package campay

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type stateResponse struct {
	Reference string `json:"reference"`
	Status    string `json:"status"`
}

func CheckStatus(apiKey, ref string) stateResponse {

	getUrl := fmt.Sprintf("https://demo.campay.net/api/transaction/%s/", ref)

	getReq, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	getReq.Header.Set("Authorization", "Token "+apiKey)
	getReq.Header.Set("Content-Type", "application/json")

	time.Sleep(20 * time.Second)

	// sending the get request
	getResp, err := http.DefaultClient.Do(getReq)
	if err != nil {
		log.Fatal(err)
	}
	// defer getResp.Body.Close()
	defer func() {
    if err := getResp.Body.Close(); 
	err != nil {
        log.Println("Error closing response body:", err)
    }
}()



	//read the response body from get requestapiKey
	var statusResponse stateResponse
	// json.NewDecoder(getReq.Body).Decode(&statusResponse)
	if err := json.NewDecoder(getReq.Body).Decode(&statusResponse); 
	err != nil {
    log.Println("JSON decode error:", err)
}

	return statusResponse

}
