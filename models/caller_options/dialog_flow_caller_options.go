package caller_options

type AgentClientOptions struct {
	AccessToken string
	ApiLang     string
	ApiVersion  string
	ApiBaseUrl  string
	SessionID   string
}

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

const (
	VERSION             = "2.0.0-beta.20"
	DEFAULT_BASE_URL    = "https://api.dialogflow.com/v1/"
	DEFAULT_API_VERSION = "20170712"
	DEFAULT_CLIENT_LANG = "en"
	API_KEY             = "a7cd00ab945249428a7f2f5841213fb3"
)
