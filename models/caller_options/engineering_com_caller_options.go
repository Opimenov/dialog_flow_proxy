//Holds default values for EngineeringClient API caller options
package caller_options

//Contains necessary values to make an Engineering.com API call.
type EngineeringClientOptions struct {
	AccessToken string
	ApiBaseUrl string
}

//Defines default EngineeringClientOptions values
//AccessToken is associated with the user of Engineering.com.
//ApiBaseURL is Engineering.com API end point.
const (
	ENGINEERING_DEFAULT_BASE_URL  =  ""
	ENGINEERING_API_KEY           =  ""
)
