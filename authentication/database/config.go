package database

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbName"`
}
