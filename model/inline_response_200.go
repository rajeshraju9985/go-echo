package model

type InlineResponse200 struct {
	Greeting string `json:"greeting,omitempty"`

	Date string `json:"date,omitempty"`

	URL string `json:"url,omitempty"`
}
