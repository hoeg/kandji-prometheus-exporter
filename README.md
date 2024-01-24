# kandji-prometheus-exporter
Prometheus exporter for Kandji.io. Get an overview of your fleet at a glance.

It exposes an endpoint for Prometheus to scrape. 
When the `/scrape` endpoint is called, the exporter calls Kandji to get a list of devices.
Metrics are calculated from the returned list.

The `/metrics` endpoint can be protected by an API key.
This is recommended since metrics might contain sensitive information about your Kandji usage.

## Exported metrics

`kandji_mac_os_version`

Number of machines on different MacOS versions.

`kandji_device_blueprint`

Number of devices on the different blueprints.

## Env vars

`KANDJI_PROM_EXPORTER_PORT`

The port to expose ser `/scrape` and `/metrics` endpoint on.

`KANDJI_PROM_EXPORTER_KANDJI_URL`

Your organizations Kandji URL. 
The URL can be found in the menu where you create the access token in the Kandji dashboard.

`KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE`

A local path to a file containing your Kandji API token.
The format of the file is to just contain the token an nothing else.

`KANDJI_PROM_EXPORTER_METRICS_API_KEY_FILE`

A local path to a file containing your secret API token for this exporter.
The format of the file is to just contain the token an nothing else.

## Requirements

Kandji API token must have permission to call the [List Device](https://api-docs.kandji.io/#78209960-31a7-4e3b-a2c0-95c7e65bb5f9) endpoint.

The token must be places in a file on the disk and pointed to by `KANDJI_PROM_EXPORTER_KANDJI_API_TOKEN_FILE`
