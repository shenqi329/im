use db_im;

select * from t_message;
select * from mysql.event;
show events;
SET GLOBAL event_scheduler = ON;

grant execute on db_im.* to im_connect@`%` identified by 'im_connect';
show grants for im_connect;
show full processlist;

show variables like '%max_connections%';
show global status like 'Max_used_connections';
set GLOBAL max_connections=1024;

drop procedure if exists t_message_insert;

DELIMITER //
create procedure t_message_insert(in message_id varchar(40),in session_id bigint(20),in user_id varchar(40),in type int(4) ,in context varchar(20000),out index_out bigint) 
begin 

declare oldindex bigint; 
start transaction; 
select max(t_message_index) into oldindex from t_message where t_message_session_id=session_id for update; 

if oldindex is NULL then 
set oldindex = 0;
end if;

insert into t_message(t_message_id,t_message_session_id,t_message_user_id,t_message_type,t_message_content,t_message_index) value(message_id,session_id,user_id,type,context,oldindex+1);
set index_out=oldindex+1;

commit; 
end;
//
DELIMITER ;

select * from t_message order by t_message_index desc limit 1000;
select count(t_message_index) from t_message;
select * from t_message order by t_message_index desc limit 1;


SET SQL_SAFE_UPDATES = 0;

-- SET @p_out=1; 
CALL t_message_get_increment_index(32,1,"a message for xxx",1,@p_out);
SELECT @p_out;

-- Error Code: 1414. OUT or INOUT argument 5 for routine db_im.t_message_get_increment_index is not a variable or NEW pseudo-variable in BEFORE trigger








