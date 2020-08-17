package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	utilnet "k8s.io/apimachinery/pkg/util/net"
	heptiotoken "sigs.k8s.io/aws-iam-authenticator/pkg/token"
)

// TokenKind defines the token type used for remote authentication to your cluster(s).
// An example NullToken and the implementation for HeptioToken (AWS) are provided.
type TokenKind int

const (
	HeptioTokenKind TokenKind = iota
	NullTokenKind
)

// NullTokenProvider always returns a nil token
type NullTokenProvider struct{}

func (t *NullTokenProvider) Kind() TokenKind {
	return NullTokenKind
}

// AuthToken contains an auth token and its expiry time
type AuthToken struct {
	Token   string
	expires time.Time
}

// Expired return whether or not the token has past its expiry time
func (t *AuthToken) Expired() bool {
	return time.Now().After(t.expires)
}

type Provider interface {
	// Token provides a time limited token
	Token() (*AuthToken, error)
	Kind() TokenKind
}

// Generator interface defines the ability to grab token providers based on the cluster name
type Generator interface {
	ProviderForCluster(cluster string) (Provider, error)
}

// GeneratorProvider is a token provider which mints tokens as well
type GeneratorProvider interface {
	Provider
	Generator
}

type heptioTokenProvider struct {
	lastToken   *AuthToken
	clusterName string
	lock        sync.Mutex
}

func (t *heptioTokenProvider) refresh() error {
	t.lock.Lock()
	defer t.lock.Unlock()

	gen, err := heptiotoken.NewGenerator(false, false)
	if err != nil {
		fmt.Printf("heptio authenticator error %v", err)
		return err
	}
	rawToken, err := gen.Get(t.clusterName)

	if err != nil {
		fmt.Printf("token data invalid %v", err)
		return err
	}

	token := AuthToken{
		Token:   rawToken.Token,
		expires: rawToken.Expiration,
	}

	t.lastToken = &token
	return nil
}

// Token returns a heptio token, refreshing it if need be
func (t *heptioTokenProvider) Token() (*AuthToken, error) {
	if t.lastToken != nil {
		if time.Now().After(t.lastToken.expires) {
			err := t.refresh()
			if err != nil {
				return nil, err
			}
		}
	} else {
		err := t.refresh()
		if err != nil {
			return nil, err
		}
	}

	return t.lastToken, nil
}

func (t *heptioTokenProvider) Kind() TokenKind {
	return HeptioTokenKind
}

// NewHeptioProvider returns an initialized heptio token provider
func NewHeptioProvider(clusterName string) Provider {
	return &heptioTokenProvider{
		clusterName: clusterName,
	}
}

type heptioGenerator struct {
	clusters sync.Map
}

func (h *heptioGenerator) ProviderForCluster(cluster string) (Provider, error) {
	provider := NewHeptioProvider(cluster)
	providerFromMap, _ := h.clusters.LoadOrStore(cluster, provider)
	return providerFromMap.(*heptioTokenProvider), nil
}

func NewHeptioGenetor() Generator {
	return &heptioGenerator{}
}

type requestCanceler interface {
	CancelRequest(*http.Request)
}

type TokenRoundtripper struct {
	RoundTripper  http.RoundTripper
	TokenProvider Provider
}

func (rt *TokenRoundtripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(req.Header.Get("Authorization")) != 0 {
		return rt.RoundTripper.RoundTrip(req)
	}

	req = utilnet.CloneRequest(req)
	token, err := rt.TokenProvider.Token()
	if err != nil {
		return nil, err

	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	return rt.RoundTripper.RoundTrip(req)
}

func (rt *TokenRoundtripper) CancelRequest(req *http.Request) {
	if canceler, ok := rt.RoundTripper.(requestCanceler); ok {
		canceler.CancelRequest(req)
	}
}
