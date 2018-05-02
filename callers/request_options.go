package callers

//Defines a holder for request options.
type RequestOptions struct {
	URI         string
	Method      string
	Body        interface{}
	QueryParams map[string]string
}
