package cvclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"crypto/tls"
	"net/http"
	"strconv"
	"time"
	_ "github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

// THis is CyberVision specific getData, not all API set are same to some
// custom t0-t1 handling
func (mc *CVClient) getData(netURL string, lastSyncUp int64,sync_up_to time.Time, offset int, limit int) ([]byte, error) {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	nextStartTime := lastSyncUp
	var lag time.Duration = 300 * time.Second
	req, err := http.NewRequest("GET", netURL, nil)
	if err != nil {
		logp.Info("Failed to connect CV API %s", err.Error())
		return nil, err
	}
	//req.Header.Add("X-Cisco-CV-API-Key", mc.Key)

	q := req.URL.Query()

	//startTime = startTime - 600
	q.Add("token",mc.Key)

	//endTime := time.Now().Add(0 - lag).Unix()
	endTime := sync_up_to.Add(lag).Unix()

	if nextStartTime >0 {
		tm := time.Unix(nextStartTime, 0).UTC()
		q.Add("start", tm.Format("2006-01-02 15:04:05"))
	}else {

		//tm,_ := time.Parse(time.RFC3339, "2000-01-01T00:00:00:000Z")
		//tm := time.Unix(tm_x, 0).UTC()
		q.Add("start", "2000-01-01 00:00:00")
	}
	//
	end_tm := time.Unix(endTime, 0).UTC()
	q.Add("end", end_tm.Format("2006-01-02 15:04:05"))

	if limit!=0{
		q.Add("limit", strconv.Itoa(limit) )
	}
	if offset !=0{
		logp.Info("Append Offset: %s", strconv.Itoa(offset))
		q.Add("offset", strconv.Itoa(offset))
	}

	req.URL.RawQuery = q.Encode()
	logp.Info("Calling API URL %+v", req.URL)
	resp, err := client.Do(req)
	if err != nil {
		logp.Info("Failed to connect CV API %s", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	client.CloseIdleConnections()
	//in lower version of golang CloseIdleConnections is not supported.
	logp.Info("%s", string(body[:]))
	return body, err
}

func (mc *CVClient) GetComponents(lastSyncUp int64,sync_up_to time.Time) (CVComponents, error) {
	netURL := fmt.Sprintf("%s/api/1.0/component", mc.URL)
	var allComponents CVComponents
	var offset int =0
	limit :=mc.Limit
	for {
		var components CVComponents
		body, err := mc.getData(netURL,lastSyncUp,sync_up_to,offset,limit)
		if err != nil {
			logp.Info("Failed to get Network List from CV API %s", err.Error())
			return allComponents, err
		}

		err = json.Unmarshal(body, &components)
		if err != nil {
			logp.Info("Failed to Unmarshal data from API %s", err.Error())
			return allComponents, err
		}

		retrieved := len(components)
		logp.Info("Components data size: %d,limit:%d",retrieved,limit)
		offset = offset + retrieved
		allComponents = append(allComponents,components...)
		if retrieved < limit || retrieved <= 0{
			logp.Info("Break data retrieving,components data size: %d,limit:%d",retrieved,limit)
			break
		}else{
			logp.Info("Continue data retrieving,components data size: %d,limit:%d",retrieved,limit)
		}

	}
	return allComponents,nil
}


func (mc *CVClient) GetSensors(lastSyncUp int64,sync_up_to time.Time) ( CVSensors, error) {
	netURL := fmt.Sprintf("%s/api/1.0/sensor", mc.URL)
	var allSensors CVSensors
	var offset int =0
	limit :=mc.Limit
	for {
		var sensors CVSensors
		body, err := mc.getData(netURL,lastSyncUp,sync_up_to,offset,limit)
		if err != nil {
			logp.Info("Failed to get data from sensor API %s", err.Error())
			return sensors, err
		}

		var sensorMap map[string]CVSensor
		err = json.Unmarshal(body, &sensorMap)
		if err != nil {
			logp.Info("Failed to Unmarshal data from API %s", err.Error())
			return sensors, err
		}

		for _, sensor := range sensorMap{
			sensors = append(sensors,sensor)
		}

		retrieved := len(sensors)
		logp.Info("Sensors data size: %d,limit:%d",retrieved,limit)
		offset = offset + retrieved
		allSensors = append(allSensors,sensors...)
		if retrieved < limit || retrieved <= 0{
			logp.Info("Break data retrieving,sensors data size: %d,limit:%d",retrieved,limit)
			break
		}else{
			logp.Info("Continue data retrieving,sensors data size: %d,limit:%d",retrieved,limit)
		}

	}
	return allSensors,nil
}

func (mc *CVClient) GetEvents(lastSyncUp int64,sync_up_to time.Time) (CVEvents, error) {
	netURL := fmt.Sprintf("%s/api/1.0/event", mc.URL)
	var allEvents CVEvents
	var offset int =0
	limit :=mc.Limit
	for {
		var events CVEvents
		body, err := mc.getData(netURL,lastSyncUp,sync_up_to,offset,limit)
		if err != nil {
			logp.Info("Failed to get data from CV events API %s", err.Error())
			return allEvents, err
		}
		err = json.Unmarshal(body, &events)
		if err != nil {
			logp.Info("Failed to Unmarshal data from API %s", err.Error())
			return allEvents, err
		}

		retrieved := len(events)
		logp.Info("Event data size: %d,limit:%d",retrieved,limit)
		offset = offset + retrieved

		allEvents = append(allEvents,events...)
		if retrieved < limit || retrieved <= 0{
			logp.Info("Break data retrieving,Event data size: %d,limit:%d",retrieved,limit)
			break
		}else{
			logp.Info("Continue data retrieving,Event data size: %d,limit:%d",retrieved,limit)
		}
	}
	return allEvents,nil

}

func (mc *CVClient) GetFlows(lastSyncUp int64,sync_up_to time.Time) (CVFlows, error) {
	netURL := fmt.Sprintf("%s/api/1.0/flows", mc.URL)
	var allFlows CVFlows
	var offset int =0
	var retrieved int =0
	limit :=mc.Limit
	for{
		var flows CVFlows
		body, err := mc.getData(netURL,lastSyncUp,sync_up_to,offset,limit)
		if err != nil {
			logp.Info("Failed to get data from CV flow API %s", err.Error())
			return flows, err
		}
		err = json.Unmarshal(body, &flows)
		if err != nil {
			logp.Info("Failed to Unmarshal data from flow API %s", err.Error())
			return flows, err
		}
		retrieved = len(flows)
		logp.Info("Flow data size: %d,limit:%d",retrieved,limit)
		offset = offset + retrieved
		allFlows = append(allFlows,flows...)
		if retrieved < limit || retrieved==0{
			logp.Info("Break data retrieving,flow data size: %d,limit:%d",retrieved,limit)
			break
		}else{
			logp.Info("Continue data retrieving,flow data size: %d,limit:%d",retrieved,limit)
		}
	}

	return allFlows, nil
}

func (mc *CVClient) GetGroups(lastSyncUp int64,sync_up_to time.Time) (CVGroups, error) {
	netURL := fmt.Sprintf("%s/api/1.0/group", mc.URL)
	var allGroups CVGroups
	var offset int =0
	var retrieved int =0
	limit :=mc.Limit
	for {
		var groups CVGroups
		body, err := mc.getData(netURL,lastSyncUp, sync_up_to,offset, limit)
		if err != nil {
			logp.Info("Failed to get data from CV group API %s", err.Error())
			return groups, err
		}
		err = json.Unmarshal(body, &groups)
		if err != nil {
			logp.Info("Failed to Unmarshal data from group API %s", err.Error())
			return groups, err
		}
		logp.Info("data from group API %+v", groups)
		retrieved =len(groups)
		offset = offset + retrieved
		allGroups = append(allGroups,groups...)
		if retrieved < limit || retrieved==0{
			logp.Info("Break data retrieving,groups data size: %d,limit:%d",retrieved,limit)
			break
		}else{
			logp.Info("Continue data retrieving,groups data size: %d,limit:%d",retrieved,limit)
		}

	}
	return allGroups, nil
}

