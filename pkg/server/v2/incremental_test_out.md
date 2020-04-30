# Incremental Unit Test Execution

To execute the test:

```bash
alecholmes@MacBook-Pro v2 % go test -timeout 30s -run TestDeltaResponseHandlers
```

```bash
2020/04/22 16:59:56 processDelta() for Type: type.googleapis.com/envoy.api.v2.ClusterLoadAssignment
2020/04/22 16:59:56 stream 1 open for type.googleapis.com/envoy.api.v2.ClusterLoadAssignment
2020/04/22 16:59:56 Starting send loop and watching for values on watches:
2020/04/22 16:59:56 Started deltaHandler() request go routine for resource Type: type.googleapis.com/envoy.api.v2.ClusterLoadAssignment
2020/04/22 16:59:56 Recieved delta request in process function
Creating a delta watch...
Found resources for type.googleapis.com/envoy.api.v2.ClusterLoadAssignment
Sent cache.DeltaResponse{DeltaRequest:envoy_api_v2.DeltaDiscoveryRequest{Node:(*envoy_api_v2_core.Node)(nil), TypeUrl:"", ResourceNamesSubscribe:[]string(nil), ResourceNamesUnsubscribe:[]string(nil), InitialResourceVersions:map[strin
g]string(nil), ResponseNonce:"", ErrorDetail:(*status.Status)(nil), XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}, Version:"1", ResourceMarshaled:false, Resources:[]types.Resource{(*envoy_api_v2.Cl
usterLoadAssignment)(0xc0002bc720)}, MarshaledResources:[][]uint8(nil)}
finished....
2020/04/22 16:59:56 OnStreamDeltaResponse() streamID: 1
Received stream response
&envoy_api_v2.DeltaDiscoveryResponse{SystemVersionInfo:"", Resources:[]*envoy_api_v2.Resource{(*envoy_api_v2.Resource)(0xc00011e540)}, TypeUrl:"type.googleapis.com/envoy.api.v2.ClusterLoadAssignment", RemovedResources:[]string(nil), 
Nonce:"1", XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}
```

```
2020/04/22 16:45:46 processDelta() for Type: type.googleapis.com/envoy.api.v2.Cluster
2020/04/22 16:45:46 stream 1 open for type.googleapis.com/envoy.api.v2.Cluster
2020/04/22 16:45:46 Starting send loop and watching for values on watches:
2020/04/22 16:45:46 Started deltaHandler() request go routine for resource Type: type.googleapis.com/envoy.api.v2.Cluster
2020/04/22 16:45:46 Recieved delta request in process function
Creating a delta watch...
Found resources for type.googleapis.com/envoy.api.v2.Cluster
Sent {DeltaRequest:{Node:<nil> TypeUrl: ResourceNamesSubscribe:[] ResourceNamesUnsubscribe:[] InitialResourceVersions:map[] ResponseNonce: ErrorDetail:<nil> XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} Version:2 Resou
rceMarshaled:false Resources:[name:"cluster0" type:EDS eds_cluster_config:<eds_config:<ads:<> resource_api_version:V2 > > connect_timeout:<seconds:5 > ] MarshaledResources:[]}
finished....
2020/04/22 16:45:46 Received values over deltaClusters chan
2020/04/22 16:45:46 OnStreamDeltaResponse() streamID: 1
resources:<name:"cluster0" resource:<type_url:"type.googleapis.com/envoy.api.v2.Cluster" value:"\n\010cluster0\032\006\n\0040\001\032\000\"\002\010\005\020\003" > > type_url:"type.googleapis.com/envoy.api.v2.Cluster" nonce:"1" 
```

```
2020/04/22 16:45:46 processDelta() for Type: type.googleapis.com/envoy.api.v2.Listener
2020/04/22 16:45:46 stream 1 open for type.googleapis.com/envoy.api.v2.Listener
2020/04/22 16:45:46 Starting send loop and watching for values on watches:
2020/04/22 16:45:46 Started deltaHandler() request go routine for resource Type: type.googleapis.com/envoy.api.v2.Listener
2020/04/22 16:45:46 Recieved delta request in process function
Creating a delta watch...
Found resources for type.googleapis.com/envoy.api.v2.RouteConfiguration
Sent {DeltaRequest:{Node:<nil> TypeUrl: ResourceNamesSubscribe:[] ResourceNamesUnsubscribe:[] InitialResourceVersions:map[] ResponseNonce: ErrorDetail:<nil> XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} Version:3 Resou
rceMarshaled:false Resources:[name:"route0" virtual_hosts:<name:"route0" domains:"*" routes:<match:<prefix:"/" > route:<cluster:"cluster0" > > > ] MarshaledResources:[]}
finished....
2020/04/22 16:45:46 OnStreamDeltaResponse() streamID: 1
resources:<name:"route0" resource:<type_url:"type.googleapis.com/envoy.api.v2.RouteConfiguration" value:"\n\006route0\022\036\n\006route0\022\001*\032\021\n\003\n\001/\022\n\n\010cluster0" > > type_url:"type.googleapis.com/envoy.api.
v2.RouteConfiguration" nonce:"1" 
```

```
2020/04/22 16:45:46 processDelta() for Type: type.googleapis.com/envoy.api.v2.Listener
2020/04/22 16:45:46 stream 1 open for type.googleapis.com/envoy.api.v2.Listener
2020/04/22 16:45:46 Starting send loop and watching for values on watches:
2020/04/22 16:45:46 Started deltaHandler() request go routine for resource Type: type.googleapis.com/envoy.api.v2.Listener
2020/04/22 16:45:46 Recieved delta request in process function
Creating a delta watch...
Found resources for type.googleapis.com/envoy.api.v2.Listener
Sent {DeltaRequest:{Node:<nil> TypeUrl: ResourceNamesSubscribe:[] ResourceNamesUnsubscribe:[] InitialResourceVersions:map[] ResponseNonce: ErrorDetail:<nil> XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} Version:4 Resou
rceMarshaled:false Resources:[name:"listener0" address:<socket_address:<address:"127.0.0.1" port_value:80 > > filter_chains:<filters:<name:"envoy.http_connection_manager" typed_config:<type_url:"type.googleapis.com/envoy.config.filte
r.network.http_connection_manager.v2.HttpConnectionManager" value:"\022\004http*\016\n\014envoy.routerj\200\001\n\032envoy.http_grpc_access_log\"b\nEtype.googleapis.com/envoy.config.accesslog.v2.HttpGrpcAccessLogConfig\022\031\n\027\
n\004echo\022\017\n\r\n\013xds_cluster\032\016\n\0040\001\032\000\022\006route0" > > > ] MarshaledResources:[]}
finished....
2020/04/22 16:45:46 OnStreamDeltaResponse() streamID: 1
resources:<name:"listener0" resource:<type_url:"type.googleapis.com/envoy.api.v2.Listener" value:"\n\tlistener0\022\017\n\r\022\t127.0.0.1\030P\032\263\002\032\260\002\n\035envoy.http_connection_manager\"\216\002\n`type.googleapis.co
m/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager\022\251\001\022\004http*\016\n\014envoy.routerj\200\001\n\032envoy.http_grpc_access_log\"b\nEtype.googleapis.com/envoy.config.accesslog.v2.HttpGrpcAccessL
ogConfig\022\031\n\027\n\004echo\022\017\n\r\n\013xds_cluster\032\016\n\0040\001\032\000\022\006route0" > > type_url:"type.googleapis.com/envoy.api.v2.Listener" nonce:"1" 
```

```
PASS
ok      github.com/envoyproxy/go-control-plane/pkg/server/v2    0.128s
```