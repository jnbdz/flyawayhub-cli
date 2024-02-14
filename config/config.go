package config

const (
	AppName     = "Flyawayhub CLI"
	SlugAppName = "flyawayhub-cli"

	// APIBaseURL is the base URL for the Flyawayhub API.
	APIBaseURL = "https://api.prod.flyawayhub.com/"

	// APIVersion is the version of the API to use.
	APIVersion = "v1"

	// UserAgent is the user agent string to use for HTTP requests.
	UserAgent = "Mozilla/5.0"

	AWSCognitoRegion   = "us-west-2"
	AWSCognitoClientID = ""

	//req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0")
	//req.Header.Set("Host", "api.prod.flyawayhub.com")
	//req.Header.Set("Origin", "https://app.prod.flyawayhub.com")
	//req.Header.Set("Referer", "https://app.prod.flyawayhub.com/")
)

// APIEndpoint returns the full URL for a given endpoint.
func APIEndpoint(endpoint string) string {
	return APIBaseURL + APIVersion + "/" + endpoint
}
