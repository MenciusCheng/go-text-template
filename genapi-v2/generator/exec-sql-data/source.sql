CREATE TABLE `platform_stat_api_relate_account` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `application_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '应用id',
    `conf_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '第三方api数据配置ID',
    `api_data_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '厂商类型',
    `account_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联账户类型,1:B站mid,2:B站内容视频广告账号',
    `account_id` varchar(255) NOT NULL DEFAULT '' COMMENT '关联账号id',
    `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注说明',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方api关联账户列表';
