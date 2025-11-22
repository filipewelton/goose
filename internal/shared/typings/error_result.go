package typings

type ErrorResult struct {
	Code    int    `json:"code"`
	Context string `json:"message"`
	Reason  string `json:"reason"`
}
