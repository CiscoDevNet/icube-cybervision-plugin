package beater

import (
	"time"
	"fmt"
	"encoding/json"
	"github.com/CiscoDevNet/icube-cybervision-plugin/config"
	"github.com/CiscoDevNet/icube-cybervision-plugin/cvclient"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

type CybervisionPoller struct {
	CVPoller
}

func NewCVPoller(cv_plugin *iCube_cybervision_plugin, config config.Config,settings *config.LocalSetting) *CybervisionPoller {
	mc := cvclient.NewCVClient(config.CVHost, config.CVKey, config.Period)
	//mc.LocalSettings = settings
	//mc.updateSettings(cv_plugin.settings)
	//fmt.Printf("Period health %+v",cv_plugin.settings)
	poller := &CybervisionPoller{}
	poller.cv_plugin = cv_plugin
	poller.config = config
	poller.mc = mc
	poller.settings = cv_plugin.settings
	return poller
}

// This is function that will call CVClient to fetch & publish data based on
// config item.  CVClient should have no understanding of beats framework except
func (p *CybervisionPoller) ConvertTimeStr(time_string string) time.Time {
	time_out,err := time.Parse(time.RFC3339, time_string)
	if err ==nil{
		return time.Now()
	}else {
		return time_out
	}

}

func (p *CybervisionPoller) GetCurrentTime() time.Time {
	return time.Now()
}

func (cp *CybervisionPoller) UpdateModuleName(valuesMap common.MapStr,moduleName string) (common.MapStr){
	valuesMap["CVModuleName"] = moduleName
	return valuesMap
}

// function that returns mapstr type.
func (p *CybervisionPoller) Run() {
	fmt.Printf("Excuting Plugins poller!\r\n")
	var optOff int =0
	var sync_up_end_to time.Time = p.GetCurrentTime()
	logp.Info("%+v", p.config)
	// Publish Network Connection Event
	logp.Info("Getting ScanComponent stat for network %+v", p.config.ScanComponent)
	//Publish Components for CV
	if p.config.ScanComponent != optOff {
		logp.Info("ScanComponent !=0, stat for network %+v", p.config.ScanComponent)
		components, err := p.mc.GetComponents(p.settings.LastSyncUp, sync_up_end_to)
		if err == nil {
			for j, component := range components {
				var jsonMap common.MapStr
				jsonEnc, _ := json.Marshal(component)
				json.Unmarshal(jsonEnc, &jsonMap)
				jsonMap = p.UpdateModuleName(jsonMap,"CVComponent")
				jsonMap["timestamp_es"]=p.ConvertTimeStr(component.CreationTime).Unix()
				event := beat.Event{
					Timestamp: p.ConvertTimeStr(component.CreationTime),
					Fields:    jsonMap,
				}
				event.SetID(component.ID)
				p.cv_plugin.client.Publish(event)
				logp.Info("Component data sent %d", j)
				logp.Info("Component data event sent,data:%+v", jsonMap)
			}
			logp.Info("Components data event sent")
		}else{
			logp.Info("No Components data event sent:nil")
		}
	}

	//Publish Sensors for CV
	if p.config.ScanSensor != optOff {
		logp.Info("ScanSensor !=0, stat for network %+v", p.config.ScanSensor)
		sensors, err := p.mc.GetSensors(p.settings.LastSyncUp, sync_up_end_to)
		if err == nil {
			for j, sensor := range sensors {
				var jsonMap common.MapStr
				jsonEnc, _ := json.Marshal(sensor)
				json.Unmarshal(jsonEnc, &jsonMap)
				jsonMap = p.UpdateModuleName(jsonMap,"CVSensor")
				jsonMap["timestamp_es"]=p.ConvertTimeStr(sensor.CreationTime).Unix()
				event := beat.Event{
					Timestamp: p.ConvertTimeStr(sensor.CreationTime),
					Fields:    jsonMap,
				}
				event.SetID(sensor.ID)
				p.cv_plugin.client.Publish(event)
				logp.Info("Sensor data event sent %d", j)
				logp.Info("Sensor data event sent,data:%+v", jsonMap)
			}
			logp.Info("Sensors data event sent")

		}else{
			logp.Info("No Sensors data event sent:nil")
		}

	}

	//Publish Flows for CV
	if p.config.ScanFlow != optOff {
		logp.Info("ScanFlow !=0, stat for network %+v", p.config.ScanFlow)
		flows, err := p.mc.GetFlows(p.settings.LastSyncUp, sync_up_end_to)
		if err == nil {
			for j, flow := range flows {
				var jsonMap common.MapStr
				jsonEnc, _ := json.Marshal(flow)
				json.Unmarshal(jsonEnc, &jsonMap)
				jsonMap = p.UpdateModuleName(jsonMap,"CVFlow")
				jsonMap["creation_time"]=flow.FirstSeen
				jsonMap["timestamp_es"]=p.ConvertTimeStr(flow.FirstSeen).Unix()
				event := beat.Event{
					Timestamp: p.ConvertTimeStr(flow.FirstSeen),
					Fields:    jsonMap,
				}
				event.SetID(flow.ID)
				p.cv_plugin.client.Publish(event)
				logp.Info("Flows event sent %d", j)
				logp.Info("Flows data event sent,data:%+v", jsonMap)
			}
			logp.Info("Flows data event sent")

		}else{
			logp.Info("No Flows data event sent:nil")
		}
	}

	//Publish Groups for CV
	if p.config.ScanGroup != optOff {
		logp.Info("ScanGroup !=0, stat for network %+v", p.config.ScanGroup)
		groups, err := p.mc.GetGroups(p.settings.LastSyncUp, sync_up_end_to)
		if err == nil {
			for j, group := range groups {
				var jsonMap common.MapStr
				jsonEnc, _ := json.Marshal(group)
				json.Unmarshal(jsonEnc, &jsonMap)
				jsonMap = p.UpdateModuleName(jsonMap,"CVGroup")
				jsonMap["timestamp_es"]=p.ConvertTimeStr(group.CreationTime).Unix()
				event := beat.Event{
					Timestamp: p.ConvertTimeStr(group.CreationTime),
					Fields:    jsonMap,
				}
				event.SetID(group.ID)
				p.cv_plugin.client.Publish(event)
				logp.Info("Group data event sent %d", j)
				logp.Info("Group data event sent,data:%+v", jsonMap)
			}
			logp.Info("Groups data event sent")

		}else {
			logp.Info("No Groups data event sent:nil")
		}
	}

	//Publish Events for CV
	if p.config.ScanEvent != optOff {
		logp.Info("ScanEvent !=0, stat for network %+v", p.config.ScanEvent)
		events, err := p.mc.GetEvents(p.settings.LastSyncUp, sync_up_end_to)
		if err == nil {
			for j, event_x := range events {
				var jsonMap common.MapStr
				jsonEnc, _ := json.Marshal(event_x)
				json.Unmarshal(jsonEnc, &jsonMap)
				jsonMap = p.UpdateModuleName(jsonMap,"CVEvent")
				jsonMap["timestamp_es"]=p.ConvertTimeStr(event_x.CreationTime).Unix()
				event := beat.Event{
					Timestamp: p.ConvertTimeStr(event_x.CreationTime),
					Fields:    jsonMap,
				}
				event.SetID(event_x.ID)
				p.cv_plugin.client.Publish(event)
				logp.Info("Event data event sent %d", j)
				logp.Info("Event data event sent,data:%+v", jsonMap)
			}
			logp.Info("Events data event sent, done!")

		}else{
			logp.Info("No Events data event sent:nil")
		}
	}

	////Publish references for CV
	//if p.config.ScanReference != 0 {
	//	references, err := p.mc.GetReferences(p.settings.LastSyncUp)
	//	if err == nil {

	//		for j, reference := range references {
	//			var jsonMap common.MapStr
	//			jsonEnc, _ := json.Marshal(event_x)
	//			json.Unmarshal(jsonEnc, &jsonMap)
	//			jsonMap = p.UpdateModuleName(jsonMap,"CVReference")
	//			event := beat.Event{
	//				Timestamp: time.Now(),
	//				Fields:    reference,
	//			}
	//			p.cv_plugin.client.Publish(event)
	//			logp.Info("Reference data event sent %d", j)
	//		}
	//		logp.Info("References data event sent")
	//
	//	}
	//}

	////Publish variables for CV
	//if p.config.ScanVariable != 0 {
	//	variables, err := p.mc.GetVariables(p.settings.LastSyncUp)
	//	if err == nil {
	//		for j, variable := range variables {
	//			event := beat.Event{
	//				Timestamp: time.Now(),
	//				Fields:    variable,
	//			}
	//			p.cv_plugin.client.Publish(event)
	//			logp.Info("Variable data event sent %d", j)
	//		}
	//		logp.Info("Variables data event sent")
	//
	//	}
	//}

	////Publish Tags for CV
	//if p.config.ScanTag != 0 {
	//	tags, err := p.mc.GetTags(p.settings.LastSyncUp)
	//	if err == nil {
	//		for j, tag := range tags {
	//			event := beat.Event{
	//				Timestamp: time.Now(),
	//				Fields:    tag,
	//			}
	//			p.cv_plugin.client.Publish(event)
	//			logp.Info("Tag data event sent %d", j)
	//		}
	//		logp.Info("Tags data event sent")
	//
	//	}
	//}
	//
	////Publish Tags for CV
	//if p.config.ScanPropRule != 0 {
	//	proprules, err := p.mc.GetPropRules(p.settings.LastSyncUp)
	//	if err == nil {
	//		for j, proprule := range proprules {
	//			event := beat.Event{
	//				Timestamp: time.Now(),
	//				Fields:    proprule,
	//			}
	//			p.cv_plugin.client.Publish(event)
	//			logp.Info("Proprule data event sent %d", j)
	//		}
	//		logp.Info("Proprules data event sent")
	//
	//	}
	//}
	//
	////Publish Tags for CV
	//if p.config.ScanPortRule != 0 {
	//	portrules, err := p.mc.GetPortRules()
	//	if err == nil {
	//		for j, portrule := range portrules {
	//			event := beat.Event{
	//				Timestamp: time.Now(),
	//				Fields:    portrule,
	//			}
	//			p.cv_plugin.client.Publish(event)
	//			logp.Info("Portrule data event sent %d", j)
	//		}
	//		logp.Info("Portrules data event sent")
	//
	//	}
	//}
	p.settings.Update(sync_up_end_to)
}
