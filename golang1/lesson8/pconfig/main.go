package pconfig

import (
	"flag"
	"fmt"
	"net/url"
)

type Config struct {
	port         int
	dbUrl        string
	jaegerUrl    string
	sentryUrl    string
	kafkaBrocker string
	appId        int
	appKey       string
}

func Configure() (Config, error) {
	var cfg = Config{}
	flag.IntVar(&cfg.port, "port", 8080, "port 0-65535")
	flag.StringVar(&cfg.dbUrl, "dburl", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "database url")
	flag.StringVar(&cfg.jaegerUrl, "jaeger", "http://jaeger:16686", "jaeger url")
	flag.StringVar(&cfg.sentryUrl, "sentry", "http://sentry:9000", "sentry url")
	flag.StringVar(&cfg.kafkaBrocker, "kafka", "http://kafka:9092", "kafka broker")
	flag.IntVar(&cfg.appId, "appid", 10000001, "app id")
	flag.StringVar(&cfg.appKey, "appkey", "Dagsdhgeyr273rrn23jurh2u3rh", "app key")
	flag.Parse()
	cfg, err := validate(cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
func validate(cfg Config) (Config, error) {
	if cfg.port > 65535 || cfg.port < 0 {
		return Config{}, fmt.Errorf("Неверное значение port")
	}
	if !isUrl(cfg.dbUrl) {
		return Config{}, fmt.Errorf("Неверное значение dbUrl")
	}
	if !isUrl(cfg.jaegerUrl) {
		return Config{}, fmt.Errorf("Неверное значение jaegerUrl")
	}
	if !isUrl(cfg.sentryUrl) {
		return Config{}, fmt.Errorf("Неверное значение sentryUrl")
	}
	if !isUrl(cfg.kafkaBrocker) {
		return Config{}, fmt.Errorf("Неверное значение kafkaBrocker")
	}
	if cfg.appId < 1000000 {
		return Config{}, fmt.Errorf("Неверное значение appId")
	}
	return cfg, nil
}
func isUrl(str string) bool {
	parsedUrl, err := url.Parse(str)
	return err == nil && parsedUrl.Scheme != "" && parsedUrl.Host != ""
}
