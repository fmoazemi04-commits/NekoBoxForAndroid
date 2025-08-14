package option

import (
	"context"
	"encoding/json"
)

// HeadlessRule represents a rule without full configuration
type HeadlessRule struct {
	Type           string
	DefaultOptions DefaultHeadlessRule
}

// DefaultHeadlessRule contains the rule matching options
type DefaultHeadlessRule struct {
	Domain        []string `json:"domain,omitempty"`
	DomainSuffix  []string `json:"domain_suffix,omitempty"`
	DomainKeyword []string `json:"domain_keyword,omitempty"`
	DomainRegex   []string `json:"domain_regex,omitempty"`
	IPCIDR        []string `json:"ip_cidr,omitempty"`
}

// TunPlatformOptions represents TUN platform-specific options
type TunPlatformOptions struct {
	HTTPProxy  *HTTPProxyOptions  `json:"http_proxy,omitempty"`
	DNSHijack  []string          `json:"dns_hijack,omitempty"`
}

// HTTPProxyOptions represents HTTP proxy options
type HTTPProxyOptions struct {
	Enabled bool   `json:"enabled,omitempty"`
	Server  string `json:"server,omitempty"`
	Port    int    `json:"port,omitempty"`
}

// Options represents the main configuration options
type Options struct {
	Log             *LogOptions              `json:"log,omitempty"`
	DNS             *DNSOptions             `json:"dns,omitempty"`
	Inbounds        []Inbound               `json:"inbounds,omitempty"`
	Outbounds       []Outbound              `json:"outbounds,omitempty"`
	Route           *RouteOptions           `json:"route,omitempty"`
	Experimental    *ExperimentalOptions    `json:"experimental,omitempty"`
}

// UnmarshalJSONContext unmarshals JSON with context
func (o *Options) UnmarshalJSONContext(ctx context.Context, data []byte) error {
	return json.Unmarshal(data, o)
}

// LogOptions represents logging configuration
type LogOptions struct {
	Disabled     bool   `json:"disabled,omitempty"`
	Level        string `json:"level,omitempty"`
	Output       string `json:"output,omitempty"`
	Timestamp    bool   `json:"timestamp,omitempty"`
}

// DNSOptions represents DNS configuration
type DNSOptions struct {
	Servers []DNSServerOptions `json:"servers,omitempty"`
	Rules   []DNSRule         `json:"rules,omitempty"`
}

// DNSServerOptions represents DNS server configuration
type DNSServerOptions struct {
	Tag     string `json:"tag,omitempty"`
	Address string `json:"address"`
}

// DNSRule represents DNS routing rule
type DNSRule struct {
	Domain []string `json:"domain,omitempty"`
	Server string   `json:"server,omitempty"`
}

// Inbound represents inbound configuration
type Inbound struct {
	Type    string      `json:"type"`
	Tag     string      `json:"tag,omitempty"`
	Options interface{} `json:",inline"`
}

// Outbound represents outbound configuration
type Outbound struct {
	Type    string      `json:"type"`
	Tag     string      `json:"tag,omitempty"`
	Options interface{} `json:",inline"`
}

// RouteOptions represents routing configuration
type RouteOptions struct {
	Rules []Rule `json:"rules,omitempty"`
}

// Rule represents routing rule
type Rule struct {
	Inbound  []string `json:"inbound,omitempty"`
	Outbound string   `json:"outbound,omitempty"`
}

// ExperimentalOptions represents experimental features
type ExperimentalOptions struct {
	V2RayAPI *V2RayStatsServiceOptions `json:"v2ray_api,omitempty"`
}

// V2RayStatsServiceOptions represents V2Ray stats service configuration
type V2RayStatsServiceOptions struct {
	Listen    string      `json:"listen,omitempty"`
	Stats     bool        `json:"stats,omitempty"`
	Enabled   bool        `json:"enabled,omitempty"`
	Outbounds []string    `json:"outbounds,omitempty"`
}