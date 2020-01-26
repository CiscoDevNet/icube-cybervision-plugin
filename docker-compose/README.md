# CyberVision system status, and system analysis demo

This demo will build a dashboard that will show the real-time CyberVision system data.
  
[Docker compose](https://docs.docker.com/compose/) start and configure the following service following:

- icube-cybervision-plugin : Beats plugin that polls CyberVision API and pushes data to elastic search
- elasticsearch: elasticsearch service. 
- kibana: Kibana service connected to elasticsearch

### Prequisites
Install [Docker](https://www.docker.com/) and [Docker compose](https://docs.docker.com/compose/).

## Instruction
Open a terminal window at the current folder.

### Config
#### 1. Update config in `config/icube-cybervision-plugin.yml`:
```
  period:   10s
  cvhost:   https://cybervision_server_name
  cvkey:    your_ics_token
  scancomponent:      1
  scansensor:         1
  scanflow:           1
  scangroup:          0
  scanevent:          1
  scanreference:      0
  scanvariable:       0
  scantag:            0
  scanpropertyrule:   0
  scanportrule:       0

```

#### 2. Update config file permission and ownership
```
sudo chmod go-w config/icube-cybervision-plugin.yml
sudo chown root:root config/icube-cybervision-plugin.yml

sudo chmod go-w config/fields.yml
sudo chown root:root config/fields.yml
```

### Start docker-compose
```
docker-compose up -d
```

You can use this command to see status or logs:
```
docker-compose ps
docker-compose logs -f
```

## Kibana settings
After around 1 min you should be able to access Kibana dashboard at [http://localhost:5601/](http://localhost:5601/).

After around 15 mins, elasticsearch should start getting data from CyberVision API.

#### 1. Create indexes
When you start getting data then add these two indexes to Kibana: `icube*.` and create time filter on "creation_time"

#### 2. Import predefined visualizations(charts) and dashboard config

In the `Management > Saved Objects`, view, click `Import` and import the `config/Kibana_icube_dashboard.json`. And config fields to the right index.

#### 3. Dashboard
Click `Dashboard` at the navigation panel and select `Dashboard`. You can change the **Time Range** at the right top to show data from different time window.







 
