package config

import (
	"fmt"
	"strconv"
)

type MailConfig struct {
	SmtpHost string
	SmtpPort int
	Username string
	Password string
	Sender   string
}

func NewMailConfig() (MailConfig, error) {
	var cfg MailConfig
	mailPort, err := strconv.Atoi(GetEnv("MAIL_SMTP_PORT", "25"))
	if err != nil {
		return cfg, fmt.Errorf("unable to parse mail port: %s", err.Error())
	}

	cfg = MailConfig{
		SmtpHost: GetEnv("MAIL_SMTP_HOST", "go-app"),
		SmtpPort: mailPort,
		Username: GetEnv("MAIL_SMTP_USERNAME", "smpt_user"),
		Password: GetEnv("MAIL_SMTP_PASSWORD", "smpt_pass"),
		Sender:   GetEnv("MAIL_SMTP_SENDER", ""),
	}

	return cfg, nil
}
