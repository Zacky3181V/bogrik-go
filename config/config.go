package config

import "os"

// config/config.go

type EmailConfig struct {
    From    string
    Subject string
    Body    string
}

type Config struct {
    Email EmailConfig
}

func Load() *Config {
    return &Config{
        Email: EmailConfig{
            From:    os.Getenv("EMAIL_FROM"),
            Subject: os.Getenv("EMAIL_SUBJECT"),
        },
    }
}