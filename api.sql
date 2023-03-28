CREATE TABLE `users`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `nickname`     varchar(128) not null default '' comment '昵称',
    `gender`       tinyint      not null default 0 comment '性别，1：男，2：女',
    `phone`        char(11)     not null default 0 comment '手机号',
    `country_code` char(12)     not null default 0 comment '手机号前缀 eg:86',
    `password`     char(32)     NOT NULL DEFAULT '' comment '密码',
    `birthday`     char(10)     not null default '' comment '生日',
    `portrait`     varchar(256) not null default '' comment '头像',
    `introduction` varchar(512) not null default '' comment '简介',
    `token`        char(32)     not null default '' comment 'token',
    `users_type`   tinyint      not null default 0 comment '1：自然用户，2：马甲用户',
    `login_time`   int          not null default 0 comment '登录时间',
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY            idx_phone(`phone`),
    KEY            idx_token (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `posting`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `uid`               bigint       not null default 0 comment '用户id',
    `content`           text null comment '帖子内容',
    `images`            varchar(512) not null default '' comment '图片',
    `score`             tinyint      not null default 0 comment '质量分，1:S,2:A,3:B,4:C',
    `attribute`         tinyint      not null default 0 comment '属性，1：自然贴，2：马甲贴',
    `posting_type`      tinyint      not null default 0 comment '类型，1：普通，2：精选',
    `audit_status`      tinyint      not null default 0 comment '审核状态，1：未审核，2：审核成功，10：审核失败',
    `audit_fail_reason` varchar(512) not null default '' comment '审核失败原因',
    `status`            tinyint      not null default 0 comment '状态，1：未发布，2:已发布，101:已删除',
    `create_time`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                 idx_uid(`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子表';

CREATE TABLE `posting_subject`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `posting_id`  bigint    not null default 0 comment '帖子id',
    `subject_id`  bigint    not null default 0 comment '话题id',
    `status`      tinyint   not null default 0 comment '状态，1:正常，101:删除',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           idx_posting_id_uid(`posting_id`, `uid`),
    KEY           idx_subject_id(`subject_id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='话题帖子关联表';

CREATE TABLE `subject`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `name`        varchar(128) not null default '' comment '话题名称',
    `status`      tinyint      not null default 0 comment '状态，1:正常，101:删除',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='话题表';

CREATE TABLE `posting_like_record`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `uid`         bigint    not null default 0 comment '用户id',
    `posting_id`  bigint    not null default 0 comment '帖子id',
    `status`      tinyint   not null default 0 comment '状态，1:已点赞，101:已取消',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           idx_posting_id_uid(`posting_id`, `uid`),
    KEY           idx_uid(`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子点赞记录表';

CREATE TABLE `posting_comment`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `code`              varchar(128) not null default '' comment '评论编号',
    `uid`               bigint       not null default 0 comment '用户id',
    `posting_id`        bigint       not null default 0 comment '帖子id',
    `pid`               bigint       not null default 0 comment '父级id',
    `level`             tinyint      not null default 0 comment '代表第几级评论',
    `attribute`         tinyint      not null default 0 comment '属性，1：自然人，2：马甲人',
    `content`           text null comment '评论内容',
    `audit_status`      tinyint      not null default 0 comment '审核状态，1：未审核，2：审核通过，10：审核未通过',
    `audit_fail_reason` varchar(512) not null default '' comment '审核未通过原因',
    `status`            tinyint      not null default 0 comment '状态，1：未发布，2:已发布，101:已删除',
    `create_time`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                 idx_posting_id_uid(`posting_id`, `uid`),
    KEY                 idx_uid(`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

CREATE TABLE `posting_comment_like_record`
(
    `id`                 bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `uid`                bigint    not null default 0 comment '用户id',
    `posting_id`         bigint    not null default 0 comment '帖子id',
    `posting_comment_id` bigint    not null default 0 comment '评论id',
    `status`             tinyint   not null default 0 comment '状态，1:已点赞，101:已取消',
    `create_time`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                  idx_posting_id_uid(`posting_id`, `uid`),
    KEY                  idx_uid(`uid`),
    KEY                  idx_posting_comment_id(`posting_comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子评论点赞记录表';

CREATE TABLE `posting_statistics`
(
    `id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `posting_id`      bigint    not null default 0 comment '帖子id',
    `statistics_type` tinyint   not null default 0 comment '统计类型，1：自然人点赞数量，2：全部点赞数量，3：自然人评论数量，4：全部评论数量',
    `num`             int       not null default 0 comment '数量',
    `create_time`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY               idx_posting_id_statistics_type(`posting_id`, `statistics_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子统计表';

CREATE TABLE `posting_comment_statistics`
(
    `id`                 bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `posting_id`         bigint    not null default 0 comment '帖子id',
    `posting_comment_id` bigint    not null default 0 comment '评论id',
    `statistics_type`    tinyint   not null default 0 comment '统计类型，1：自然人点赞数量，2：全部点赞数量，3：自然人评论数量，4：全部评论数量',
    `num`                int       not null default 0 comment '数量',
    `create_time`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY                  idx_posting_comment_statistics(`posting_comment_id`, `statistics_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子评论统计表';
