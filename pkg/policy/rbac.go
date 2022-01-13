package policy

import (
	rbacv3 "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type RBACBuilder struct {
	Opts *RBACOpts
}

// RBACOpts is a wrapper configuration object that allows us to break down
// the RBAC policy object into smaller pieces that users can digest quickly.
type RBACOpts struct {
	Name     string
	Action   rbacv3.RBAC_Action
	Policies map[string]*rbacv3.Policy
}

type RBACOpt func(*RBACOpts)

// WithAction specifies policy creation per the required access control verb.
func WithAction(action rbacv3.RBAC_Action) RBACOpt {
	return func(opts *RBACOpts) {
		opts.Action = action
	}
}

// WithHeaderMatchers allows policy creation with user defined header matching rules.
func WithHeaderMatchers() RBACOpt {
	return func(opts *RBACOpts) {
	}
}

// NewRBACBuilder is a constructor for RBACBuilder that takes various
// args for a singular RBAC policy.
func NewRBACBuilder() *RBACBuilder {
	return &RBACBuilder{}
}

func (r *RBACBuilder) AddPolicy(name string, options ...RBACOpt) error {
	if r.Opts.Policies == nil {
		r.Opts.Policies = make(map[string]*rbacv3.Policy)
	}

	r.Opts.Policies[r.Opts.Name] = &rbacv3.Policy{}
	// Create our permissions
	perms := make([]*rbacv3.Permission, 0)
	perms = append(perms, &rbacv3.Permission{})

	r.Opts.Policies[r.Opts.Name].Permissions = perms
	// Create our policy principals
	principals := make([]*rbacv3.Principal, 0)
	principals = append(principals, &rbacv3.Principal{})

	r.Opts.Policies[r.Opts.Name].Principals = principals

	return nil
}

func (r *RBACBuilder) Build() ([]byte, error) {
	return proto.Marshal(&rbacv3.RBAC{
		Action:   r.Opts.Action,
		Policies: r.Opts.Policies,
	})
}

func (b *RBACBuilder) BuildJSON() ([]byte, error) {
	return protojson.Marshal(&rbacv3.RBAC{
		Action:   b.Opts.Action,
		Policies: b.Opts.Policies,
	})
}
