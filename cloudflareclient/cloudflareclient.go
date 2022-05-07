package cloudflareclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type UpdateRecordBody struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Ttl     int    `json:"ttl"`
}

type zoneInfoResponse struct {
	Success bool
	Result  []ZoneResult
}

type recordInfoResponse struct {
	Success bool
	Result  []RecordResult
}

type UpdateUpdateResponse struct {
	Success bool
}

type RecordResult struct {
	Id   string
	Name string
	Type string
}

type ZoneResult struct {
	Id   string
	Name string
}

type CloudFlareConnectionInfo struct {
	Email  string
	ApiKey string
}

type NewRecordInfo struct {
	Domain string
	Tld    string
	Ip     string
}

func cloudflareUrl() string {
	return "https://api.cloudflare.com/client/v4"
}

func UpdateRemoteRecord(connectionInfo CloudFlareConnectionInfo, recordInfo NewRecordInfo, zoneId string, recordId string) UpdateUpdateResponse {
	requestPath := fmt.Sprintf("zones/%s/dns_records/%s", zoneId, recordId)
	requestBody := UpdateRecordBody{
		Type:    "A",
		Name:    recordInfo.Domain,
		Content: recordInfo.Ip,
		Ttl:     1,
	}

	phaseJson, err := json.Marshal(requestBody)

	if err != nil {
		log.Fatalln("Unable to phrase update record request body")
	}

	log.Printf("Update record with follow data %s", phaseJson)

	resp, err := newCloudFlareRequest("PUT", requestPath, connectionInfo, bytes.NewBuffer(phaseJson))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var response UpdateUpdateResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Success {
		log.Println("Successfully Update Record information")
	} else {
		log.Fatalln("Fail to Update cloudflare record infomation")
	}

	return response
}

func FetchRecordInfo(connectionInfo CloudFlareConnectionInfo, zoneId string) recordInfoResponse {
	requestPath := fmt.Sprintf("zones/%s/dns_records", zoneId)
	resp, err := newCloudFlareRequest("GET", requestPath, connectionInfo, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	var response recordInfoResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Success {
		log.Println("Successfully get Record information")
	} else {
		log.Println("Fail to get cloudflare record infomation exiting")
		log.Fatalln(response)
	}

	return response

}

func FetchZoneInfo(connectionInfo CloudFlareConnectionInfo) zoneInfoResponse {
	resp, err := newCloudFlareRequest("GET", "zones", connectionInfo, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var response zoneInfoResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if response.Success {
		log.Println("Successfully get Zone information")
	} else {
		log.Println("Fail to get cloudflare zone infomation exiting")
		log.Fatalln(response)
	}

	return response
}

func newCloudFlareRequest(method string, requestPath string, connection CloudFlareConnectionInfo, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", cloudflareUrl(), requestPath)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest(method, url, body)

	req.Header.Set("X-Auth-Email", connection.Email)
	req.Header.Set("X-Auth-Key", connection.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}
