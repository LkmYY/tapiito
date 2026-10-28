-- 访问控制：暗号 + 一次性 ticket + 水印 + 自动跳转
ALTER TABLE nfc_tag
  -- none 直出 / passcode 仅暗号 / ticket 仅一次性 / both 暗号换 ticket
  ADD COLUMN access_mode VARCHAR(16) NOT NULL DEFAULT 'none' COMMENT '访问控制模式',
  -- bcrypt hash，绝不存明文
  ADD COLUMN passcode_hash VARCHAR(72) NULL COMMENT '暗号 bcrypt 哈希',
  ADD COLUMN passcode_hint VARCHAR(64) NULL COMMENT '暗号输入提示语',
  ADD COLUMN ticket_ttl_seconds INT NOT NULL DEFAULT 3600 COMMENT 'ticket 有效期（秒）',
  ADD COLUMN auto_close_seconds INT NOT NULL DEFAULT 0 COMMENT '到点跳转品牌页（0=关闭）',
  ADD COLUMN watermark_enabled TINYINT NOT NULL DEFAULT 0 COMMENT '角标水印';

-- 一次性 ticket：绑定浏览器会话，跨会话失效（防转发）
CREATE TABLE IF NOT EXISTS access_ticket (
  id          BIGINT      NOT NULL AUTO_INCREMENT,
  nfc_tag_id  BIGINT      NOT NULL,
  ticket      CHAR(64)    NOT NULL COMMENT '随机 32 字节 hex',
  session_id  VARCHAR(64) NULL     COMMENT '首次使用时绑定的浏览器会话',
  issued_at   DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at  DATETIME    NOT NULL,
  used_at     DATETIME    NULL     COMMENT '首次使用时间',
  issued_ip   VARCHAR(64) NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uk_ticket (ticket),
  KEY ix_tag_expires (nfc_tag_id, expires_at),
  KEY ix_expires (expires_at),
  CONSTRAINT fk_ticket_tag FOREIGN KEY (nfc_tag_id) REFERENCES nfc_tag (id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT '一次性访问凭证';
