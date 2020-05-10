package check

// Property model describes a single key value pair for a service(parameters). An example could be a port for HTTP checking
type Property struct {
	Id int64 `json:"id,omitempty"`

	ServiceId int64 `json:"service_id"`

	Key string `json:"key,omitempty"`

	Value string `json:"value"`

	Status string `json:"status,omitempty"`
}
