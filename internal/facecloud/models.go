package facecloud

import "fmt"

type ResponseBase struct {
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"status_code,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

type ResponseDetect struct {
	ResponseBase

	Data []Data `json:"data,omitempty"`
}

type Data struct {
	BBox         BBox         `json:"bbox"`
	Demographics Demographics `json:"demographics"`
}

type BBox struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Height int `json:"height"`
	Width  int `json:"width"`
}

func (b *BBox) String() string {
	return fmt.Sprintf("%d %d %d %d", b.X, b.Y, b.Height, b.Width)
}

type Demographics struct {
	Age    Age    `json:"age"`
	Gender string `json:"gender"`
}

type Age struct {
	Mean     float64 `json:"mean"`
	Variance float64 `json:"variance"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	ResponseBase

	Data struct {
		AccessToken string `json:"access_token"`
	} `json:"data,omitempty"`
	StatusCode int `json:"status_code,omitempty"`
}
