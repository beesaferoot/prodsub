package util

type ProductAttribute struct {
	FileSize           int64   `json:"file_size,string,omitempty"`
	DownloadLink       string  `json:"download_link,omitempty"`
	Weight             float64 `json:"weight,omitempty"`
	Dimensions         string  `json:"dimensions,omitempty"`
	SubscriptionPeriod string  `json:"subscription_period,omitempty"`
	RenewalPrice       float64 `json:"renewal_price,omitempty"`
}
