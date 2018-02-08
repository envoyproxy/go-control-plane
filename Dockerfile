# Integration test docker file
FROM envoyproxy/envoy:latest
ADD sample /sample
ADD build/integration.sh build/integration.sh
ADD bin/test /bin/test
ENTRYPOINT ["build/integration.sh"]
