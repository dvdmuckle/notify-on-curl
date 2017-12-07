# notify-on-curl
Simple service to trigger desktop notification on curl. This is explicitly designed to work with the `motion` camera and motion detector as a manner to receive notifications of motion on some machine with a static IP. As such, this is really simple and not at all secure. Use at your own risk.
## Get it
```bash
go get -u github.com/dvdmuckle/notify-on-curl
```
You can set an environment variable `PORT` to set the port to run the service on. The service will print out a randomly generated key, which will be used when curling the service.
## Usage
Send a POST to the host where the service is running. For example, if running the service on localhost:
```bash
curl -XPOST localhost:8080/$KEY
```
Where `$KEY` is the random key received from first running.
