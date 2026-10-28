// Package sms 通过 Twilio REST API 发送短信/彩信通知。
//
// 未配置 Twilio 时不报错，返回 sent=false，避免公开提交接口因外部依赖不可用而失败。
package sms

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lkmyy/tapiito-go/internal/config"
)

type Sender struct {
	cfg    config.TwilioConfig
	client *http.Client
	log    *slog.Logger
}

func New(cfg config.TwilioConfig, log *slog.Logger) *Sender {
	return &Sender{
		cfg:    cfg,
		client: &http.Client{Timeout: 8 * time.Second},
		log:    log,
	}
}

// Enabled 表示是否配置了完整凭证。
func (s *Sender) Enabled() bool {
	return s.cfg.AccountSID != "" && s.cfg.AuthToken != "" && s.cfg.From != ""
}

// Send 发送短信；mediaURLs 非空时按彩信发送。返回是否成功。
func (s *Sender) Send(ctx context.Context, to, body string, mediaURLs []string) bool {
	if !s.Enabled() {
		s.log.Debug("twilio not configured, skip sending", "to", to)
		return false
	}
	endpoint := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", s.cfg.AccountSID)

	form := url.Values{}
	form.Set("To", to)
	form.Set("From", s.cfg.From)
	form.Set("Body", body)
	for _, m := range mediaURLs {
		form.Add("MediaUrl", m)
	}

	// 有限重试：外部服务抖动不应无限挂住请求
	var lastErr error
	for attempt := 0; attempt < 2; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint,
			strings.NewReader(form.Encode()))
		if err != nil {
			return false
		}
		req.SetBasicAuth(s.cfg.AccountSID, s.cfg.AuthToken)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := s.client.Do(req)
		if err == nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				return true
			}
			lastErr = fmt.Errorf("twilio status %d", resp.StatusCode)
		} else {
			lastErr = err
		}
	}
	s.log.Error("twilio send failed", "to", to, "err", lastErr)
	return false
}
