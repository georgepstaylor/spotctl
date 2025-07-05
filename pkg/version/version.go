package version

// Version information - these will be set via ldflags during build
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// GetUserAgent returns the User-Agent string for the CLI
func GetUserAgent() string {
	return "spotctl/" + Version
}
