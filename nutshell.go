package gonutshell

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Duration(2*time.Second))
}

func AddContact(username, apikey, contactName, contactEmail string) (userid string, err error) {
	str := `{"jsonrpc":"2.0","method":"newContact","params":{"contact":{"name":"` + contactName + `", "email":"` + contactEmail + `"}},"id": "mark@crowdstitch.com"}`
	qbodyreader := strings.NewReader(str)
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://app01.nutshell.com/api/v1/json", qbodyreader)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(username, apikey)
	contenttype := "application/json"
	req.Header.Add("Content-Type", contenttype)
	resp, err := client.Do(req)
	fmt.Printf("Response: %v, Error: %v\n", resp, err)
	fmt.Println("Respnse headers:", resp.Header)
	rbb, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Respnse body:", string(rbb))
	fmt.Println("Request content length: ", resp.Request.ContentLength)

	var j map[string]interface{}
	err = json.Unmarshal(rbb, &j)
	fmt.Printf("json:%v\n", j)

	result := j["result"].(map[string]interface{})
	userid = fmt.Sprintf("%v", result["id"])
	return userid, err
}

func AddLead(username, apikey, leadtitle, contactid string) error {
	str := `{"jsonrpc":"2.0","method":"newLead","params":{"lead":{"contacts":[{"id":` + contactid + `}]}},"id": "mark@crowdstitch.com"}`
	fmt.Printf("Payload: %v\n", str)
	qbodyreader := strings.NewReader(str)
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://app01.nutshell.com/api/v1/json", qbodyreader)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, apikey)
	contenttype := "application/json"
	req.Header.Add("Content-Type", contenttype)
	resp, err := client.Do(req)
	fmt.Printf("Response: %v, Error: %v\n", resp, err)
	fmt.Println("Respnse headers:", resp.Header)
	rbb, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Respnse body:", string(rbb))
	fmt.Println("Request content length: ", resp.Request.ContentLength)

	var j map[string]interface{}
	err = json.Unmarshal(rbb, &j)
	fmt.Printf("json:%v\n", j)

	return err
}
