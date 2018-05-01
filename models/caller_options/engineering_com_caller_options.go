//Holds default values for the Engineering.com API caller options
package caller_options

//Contains necessary values to make an Engineering.com API call.
type EngineeringClientOptions struct {
	AccessToken string
	ApiBaseUrl string
}

//AccessToken is associated with the user of Engineering.com.
//ApiBaseURL is Engineering.com API end point.
const (
	ENGINEERING_DEFAULT_BASE_URL  =  ""
	ENGINEERING_API_KEY           =  ""
)
