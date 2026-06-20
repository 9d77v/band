package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strconv"
)

type Sender struct {
	Conf Conf
}

func NewSender(conf Conf) *Sender {
	return &Sender{Conf: conf}
}

// buildMsg 构建邮件原始消息
func buildMsg(fromAddress, to, subject, body string) []byte {
	return []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", fromAddress, to, subject, body))
}

// isSSLPort 判断是否为 SSL/TLS 直连端口
func isSSLPort(portStr string) bool {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return false
	}
	return port == 465
}

// SendEmail 发送邮件
// to: 收件人邮箱
// subject: 邮件主题
// body: 邮件正文（HTML格式）
func (s *Sender) SendEmail(to, subject, body string) error {
	conf := s.Conf
	addr := fmt.Sprintf("%s:%s", conf.SMTPHost, conf.SMTPPort)
	msg := buildMsg(conf.FromAddress, to, subject, body)

	if isSSLPort(conf.SMTPPort) {
		return sendMailSSL(addr, conf.FromAddress, to, msg, conf.SMTPHost, conf.Username, conf.Password)
	}
	auth := smtp.PlainAuth("", conf.Username, conf.Password, conf.SMTPHost)
	return smtp.SendMail(addr, auth, conf.FromAddress, []string{to}, msg)
}

// sendMailSSL 通过 SSL/TLS 直连发送邮件（用于 465 端口）
func sendMailSSL(addr, from, to string, msg []byte, host, username, password string) error {
	tlsConfig := &tls.Config{ServerName: host}
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("tls dial failed: %w", err)
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("smtp new client failed: %w", err)
	}
	defer client.Close()

	auth := smtp.PlainAuth("", username, password, host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("smtp auth failed: %w", err)
	}
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("smtp mail from failed: %w", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("smtp rcpt failed: %w", err)
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("smtp data failed: %w", err)
	}
	if _, err = w.Write(msg); err != nil {
		return fmt.Errorf("write msg failed: %w", err)
	}
	if err = w.Close(); err != nil {
		return fmt.Errorf("close writer failed: %w", err)
	}
	return client.Quit()
}
