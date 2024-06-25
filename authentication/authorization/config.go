package jwt

type Config struct {
	Secret string `json:"secret"`
	Expiry int    `json:"expiry"`
}
