-- ----------------------------
-- Table structure for t_link
-- ----------------------------
DROP TABLE IF EXISTS `t_link`;
CREATE TABLE `t_link` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `uuid` char(36) NOT NULL COMMENT 'UUID',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '链接标题',
    `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '链接描述',
    `url` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址',
    `status` smallint NOT NULL DEFAULT 0 COMMENT '资源状态:0默认|1被用户删除',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_uuid` (`uuid`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci COMMENT='基础链接表';
