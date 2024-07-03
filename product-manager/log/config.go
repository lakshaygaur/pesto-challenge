package log

type LogEnvironment string

const (
	dev  LogEnvironment = "development"
	prod LogEnvironment = "production"
)

// config for logger
type Config struct {
	Environment LogEnvironment `jsons:"environment"`
}
