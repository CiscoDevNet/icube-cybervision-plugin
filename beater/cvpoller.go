package beater

import (
	"time"
	"github.com/CiscoDevNet/icube-cybervision-plugin/config"
	"github.com/CiscoDevNet/icube-cybervision-plugin/cvclient"
	//"github.com/elastic/beats/libbeat/common"
)

type CVPoller struct {
	cv_plugin * iCube_cybervision_plugin
	config     config.Config
	timeout    time.Duration
	mc         cvclient.CVClient
	settings  * config.LocalSetting
}

//type CVPolleriIntf interface {
//	Run()
//}
