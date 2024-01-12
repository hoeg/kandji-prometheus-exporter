# kandji-prometheus-exporter
Prometheus exporter for Kandji.io. Get an overview of your fleet at a glance.

It exposes an endpoint for Prometheus to scrape. 
When the `/scrape` endpoint is called, the exporter makes an HTTPS call to kandji to get a list of devices.
The different versions are counted and written to prometheus such that we can get an overview of how many
devices are on different MacOS versions.

## Env vars

`KANDJI_PROM_EXPORTER_PORT`

`KANDJI_PROM_EXPORTER_KANDJI_URL`

`KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE`

## Requirements

Kandji API token must have permission to call the [List Device](https://api-docs.kandji.io/#78209960-31a7-4e3b-a2c0-95c7e65bb5f9) endpoint.

The token must be places in a file on the disk and pointed to by `KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE`
