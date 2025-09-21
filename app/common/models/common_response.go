package models

// standard response
type CommonResponse struct {
	Status        int         `json:"status"`
	Error         bool        `json:"error"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
	EncryptedData interface{} `json:"encrypted_data,omitempty"`
	Meta          interface{} `json:"meta,omitempty"`
}
