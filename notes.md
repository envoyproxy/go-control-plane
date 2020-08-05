# NOTES

All unit tests pass for the delta snapshot cache. They also pass for the delta server code.

For some reason the management server will respond to envoys initial request of the runtime payload, then die. Need to dig into that



Running:

./test --xds=delta  -debug=true 2> xds.delta.log

envoy -c ../sample/bootstrap-delta.yaml --drain-time-s 1 -l debug 2> envoy.delta.log


