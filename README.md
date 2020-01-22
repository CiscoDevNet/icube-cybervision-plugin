[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/CiscoDevNet/icube-cybervision-plugin)
# icube-cybervision-plugin

This is elastic beats (https://www.elastic.co/products/beats) plugin for CyberVision API. icube-cybervision-plugin enables data collection and analytics pipeline for multiple CyberVision API.
- CyberVision Sensor API
- CyberVision Component API
- CyberVision Event API
- CyberVision Flow API

This pipeline enables analyzing CybveVision Components, Sensors, Events, and Flow Data to provide more detail analytics data for ICS(Industrial Control System).

#### CyberVision Sensor API
Route:
GET /api/1.0/sensor

Example:
GET /api/1.0/sensor?token=YOUR_TOKEN&status=active&status=inactive

Retrieve whole list of sensors(active and inactive) used by Cyber Vision. Some of the key sensor parameters that

CyberVision Sensor API exposes are
	- status				: Sensor's status, active or inactive
	- recording				: Recording status
	- recording_last_size	: last size of recording
	- name					: friendly name of sensor
	- recording_last_stop	: last stop time of recording
	- auto_config			: sensor is auto configured or not
	- flow_rate_day			: flow rate of day
	- creation_time			: creation time of sensor
	- flows_percentage		: flows percentage info
	- flow_count_day		: flow count of day
	- filter				: filter options
	- version				: version information
	- recording_last_start  : last start of recording
	- ip					: IP of sensor
	- serial_number			: Serial number of sensor
	- last_active_time		: Last active time of sensor
	- id					: ID of sensor

#### CyberVision Component API

Route:
GET /api/1.0/component
Example:
Retrieve up to 5 live components tagged has PLC with vendor "Siemens AG - Industrial Automation - EWA" and with has property "s7plus-modulever" equals to "V4.0" and has property "s7plus- moduleref":
GET /api/1.0/component?token=YOUR_TOKEN&vendor=Siemens+AG+-+Industrial+Automation+-+EWA&limit=5&generation=0&property=s7plus-modulever%3DV4.0&property=s7plus-moduleref&tag=PLC


A Component is a node of the network having had communications. The API allows to retrieve the whole list of Components having been created by the SCV ordered by their last active time.

CyberVision Component API exposes are
	- id					: ID of component
	- creation_time			: Creation time of component
	- last_active_time		: Last active time of component
	- ip					: IP address
	- mac					: MAC address
	- name					: Name of Component
	- model_name			: Model name
	- model_ref				: Model reference
	- fw_version			: Firmware version
	- hw_version			: Hardware version
	- serial_number			: Serial number
	- vendor_name			: Vendor name
	- project_name			: Project name
	- project_version		: Project version
	- group					: Group name-group name will be used in getting Group API
	- tags					: Tags info
	- properties			: Properties
	- vulnerabilities		: Vulnerability list

Vulnerability that discovered on components

	- creation_time			: Creation time
	- last_update			: Last update
	- matching_time			: Matching time
	- id					: ID
	- cve					: CVE
	- vendor_id				: Vendor ID
	- title					: Title
	- summary				: Summary
	- full_description		: Full Description
	- solution				: Solution
	- CVSS					: CVSS
	- CVSS_temporal			: CVSS temporal
	- CVSS_vector_string	: CVSS vector string
	- links					: Relevant Links about Vulnerability

Links about Vulnerability
	- id					: Link ID
	- title					: Title
	- link					: Link Url
	- order					: Order



#### CyberVision Flow API

Route:
GET /api/1.0/flows/?token=YOUR_TOKEN

Example:
Retrieve flow between 2016-11-03 and 2017-11-14 12:05 with tag "Stop CPU" and "Write Var" with industrial as network category and has property "s7plus-moduleref" and with property "s7plus- modulever" should be equal to "V1.6" and with property "s7plus-io-ver" should be equal to "V2.0.2":

GET /api/1.0/flows/?token=YOUR_TOKEN&start=2016-11-03+00%3A00&tag=Stop+CPU&tag=Write+Var&limit=2&offset=0&generation=0&end=2017-11-14+12%3A05&property=s7plus-io-ver%3DV2.0.2&property=s7plus-modulever%3DV1.6&property=s7plus-moduleref&order=asc&netcat=field&netcat=control

CyberVision Flow API exposes are

	- id				: ID
    - sensor_id			: Sensor ID
    - src 				: Flow source, refer to FlowEndPoint
	- dst				: Flow destination,refer to FlowEndPoint
    - network_category	: Network category, such as control, field...
	- ethertype			: EtherType,such as IPv4,IPv6
    - protocol			: Protocol, such as TCP, UDP
    - first_seen		: First seen time
	- properties		: Properties about the flow
	- last_seen			: Last seen time
	- tags				: Tags info about the flow

FlowEndPoint:
	- ip 				: IP address
	- port 				: Port
	- mac				: MAC address
	- component			: Component Info, ID

#### CyberVision Event API

Route:

GET /api/1.0/event?token=YOUR_TOKEN

Example:
Retrieve the 5 events between the 2016-11-03 00:00 and the 2017-11-13 17:23, with "high" or "very- high" as severity and in "Identification" as category related to a component which have "d4:ae:52:aa:dc:93" as MAC address and "10.4.0.46" as IP address sorted by most recent first:
GET /api/1.0/event?token=YOUR_TOKEN&category=Identification&start=2016-11-03+00%3A00&mac=d4%3Aae%3A52%3Aaa%3Adc%3A9

CyberVision Event API exposes are
	- id				: ID
	- category			: Category
	- severity			: Severity, such as High,Medium,Low
	- family			: family, such as "Cisco Cyber vision",Component,Property,PLC Control...
	- component			: Component info, such as component id
	- creation_time		: Creation time
	- message			: Message description about the event
	- type				: Type, such as Classification,Behavioral...
	- short_message		: Short message about the event


This beats plugin polls the CyberVision APIs for these CybverVision stats and allows sending Stats
to elasticsearch or any of ouput service supported by
beats. (https://www.elastic.co/guide/en/logstash/current/output-plugins.html).

## icube-cybervision-plugin Data pipeline

![icube-cybervision-plugin pipeline](./docs/media/icube-cybervision-plugin-pipeline.png)

### sample dashboard
Sample kibana dashboards are included as part of docker compose is

#### icube-cybervision-plugin API: sample dashboard

![icube-cybervision-plugin API: sample dashboard](./docs/media/icube-cybervision-plugin-dashboard.png)

# Configuring icube-cybervision-plugin
Supports following plugin specific configs in icube-cybervision-plugin.yml file
-  period: Polling interval , recommended value 10s to 600s
-  cvhost: URL for cv API endpoint in format, https://192.168.1.10
-  cvkey:  CyberVision API token

Following are fields to control sections that will be collected by data collector
'1' represent data will be collected and '0' means data will be skipped(scangroup,scanreference,scanvariable,scantag,scanpropertyrule, and scanportrule have not implemented in current version yet).
- scancomponent:      1		Scan CyberVision Component API
- scansensor:         1		Scan CyberVision Sensor API
- scanflow:           1		Scan CyberVision Flow API
- scangroup:          0
- scanevent:          1		Scan CyberVision Event API
- scanreference:      0
- scanvariable:       0
- scantag:            0
- scanpropertyrule:   0
- scanportrule:       0


All these field are configured in icube-cybervision-plugin.yml config file


## Running icube-cybervision-plugin
### As docker-composer (**Recommended**)
https://github.com/CiscoDevNet/icube-cybervision-plugin/blob/master/docker-compose/README.md

### As binary
```
icube-cybervision-plugin -e -d *
```

## Building Source code
https://github.com/CiscoDevNet/icube-cybervision-plugin/blob/master/DEVELOPMENT.md
