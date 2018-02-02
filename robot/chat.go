package robot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
)

func getUserMacAddr() string {
	var hardAddr string
	netitfs, err := net.Interfaces()
	if err != nil {
		hardAddr = "anonymous"
	} else {
		hardAddr = netitfs[0].HardwareAddr.String()
		if hardAddr == "" {
			hardAddr = netitfs[3].HardwareAddr.String()
		}
	}
	return hardAddr
}

func GetRobotResp(info, loc string) ([]byte, error) {
	reqStruct := LocalReq{
		Key:    robotApiKey,
		Info:   info,
		Loc:    loc,
		UserId: getUserMacAddr(),
	}
	reqJson, err := json.Marshal(reqStruct)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(reqJson)
	request, err := http.NewRequest("POST", robotAprUrl, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/jsonl;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}
