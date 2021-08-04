package xds

// Options are configuration settings for the discovery object
type Options struct {
	NodeID        string
	Zone          string
	Cluster       string
	ResourceNames []string // List of Envoy resource names to subscribe to
	ResourceType  string   // ex: type.googleapis.com/envoy.api.v2.ClusterLoadAssignment
}

// Option follows the functional opts pattern
type Option func(*Options)

// WithNode will inject the node id into the configuration object
func WithNode(id string) Option {
	return func(o *Options) {
		o.NodeID = id
	}
}

// WithZone will specificy which zone to use in the xDS discovery request
func WithZone(zone string) Option {
	return func(o *Options) {
		o.Zone = zone
	}
}

// WithCluster will specificy which cluster the request is announcing as
func WithCluster(cluster string) Option {
	return func(o *Options) {
		o.Cluster = cluster
	}
}

// WithResourceNames will inject a list of resources the user wants to place watches on
func WithResourceNames(names []string) Option {
	return func(o *Options) {
		o.ResourceNames = names
	}
}

// WithResourceType will inject the specific resource type that a user wants to stream
func WithResourceType(resource string) Option {
	return func(o *Options) {
		o.ResourceType = resource
	}
}
