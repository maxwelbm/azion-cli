package contracts

import (
	"os"

	sdk "github.com/aziontech/azionapi-go-sdk/edgeapplications"
)

type FileOps struct {
	Path        string
	MimeType    string
	FileContent *os.File
	VersionID   string
}

type BuildInfo struct {
	Preset        string
	Mode          string
	Entry         string
	NodePolyfills string
	OwnWorker     string
	IsFirewall    bool
}

type DevInfo struct {
	IsFirewall string
}

type ListOptions struct {
	Details           bool
	OrderBy           string
	Sort              string
	Page              int64
	PageSize          int64
	Filter            string
	ContinuationToken string
}

type DescribeOptions struct {
	OutPath string
	Format  string
}

type AzionApplicationOptions struct {
	Test          func(path string) error      `json:"-"`
	Name          string                       `json:"name"`
	Bucket        string                       `json:"bucket"`
	Preset        string                       `json:"preset"` // framework: react, next, vue, angular and etc
	Mode          string                       `json:"mode"`   // deliver == ssg, compute == ssr
	Env           string                       `json:"env"`
	Prefix        string                       `json:"prefix"`
	NotFirstRun   bool                         `json:"not-first-run"`
	Function      AzionJsonDataFunction        `json:"function"`
	Application   AzionJsonDataApplication     `json:"application"`
	Domain        AzionJsonDataDomain          `json:"domain"`
	RtPurge       AzionJsonDataPurge           `json:"rt-purge"`
	Origin        []AzionJsonDataOrigin        `json:"origin"`
	RulesEngine   AzionJsonDataRulesEngine     `json:"rules-engine"`
	CacheSettings []AzionJsonDataCacheSettings `json:"cache-settings"`
}

type AzionApplicationSimple struct {
	Name        string                   `json:"name"`
	Type        string                   `json:"type"`
	Domain      AzionJsonDataDomain      `json:"domain"`
	Application AzionJsonDataApplication `json:"application"`
}

type AzionApplicationConfig struct {
	InitData    InitConf    `json:"init"`
	BuildData   BuildConf   `json:"build"`
	PublishData PublishConf `json:"publish"`
}

type InitConf struct {
	Cmd        string `json:"cmd"`
	Env        string `json:"env"`
	OutputCtrl string `json:"output-ctrl"`
	Default    string `json:"default"`
}

type BuildConf struct {
	Cmd        string `json:"cmd"`
	Env        string `json:"env"`
	OutputCtrl string `json:"output-ctrl"`
	Default    string `json:"default"`
}

type PublishConf struct {
	Cmd        string `json:"pre_cmd"`
	Env        string `json:"env"`
	OutputCtrl string `json:"output-ctrl"`
	Default    string `json:"default"`
}

type CacheConf struct {
	PurgeOnPublish bool `json:"purge_on_publish"`
}

type AzionJsonDataFunction struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	File         string `json:"file"`
	Args         string `json:"args"`
	InstanceID   int64  `json:"instance-id"`
	InstanceName string `json:"instance-name"`
	CacheId      int64  `json:"cache-id"`
}

type AzionJsonDataApplication struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type AzionJsonDataOrigin struct {
	OriginId  int64    `json:"origin-id"`
	OriginKey string   `json:"origin-key"`
	Name      string   `json:"name"`
	Address   []string `json:"address,omitempty"`
}

type AzionJsonDataDomain struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	DomainName string `json:"domain_name"`
	Url        string `json:"url"`
}

type AzionJsonDataPurge struct {
	PurgeOnPublish bool `json:"purge_on_publish"`
}

type AzionJsonDataRulesEngine struct {
	Created bool                 `json:"created"`
	Rules   []AzionJsonDataRules `json:"rules"`
}

type AzionJsonDataRules struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type AzionJsonDataCacheSettings struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Manifest struct {
	CacheSettings []CacheSetting `json:"cache"`
	Origins       []Origin       `json:"origin"`
	Rules         []RuleEngine   `json:"rules"`
}

type CacheSetting struct {
	Name                           *string  `json:"name,omitempty"`
	BrowserCacheSettings           *string  `json:"browser_cache_settings,omitempty"`
	BrowserCacheSettingsMaximumTtl *int64   `json:"browser_cache_settings_maximum_ttl,omitempty"`
	CdnCacheSettings               *string  `json:"cdn_cache_settings,omitempty"`
	CdnCacheSettingsMaximumTtl     *int64   `json:"cdn_cache_settings_maximum_ttl,omitempty"`
	CacheByQueryString             *string  `json:"cache_by_query_string,omitempty"`
	QueryStringFields              []string `json:"query_string_fields,omitempty"`
	EnableQueryStringSort          *bool    `json:"enable_query_string_sort,omitempty"`
	CacheByCookies                 *string  `json:"cache_by_cookies,omitempty"`
	CookieNames                    []string `json:"cookie_names,omitempty"`
	AdaptiveDeliveryAction         *string  `json:"adaptive_delivery_action,omitempty"`
	DeviceGroup                    []int32  `json:"device_group,omitempty"`
	EnableCachingForPost           *bool    `json:"enable_caching_for_post,omitempty"`
	L2CachingEnabled               *bool    `json:"l2_caching_enabled,omitempty"`
	IsSliceConfigurationEnabled    *bool    `json:"is_slice_configuration_enabled,omitempty"`
	IsSliceEdgeCachingEnabled      *bool    `json:"is_slice_edge_caching_enabled,omitempty"`
	IsSliceL2CachingEnabled        *bool    `json:"is_slice_l2_caching_enabled,omitempty"`
	SliceConfigurationRange        *int64   `json:"slice_configuration_range,omitempty"`
	EnableCachingForOptions        *bool    `json:"enable_caching_for_options,omitempty"`
	EnableStaleCache               *bool    `json:"enable_stale_cache,omitempty"`
	L2Region                       *string  `json:"l2_region,omitempty"`
}

type Origin struct {
	Name       string `json:"name"`
	OriginType string `json:"origin_type,omitempty"`
	Bucket     string `json:"bucket,omitempty"`
	Prefix     string `json:"prefix,omitempty"`
}

type RuleEngine struct {
	Name        string                         `json:"name"`
	Description *string                        `json:"description,omitempty"`
	Criteria    [][]sdk.RulesEngineCriteria    `json:"criteria,omitempty"`
	Behaviors   []sdk.RulesEngineBehaviorEntry `json:"behaviors,omitempty"`
}

type RulesEngineBehaviorEntry struct {
	RulesEngineBehaviorObject *sdk.RulesEngineBehaviorObject
	RulesEngineBehaviorString *sdk.RulesEngineBehaviorString
}
