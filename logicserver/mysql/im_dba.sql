use db_im;

drop table `t_token`;
CREATE TABLE `t_token` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` varchar(40) NOT NULL COMMENT '授权码',
    `device_id` varchar(36) NOT NULL  COMMENT '设备id',
    `app_id` varchar(36) NOT NULL  COMMENT '应用id',
    `platform` varchar(10) NOT NULL  COMMENT '平台',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
    `login_time` datetime  COMMENT '登录日期',
    `logout_time` datetime COMMENT '登出日期',
    `sync_key` bigint(20) NOT NULL DEFAULT 0 COMMENT '同步key',
     PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

drop table `t_session`;
create table `t_session`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `app_id` varchar(200) NOT NULL COMMENT '应用id',
    `create_user_id` varchar(200) NOT NULL COMMENT '创建用户id',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会话表';

drop table `t_session_map`;
create table `t_session_map`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `session_id` bigint(20) NOT NULL ,
    `user_id` varchar(200) NOT NULL,
	PRIMARY KEY (`id`),
    UNIQUE KEY `session_user` (`session_id`,`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会话-用户映射表';

drop table `t_message`;
create table `t_message`(
    `id` varchar(40) NOT NULL COMMENT '主键',
    `session_id` varchar(40) NOT NULL COMMENT '用户id',
    `user_id` varchar(40) NOT NULL COMMENT '用户id',
    `type` int(4) NOT NULL COMMENT '消息类型',
    `content` varchar(20000) NOT NULL COMMENT '消息内容',
    `sync_key` bigint(20) NOT NULL,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
    UNIQUE KEY `user_message_index` (`user_id`,`sync_key`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='消息表';

start transaction;
select max(sync_key) from t_message where user_id = "1";
commit;

-- [CONSTRAINT symbol] FOREIGN KEY [id] (index_col_name, …)
-- REFERENCES tbl_name (index_col_name, …)
-- [ON DELETE {RESTRICT | CASCADE | SET NULL | NO ACTION | SET DEFAULT}]
-- [ON UPDATE {RESTRICT | CASCADE | SET NULL | NO ACTION | SET DEFAULT}]
-- ALTER TABLE news_info[子表名] ADD CONSTRAINT FK_news_info_news_type[约束名] FOREIGN KEY (info_id)[子表列] REFERENCES news_type[主表名] (id)[主表列] ; 
-- alter table `t_message` add constraint `t_message_session_id` foreign key `t_message_session_id` references `t_session` `t_session_id`; 

