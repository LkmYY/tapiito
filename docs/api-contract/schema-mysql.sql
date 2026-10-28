CREATE TABLE admin_user (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	username VARCHAR(64) NOT NULL, 
	password_hash VARCHAR(256) NOT NULL, 
	`role` VARCHAR(32), 
	created_at DATETIME, 
	PRIMARY KEY (id)
);

CREATE INDEX ix_admin_user_id ON admin_user (id);

CREATE UNIQUE INDEX ix_admin_user_username ON admin_user (username);

CREATE TABLE audit_log (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	user_type VARCHAR(16) NOT NULL, 
	username VARCHAR(64) NOT NULL, 
	action VARCHAR(64) NOT NULL, 
	resource_type VARCHAR(64), 
	resource_id VARCHAR(128), 
	detail TEXT, 
	ip_address VARCHAR(64), 
	created_at DATETIME, 
	PRIMARY KEY (id)
);

CREATE INDEX ix_audit_log_created_at ON audit_log (created_at);

CREATE INDEX ix_audit_log_id ON audit_log (id);

CREATE TABLE feedback (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	order_no VARCHAR(64), 
	content TEXT NOT NULL, 
	status VARCHAR(16), 
	created_ip VARCHAR(64), 
	created_at DATETIME, 
	adopted_at DATETIME, 
	PRIMARY KEY (id)
);

CREATE INDEX ix_feedback_order_no ON feedback (order_no);

CREATE INDEX ix_feedback_status ON feedback (status);

CREATE INDEX ix_feedback_id ON feedback (id);

CREATE INDEX ix_feedback_created_at ON feedback (created_at);

CREATE TABLE user (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	username VARCHAR(64) NOT NULL, 
	password_hash VARCHAR(256) NOT NULL, 
	phone VARCHAR(32), 
	email VARCHAR(128), 
	`role` VARCHAR(32), 
	created_at DATETIME, 
	updated_at DATETIME, 
	PRIMARY KEY (id), 
	UNIQUE (phone), 
	UNIQUE (email)
);

CREATE INDEX ix_user_id ON user (id);

CREATE UNIQUE INDEX ix_user_username ON user (username);

CREATE TABLE nfc_tag (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	nfc_id VARCHAR(128) NOT NULL, 
	name VARCHAR(128), 
	phone VARCHAR(32), 
	user_id INTEGER, 
	status VARCHAR(16), 
	template_type VARCHAR(64) NOT NULL, 
	custom_template VARCHAR(64), 
	template_payload JSON, 
	order_no VARCHAR(32), 
	created_at DATETIME, 
	updated_at DATETIME, 
	PRIMARY KEY (id), 
	FOREIGN KEY(user_id) REFERENCES user (id)
);

CREATE INDEX ix_nfc_tag_order_no ON nfc_tag (order_no);

CREATE INDEX ix_nfc_tag_id ON nfc_tag (id);

CREATE UNIQUE INDEX ix_nfc_tag_nfc_id ON nfc_tag (nfc_id);

CREATE TABLE note (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	nfc_tag_id INTEGER NOT NULL, 
	content TEXT NOT NULL, 
	created_at DATETIME, 
	created_ip VARCHAR(64), 
	PRIMARY KEY (id), 
	FOREIGN KEY(nfc_tag_id) REFERENCES nfc_tag (id)
);

CREATE INDEX ix_note_id ON note (id);

CREATE TABLE `order` (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	order_no VARCHAR(32) NOT NULL, 
	template_type VARCHAR(64) NOT NULL, 
	custom_template VARCHAR(64), 
	template_payload JSON, 
	tag_name VARCHAR(128), 
	external_order_no VARCHAR(64), 
	quantity INTEGER NOT NULL, 
	amount NUMERIC(10, 2) NOT NULL, 
	contact_name VARCHAR(64) NOT NULL, 
	contact_phone VARCHAR(32) NOT NULL, 
	contact_address VARCHAR(256) NOT NULL, 
	contact_zipcode VARCHAR(16), 
	contact_note VARCHAR(256), 
	status VARCHAR(16), 
	nfc_tag_id INTEGER, 
	nfc_id VARCHAR(128), 
	created_at DATETIME, 
	updated_at DATETIME, 
	PRIMARY KEY (id), 
	FOREIGN KEY(nfc_tag_id) REFERENCES nfc_tag (id)
);

CREATE INDEX ix_order_status ON `order` (status);

CREATE UNIQUE INDEX ix_order_order_no ON `order` (order_no);

CREATE INDEX ix_order_created_at ON `order` (created_at);

CREATE INDEX ix_order_id ON `order` (id);

CREATE TABLE photo (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	nfc_tag_id INTEGER NOT NULL, 
	oss_key VARCHAR(256) NOT NULL, 
	mime_type VARCHAR(64), 
	size INTEGER, 
	created_at DATETIME, 
	created_ip VARCHAR(64), 
	PRIMARY KEY (id), 
	FOREIGN KEY(nfc_tag_id) REFERENCES nfc_tag (id)
);

CREATE INDEX ix_photo_id ON photo (id);

CREATE TABLE visit_stats (
	id INTEGER NOT NULL AUTO_INCREMENT, 
	nfc_tag_id INTEGER NOT NULL, 
	ip_address VARCHAR(64), 
	user_agent VARCHAR(512), 
	referer VARCHAR(512), 
	country VARCHAR(64), 
	region VARCHAR(64), 
	city VARCHAR(64), 
	created_at DATETIME, 
	PRIMARY KEY (id), 
	FOREIGN KEY(nfc_tag_id) REFERENCES nfc_tag (id)
);

CREATE INDEX ix_visit_stats_id ON visit_stats (id);
