use db_im;

drop table `t_token`;
CREATE TABLE `t_token` (
	`t_token_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `t_token_user_id` varchar(40) NOT NULL COMMENT '授权码',
    `t_token_device_id` varchar(36) NOT NULL  COMMENT '设备id',
    `t_token_app_id` varchar(36) NOT NULL  COMMENT '应用id',
    `t_token_platform` varchar(10) NOT NULL  COMMENT '平台',
    `t_token_create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
    `t_token_login_time` datetime  COMMENT '登录日期',
    `t_token_logout_time` datetime COMMENT '登出日期',
     PRIMARY KEY (`t_token_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

drop table `t_session`;
create table `t_session`(
    `t_session_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `t_session_app_id` varchar(200) NOT NULL COMMENT '应用id',
    `t_session_create_user_id` varchar(200) NOT NULL COMMENT '创建用户id',
	PRIMARY KEY (`t_session_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会话表';

drop table `t_session_map`;
create table `t_session_map`(
    `t_session_map_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `t_session_map_session_id` bigint(20) NOT NULL ,
    `t_session_map_user_id` varchar(200) NOT NULL,
	PRIMARY KEY (`t_session_map_id`),
    UNIQUE KEY `t_session_session_user` (`t_session_map_session_id`,`t_session_map_user_id`)
    -- FOREIGN KEY `t_session_map_session_id` (`t_session_map_session_id`) REFERENCES `t_session` (`t_session_id`) ON DELETE CASCADE ON UPDATE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会话-用户映射表';

drop table `t_message`;
create table `t_message`(
    `t_message_id` varchar(40) NOT NULL COMMENT '主键',
    `t_message_session_id` varchar(40) NOT NULL COMMENT '用户id',
    `t_message_user_id` varchar(40) NOT NULL COMMENT '用户id',
    `t_message_type` int(4) NOT NULL COMMENT '消息类型',
    `t_message_content` varchar(20000) NOT NULL COMMENT '消息内容',
    `t_message_index` bigint(20) NOT NULL,
    `t_message_create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成时间',
    -- PRIMARY KEY (`t_message_id`),
    UNIQUE KEY `t_message_user_message_index` (`t_message_user_id`,`t_message_index`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='消息表';

start transaction;	
select max(t_message_index) from t_message where t_message_user_id = "1";
commit;


-- [CONSTRAINT symbol] FOREIGN KEY [id] (index_col_name, …)
-- REFERENCES tbl_name (index_col_name, …)
-- [ON DELETE {RESTRICT | CASCADE | SET NULL | NO ACTION | SET DEFAULT}]
-- [ON UPDATE {RESTRICT | CASCADE | SET NULL | NO ACTION | SET DEFAULT}]
-- ALTER TABLE news_info[子表名] ADD CONSTRAINT FK_news_info_news_type[约束名] FOREIGN KEY (info_id)[子表列] REFERENCES news_type[主表名] (id)[主表列] ; 
-- alter table `t_message` add constraint `t_message_session_id` foreign key `t_message_session_id` references `t_session` `t_session_id`; 

