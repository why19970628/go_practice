CREATE TABLE `bf_slot_list`
(
    `id`             bigint(11) NOT NULL AUTO_INCREMENT,
    `machine_id`     bigint(20) NOT NULL COMMENT '机器序号',
    `machine_name`   varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '机器名称',
    `require_level`  bigint(20) NOT NULL,
    `card_type`      varchar(64)                                                   NOT NULL COMMENT '类型',
    `is_temp_unlock` tinyint(1),
    `add_time`       bigint(20) DEFAULT NULL,
    `update_time`    bigint(20) DEFAULT NULL,
    `del_time`       bigint(20) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY              `machine_id` (`machine_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='大厅机器排序，VIP和非VIP分开排列';