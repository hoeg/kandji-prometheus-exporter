# kandji-prometheus-exporter
Prometheus exporter for Kandji.io. Get an overview of your fleet at a glance.

It exposes an endpoint for Prometheus to scrape. 
When the `/scrape` endpoint is called, the exporter makes an HTTPS call to kandji with a JSON payload and an API token in the header.

## Usage

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd prometheus-exporter
