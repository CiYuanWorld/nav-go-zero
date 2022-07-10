-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `uuid` char(36) NOT NULL COMMENT 'UUID',
    `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
    `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像地址',
    `intro` varchar(255) NOT NULL DEFAULT '' COMMENT '一句话介绍',
    `status` smallint NOT NULL DEFAULT 0 COMMENT '资源状态:0默认',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_uuid` (`uuid`),
    UNIQUE KEY `uni_username` (`username`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci COMMENT='基础用户表';
