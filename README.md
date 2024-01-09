# kandji-prometheus-exporter
Prometheus exporter for Kandji.io. Get an overview of your fleet at a glance.

It exposes an endpoint for Prometheus to scrape. 
When the `/scrape` endpoint is called, the exporter makes an HTTPS call to kandji to get a list of devices.
The different versions are counted and written to prometheus such that we can get an overview of how many
devices are on different MacOS versions.

## Usage

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd prometheus-exporter
