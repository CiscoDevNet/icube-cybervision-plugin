package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/CiscoDevNet/icube-cybervision-plugin/config"
	_ "github.com/CiscoDevNet/icube-cybervision-plugin/cvclient"
)

type iCube_cybervision_plugin struct {
	done   chan struct{}
	config config.Config
	settings *config.LocalSetting
	client beat.Client
	b      *beat.Beat
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	settings := &config.LocalSetting{
		SettingFile: "/tmp/settings.json",
	}

	settings.Load("/tmp/settings.json")

	bt := &iCube_cybervision_plugin{
		done:   make(chan struct{}),
		config: c,
		b:      b,
		settings: settings,
	}

	return bt, nil
}

func (bt *iCube_cybervision_plugin) Run(b *beat.Beat) error {
	logp.Info("icube-cybervision-plugin is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	// health api poller
	ticker := time.NewTicker(bt.config.Period)
	poller := NewCVPoller(bt, bt.config,bt.settings)

	fmt.Printf("Period CyberVision %+v \r\n", bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
			poller.Run()
		}
	}
}

func (bt *iCube_cybervision_plugin) Stop() {
	bt.client.Close()
	close(bt.done)
}
