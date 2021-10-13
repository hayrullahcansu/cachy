package request

type CreateCacheItemRequest struct {
	TimeSpan int         `json:"time_span"`
	Data     interface{} `json:"data,omitempty"`
}
