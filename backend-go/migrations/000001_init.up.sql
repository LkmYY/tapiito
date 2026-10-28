-- 000001_init.up.sql
-- Tapiito 初始表结构：由 FastAPI 时期的 ensure_schema() 沉淀为正式 migration。
-- 约定：全部 InnoDB + utf8mb4；时间列保留 DATETIME（与旧数据一致）。

CREATE TABLE IF NOT EXISTS admin_user (
  id            BIGINT       NOT NULL AUTO_INCREMENT,
  username      VARCHAR(64)  NOT NULL,
  password_hash VARCHAR(256) NOT NULL,
  `role`        VARCHAR(32)  NOT NULL DEFAULT 'admin',
  created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_admin_user_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `user` (
  id            BIGINT       NOT NULL AUTO_INCREMENT,
  username      VARCHAR(64)  NOT NULL,
  password_hash VARCHAR(256) NOT NULL,
  phone         VARCHAR(32)  NULL,
  email         VARCHAR(128) NULL,
  avatar        VARCHAR(256) NULL,
  `status`      VARCHAR(16)  NOT NULL DEFAULT 'active',
  created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_user_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 名片主表：公开读的唯一起点，nfc_id 必须唯一
CREATE TABLE IF NOT EXISTS nfc_tag (
  id               BIGINT       NOT NULL AUTO_INCREMENT,
  nfc_id           VARCHAR(128) NOT NULL,
  name             VARCHAR(128) NULL,
  phone            VARCHAR(32)  NULL,
  user_id          BIGINT       NULL,
  `status`         VARCHAR(16)  NOT NULL DEFAULT 'enabled',
  template_type    VARCHAR(64)  NOT NULL,
  custom_template  VARCHAR(64)  NULL,
  template_payload JSON         NULL,
  order_no         VARCHAR(32)  NULL,
  created_at       DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at       DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_nfc_tag_nfc_id (nfc_id),
  KEY ix_nfc_tag_status_created (status, id),
  KEY ix_nfc_tag_user_created (user_id, id),
  KEY ix_nfc_tag_order_no (order_no),
  CONSTRAINT fk_nfc_tag_user FOREIGN KEY (user_id) REFERENCES `user` (id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS note (
  id          BIGINT      NOT NULL AUTO_INCREMENT,
  nfc_tag_id  BIGINT      NOT NULL,
  content     TEXT        NOT NULL,
  created_ip  VARCHAR(64) NULL,
  created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY ix_note_tag_created (nfc_tag_id, id),
  CONSTRAINT fk_note_tag FOREIGN KEY (nfc_tag_id) REFERENCES nfc_tag (id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS photo (
  id          BIGINT       NOT NULL AUTO_INCREMENT,
  nfc_tag_id  BIGINT       NOT NULL,
  oss_key     VARCHAR(512) NOT NULL,
  mime_type   VARCHAR(128) NULL,
  `size`      BIGINT       NULL,
  created_ip  VARCHAR(64)  NULL,
  created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY ix_photo_tag_created (nfc_tag_id, id),
  CONSTRAINT fk_photo_tag FOREIGN KEY (nfc_tag_id) REFERENCES nfc_tag (id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 访问明细：仅保留近期数据，长期统计看 visit_stats_daily
CREATE TABLE IF NOT EXISTS visit_stats (
  id          BIGINT       NOT NULL AUTO_INCREMENT,
  nfc_tag_id  BIGINT       NOT NULL,
  ip_address  VARCHAR(64)  NULL,
  user_agent  VARCHAR(512) NULL,
  referer     VARCHAR(512) NULL,
  -- 与 FastAPI 的 VisitStats 模型保持一致（统计接口会返回这三个字段）
  country     VARCHAR(64)  NULL,
  region      VARCHAR(64)  NULL,
  city        VARCHAR(64)  NULL,
  created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY ix_visit_stats_tag_created (nfc_tag_id, created_at),
  KEY ix_visit_stats_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 日聚合表：管理后台统计查这里，不再 COUNT(*) 扫明细
CREATE TABLE IF NOT EXISTS visit_stats_daily (
  id          BIGINT   NOT NULL AUTO_INCREMENT,
  nfc_tag_id  BIGINT   NOT NULL,
  stat_date   DATE     NOT NULL,
  `count`     BIGINT   NOT NULL DEFAULT 0,
  updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_visit_daily_tag_date (nfc_tag_id, stat_date),
  KEY ix_visit_daily_date (stat_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 硬件订单
CREATE TABLE IF NOT EXISTS `order` (
  id                BIGINT       NOT NULL AUTO_INCREMENT,
  order_no          VARCHAR(32)  NOT NULL,
  template_type     VARCHAR(64)  NOT NULL,
  custom_template   VARCHAR(64)  NULL,
  template_payload  JSON         NULL,
  tag_name          VARCHAR(128) NULL,
  external_order_no VARCHAR(64)  NULL,
  quantity          INT          NOT NULL DEFAULT 1,
  amount            DECIMAL(10,2) NOT NULL DEFAULT 0.00,
  contact_name      VARCHAR(64)  NOT NULL,
  contact_phone     VARCHAR(32)  NOT NULL,
  contact_address   VARCHAR(256) NOT NULL,
  contact_zipcode   VARCHAR(16)  NULL,
  contact_note      VARCHAR(256) NULL,
  `status`          VARCHAR(16)  NOT NULL DEFAULT 'pending',
  nfc_tag_id        BIGINT       NULL,
  nfc_id            VARCHAR(128) NULL,
  created_at        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY uk_order_no (order_no),
  KEY ix_order_status_created (status, id),
  KEY ix_order_external (external_order_no),
  KEY ix_order_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 用户意见
CREATE TABLE IF NOT EXISTS feedback (
  id          BIGINT      NOT NULL AUTO_INCREMENT,
  order_no    VARCHAR(64) NULL,
  content     TEXT        NOT NULL,
  `status`    VARCHAR(16) NOT NULL DEFAULT 'pending',
  created_ip  VARCHAR(64) NULL,
  created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  adopted_at  DATETIME    NULL,
  PRIMARY KEY (id),
  KEY ix_feedback_status_created (status, id),
  KEY ix_feedback_order_no (order_no)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 审计日志
CREATE TABLE IF NOT EXISTS audit_log (
  id            BIGINT       NOT NULL AUTO_INCREMENT,
  user_type     VARCHAR(16)  NOT NULL,
  username      VARCHAR(64)  NOT NULL,
  action        VARCHAR(64)  NOT NULL,
  resource_type VARCHAR(64)  NULL,
  resource_id   VARCHAR(128) NULL,
  detail        TEXT         NULL,
  ip_address    VARCHAR(64)  NULL,
  created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY ix_audit_log_created (created_at),
  KEY ix_audit_log_user (username, created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
