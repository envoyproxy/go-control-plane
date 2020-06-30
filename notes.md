NOTES

Snapshot cache delta unit test is failing intermittenly because the Cluster and Endpoint watches are failing to return a snapshot
But the other resources types always pass regardless of the test case

As of right now, I am noticing that the management server is only creating new watches for a singular resource type when it receives a new snapshot



Running:

./test --xds=delta  -debug=true 2> xds.delta.log
envoy -c ../sample/bootstrap-delta.yaml --drain-time-s 1 -l debug 2> envoy.delta.log


