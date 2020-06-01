# doubtnut

Pdf Service

## Installation

```
OS Environment - Linux (Ubuntu)
Make Sure Go and go mod is Installed in the system
Install dependencies
"go mod init [path]"
"go get ./..."

"go run main.go" to start the server
U
ses SendGrid to send the emails,Currently email are hardcoded but can be made dynamic based on userid
Add SendGrid Api key in .env
Generated pdf files in static/pdfs/ folder

```

Payload

```
curl -X POST \
  http://127.0.0.1:8080/api/v1/sendPdf/1212 \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
    "1": "As light from a star spreads out and weakens, do gaps form between the photons?",
    "2": "Can momentum be hidden to human eyes like how kinetic energy can be hidden as heat?",
    "3":"Can you make a shock wave of light by breaking the light barrier just like supersonic airplanes break the sound barrier?"
}'
```
