// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period        time.Duration `config:"period"`
	CVHost        string        `config:"cvhost"`
	CVKey         string        `config:"cvkey"`
	ScanComponent int 	    `config:"scancomponent"`
	ScanSensor    int 	    `config:"scansensor"`
	ScanFlow      int 	    `config:"scanflow"`
	ScanGroup    int 	    `config:"scangroup"`
	ScanEvent     int 	    `config:"scanevent"`
	ScanReference int 	    `config:"scanreference"`
	ScanVariable  int 	    `config:"scanvariable"`
	ScanTag       int 	    `config:"scantag"`
	ScanPropRule  int 	    `config:"scanpropertyrule"`
	ScanPortRule  int 	    `config:"scanportrule"`
	//
	//CVNetworkIDs  []string      `config:"cvnewtorkids"`
	//CVOrgID       string        `config:"cvorgid"`
	//NwConnStat        int           `config:"nwconnstat"`
	//NwLatencyStat     int           `config:"nwlatencystat"`
	//DeviceConnStat    int           `config:"devconnstat"`
	//DeviceLatencyStat int           `config:"devlatencystat"`
	//ClientConnStat    int           `config:"clconnstat"`
	//ClientLatencyStat int           `config:"cllatencystat"`
	//ScanSecret        string        `config:"scanSecret"`
	//ScanValidator     string        `config:"scanValidator"`
	//ScanEnable        int           `config:"scanEnable"`
	//VideoPeriod       time.Duration `config:"videoPeriod"`
	//CameraZoneList    []string      `config:"cameraZoneList"`
	//
	//CVNetworkIDsAll map[string]string
}

var DefaultConfig = Config{
	Period:     1 * time.Minute,
	CVHost: "https://localhost:443",
}
