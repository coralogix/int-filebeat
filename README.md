# int-filebeat
This repository is plugin for filebeat.
The purpose of this plugin is send logs directly to Coralogix and validate them before.

Installation:
1. create code's repository file
```
go mod init filebeat_output_http.go
```
3. get all relevant packages:
```
go get github.com/elastic/beats/libbeat/beat
go get github.com/elastic/beats/libbeat/common
go get github.com/elastic/beats/libbeat/outputs
```
