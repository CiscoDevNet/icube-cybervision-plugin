package cvclient

import (
	"io/ioutil"
	"crypto/tls"
	"net/http"
	"time"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/CiscoDevNet/icube-cybervision-plugin/config"
)

type CVClient struct {
	URL         string
	Key         string
	Period      time.Duration
	LocalSettings *config.LocalSetting
	Limit   int  `default:2000`
}

func NewCVClient(url string, key string, period time.Duration) CVClient {
	return CVClient{
		URL:         url,
		Key:         key,
		Period:      period,
		Limit:	     500,
	}
}

func (cvc *CVClient) updateSettings(settings *config.LocalSetting) bool {
	cvc.LocalSettings = settings
	return true
}

// THis is health API specific getData, not all API set are same to some
// custom t0-t1 handling
func (cvc *CVClient) getDataQueryParam(netURL string, params map[string]string) ([]byte, error) {

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	req, err := http.NewRequest("GET", netURL, nil)
	if err != nil {
		logp.Info("Failed to connect CV API %s", err.Error())
		return nil, err
	}
	//req.Header.Add("X-Cisco-CV-API-Key", mc.Key)

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	q.Add("token",cvc.Key)
	req.URL.RawQuery = q.Encode()

	logp.Info("Calling API URL %+v, %v", req.URL, req.URL.RawQuery)
	resp, err := client.Do(req)
	if err != nil {
		logp.Info("Failed to connect CV API %s", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	logp.Info("%s", string(body[:]))
	return body, err
}
