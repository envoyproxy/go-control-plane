module github.com/envoyproxy/go-control-plane/examples/dyplomat

go 1.16

require (
	github.com/envoyproxy/go-control-plane v0.9.6
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/grpc v1.30.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v11.0.0+incompatible
	sigs.k8s.io/aws-iam-authenticator v0.5.1
)

replace (
	k8s.io/api => k8s.io/api v0.18.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
)
