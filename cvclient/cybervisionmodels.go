package cvclient

import (
	"github.com/elastic/beats/libbeat/common"
)

type CVSensor struct {
	Status string `json:"status"`
	Recording bool `json:"recording"`
	RecordingLastSize int `json:"recording_last_size"`
	Name string `json:"name"`
	RecordingLastStop string `json:"recording_last_stop"`
	AutoConfig bool `json:"auto_config"`
	FlowRateDay int `json:"flow_rate_day"`
	CreationTime string `json:"creation_time"`
	FlowsPercentage string `json:"flows_percentage"`
	FlowCountDay int `json:"flow_count_day"`
	Filter CVSensorFilter `json:"filter,omitempty"`
	Version string `json:"version"`
	RecordingLastStart string `json:"recording_last_start"`
	IP string `json:"ip"`
	SerialNumber string `json:"serial_number"`
	LastActiveTime string `json:"last_active_time"`
	ID string `json:"id"`
}

type CVSensorFilter struct {
	CustomInput string `json:"custom_input"`
	CaptureMode string `json:"capture_mode"`
}

type CVSensors []CVSensor

type CVLink struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Link string `json:"link"`
	Order int `json:"order"`
}

type CVulnerability struct {
	CreationTime string `json:"creation_time"`
	LastUpdate string `json:"last_update"`
	MatchingTime string `json:"matching_time"`
	ID string `json:"id"`
	CVE string `json:"cve"`
	VendorId string `json:"vendor_id"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	FullDescription string `json:"full_description"`
	Solution string `json:"solution"`
	CVSS float64 `json:"CVSS"`
	CVSSTemporal float64 `json:"CVSS_temporal"`
	CVSSVectorString string `json:"CVSS_vector_string"`
	Links []CVLink `json:"links,omitempty"`
}


type CVComponent struct {
	ID string `json:"id"`
	CreationTime string `json:"creation_time"`
	LastActiveTime string `json:"last_active_time"`
	IP string `json:"ip"`
	MAC string `json:"mac"`
	Name string `json:"name"`
	ModelName string `json:"model_name"`
	ModelRef string `json:"model_ref"`
	FWVersion string `json:"fw_version"`
	HWVersion string `json:"hw_version"`
	SerialNumber string `json:"serial_number"`
	VendorName string `json:"vendor_name"`
	ProjectName string `json:"project_name"`
	ProjectVersion string `json:"project_version"`
	Group string `json:"group,omitempty"` //group name will be used in getting Group API
	Tags    common.MapStr `json:"tags,omitempty"`
	Properties common.MapStr `json:"properties,omitempty"`
	Vulnerabilities []CVulnerability `json:"vulnerabilities,omitempty"`
}

type CVComponents []CVComponent

type CVFlow struct {
	Src CVFlowEndPoint `json:"src"`
	Tags common.MapStr `json:"tags"`
	Dst CVFlowEndPoint `json:"dst"`
	ID string `json:"id"`
	NetworkCategory string `json:"network_category"`
	FirstSeen string `json:"first_seen"`
	EtherType string `json:"ethertype"`
	Protocol string `json:"protocol"`
	Properties common.MapStr `json:"properties,omitempty"`
	LastSeen string `json:"last_seen"`
	SensorId string `json:"sensor_id"`
}

type CVFlows []CVFlow


type CVFlowEndPoint struct {
	IP string `json:"ip"`
	MAC string `json:"mac"`
	Component common.MapStr `json:"component,omitempty"`
	Port int `json:"port"`
}

type CVGroup struct {
	ID string `json:"id"`
	CreationTime string `json:"creation_time,omitempty"`
	Label string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Comments string `json:"comments,omitempty"`
	Color string `json:"color,omitempty"`
	IndustrialImpact string `json:"industrial_impact,omitempty"`
	Properties []common.MapStr `json:"properties,omitempty"`
	Components CVComponents `json:"components,omitempty"`
}

type CVGroups []CVGroup

type CVEvent struct {
	ID string `json:"id"`
	Category string `json:"category"`
	Severity string `json:"severity"`
	Family string `json:"family"`
	Component common.MapStr `json:"component"`
	CreationTime string `json:"creation_time"`
	Message string `json:"message"`
	Type string `json:"type"`
	ShortMessage string `json:"short_message"`
}

type CVEvents []CVEvent

func (cp *CVComponent) GetMapStr(addlnKVP map[string]string) ([]common.MapStr, error){
	var mapStrArr []common.MapStr
	//mapStr := common.MapStr{
	//	"type":                stattype,
	//	"datatype":            sd.Type,
	//}
	//
	return mapStrArr,nil
}
