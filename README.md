# kandji-prometheus-exporter
Prometheus exporter for Kandji.io. Get an overview of your fleet at a glance.

It exposes an endpoint for Prometheus to scrape. 
When the `/scrape` endpoint is called, the exporter calls Kandji to get a list of devices.
Metrics are calculated from the returned list.

## Exported metrics

`kandji_mac_os_version`

`kandji_device_blueprint`

## Env vars

`KANDJI_PROM_EXPORTER_PORT`

The port to expose ser `/scrape` and `/metrics` endpoint on.

`KANDJI_PROM_EXPORTER_KANDJI_URL`

Your organizations Kandji URL. 
The URL can be found in the menu where you create the access token in the Kandji dashboard.

`KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE`

A local path to a file containing your API token.
The format of the file is to just contain the token an nothing else.

## Requirements

Kandji API token must have permission to call the [List Device](https://api-docs.kandji.io/#78209960-31a7-4e3b-a2c0-95c7e65bb5f9) endpoint.

The token must be places in a file on the disk and pointed to by `KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE`
