# HTTP Upstream Test Server

This server is used by the integration tests as an upstream server.
It simply responds to HTTP GET requests with a canned message that is
then proxied to the test framework via Envoy.

For more info see the ```build/integration.sh``` script which runs the
upstream server, and ```pkg/test/main/main.go``` for the test
framework that sends and checks the requests.
