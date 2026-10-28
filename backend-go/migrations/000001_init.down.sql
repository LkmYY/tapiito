-- 000001_init.down.sql
-- 回滚：按外键依赖逆序删除。生产环境执行前请确认已备份。

DROP TABLE IF EXISTS audit_log;
DROP TABLE IF EXISTS feedback;
DROP TABLE IF EXISTS `order`;
DROP TABLE IF EXISTS visit_stats_daily;
DROP TABLE IF EXISTS visit_stats;
DROP TABLE IF EXISTS photo;
DROP TABLE IF EXISTS note;
DROP TABLE IF EXISTS nfc_tag;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS admin_user;
