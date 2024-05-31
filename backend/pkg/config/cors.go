package config

type CorsConfig struct {
	AllowedOrigin string
}

func NewCorsConfig() *CorsConfig {
	return &CorsConfig{
		AllowedOrigin: EnvOrDefault("ALLOWED_ORIGIN", defaultCfg.Cors.AllowedOrigin),
	}
}
