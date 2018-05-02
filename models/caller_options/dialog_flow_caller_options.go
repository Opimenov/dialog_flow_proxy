//Holds default values for various API caller options
package caller_options

//Contains necessary values to make DialogFlow API call.
type AgentClientOptions struct {
	AccessToken string
	ApiLang     string
	ApiVersion  string
	ApiBaseUrl  string
	SessionID   string
}

//Dictionary of available languages
var AVAILABLE_LANGUAGES = map[string]string{
	"EN":    "en",
	"DE":    "de",
	"ES":    "es",
	"PT_BR": "pt-BR",
	"ZH_HK": "zh-HK",
	"ZH_CN": "zh-CN",
	"ZH_TW": "zh-TW",
	"FR":    "fr",
	"NL":    "nl",
	"IT":    "it",
	"JA":    "ja",
	"KO":    "ko",
	"RU":    "ru",
	"UK":    "uk",
	"PT":    "pt",
}

//Defines default AgentClientOptions values
const (
	//DialogFlow implementation version
	VERSION             = "2.0.0-beta.20"
	//DialogFlow API endpoint
	DEFAULT_BASE_URL    = "https://api.dialogflow.com/v1/"
	DEFAULT_API_VERSION = "20170712"
	DEFAULT_CLIENT_LANG = "en"
	API_KEY             = "a7cd00ab945249428a7f2f5841213fb3"
)
