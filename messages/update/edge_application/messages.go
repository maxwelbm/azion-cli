package edge_application

var (
	Usage                       = "edge-application --id <application_id> [flags]"
	ShortDescription            = "Modifies an Edge Application"
	LongDescription             = "Modifies an Edge Application based on its ID to update its name, activity status, and other attributes"
	FlagID                      = "The Edge Application's id"
	FlagName                    = "The Edge Application's name"
	FlagDeliveryProtocol        = "The Edge Application's Delivery Protocol"
	FlagHttpPort                = "The Edge Application's Http Port"
	FlagHttpsPort               = "The Edge Application's Https Port"
	FlagMinimumTlsVersion       = "The Edge Application's Minimum Tls Version"
	FlagApplicationAcceleration = "Whether the Edge Application has Application Acceleration active or not"
	FlagCaching                 = "Whether the Edge Application has Caching active or not"
	FlagDeviceDetection         = "Whether the Edge Application has Device Detection active or not"
	FlagEdgeFirewall            = "Whether the Edge Application has Edge Firewall active or not"
	FlagEdgeFunctions           = "Whether the Edge Application has Edge Functions active or not"
	FlagImageOptimization       = "Whether the Edge Application has Image Optimization active or not"
	FlagL2Caching               = "Whether the Edge Application has L2 Caching active or not"
	FlagLoadBalancer            = "Whether the Edge Application has Load Balancer active or not"
	RawLogs                     = "Whether the Edge Application has Raw Logs active or not"
	WebApplicationFirewall      = "Whether the Edge Application has Web Application Firewall active or not"
	FlagIn                      = "Given path and JSON file to automatically update the Edge Application attributes; you can use - for reading from stdin"
	OutputSuccess               = "Updated Edge Application with ID %d\n"
	HelpFlag                    = "Displays more information about the update subcommand"
)
