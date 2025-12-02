-- ============================================
-- donkey-admin-backend 数据库初始化脚本
-- ============================================

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS `donkey_user_center` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `donkey_user_center`;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
CREATE TABLE IF NOT EXISTS  `t_user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登陆名',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '称呼',
  `avatar` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `status` int NOT NULL DEFAULT '1' COMMENT '是否在职，1:正常，0:禁止',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_name` (`name`) USING BTREE,
  UNIQUE KEY `idx_phone` (`phone`) USING BTREE,
  UNIQUE KEY `idx_email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=117 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_verification
-- ----------------------------
CREATE TABLE IF NOT EXISTS `t_user_verification` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int unsigned NOT NULL DEFAULT '0' COMMENT '用户ID，0表示未注册用户（注册验证场景）',
  `type` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '验证类型: email_register, phone_register, email_bind, phone_bind, email_reset, phone_reset 等',
  `code` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '验证码或令牌',
  `target` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '目标值: 邮箱地址、手机号等',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-未使用, 1-已使用, 2-已过期',
  `expires_at` datetime NOT NULL COMMENT '过期时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `used_at` datetime DEFAULT NULL COMMENT '使用时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户验证码/令牌表';


SET FOREIGN_KEY_CHECKS = 1;

-- ============================================
-- 初始化完成
-- ============================================

