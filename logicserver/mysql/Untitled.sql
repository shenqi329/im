use db_im;

SET SQL_SAFE_UPDATES = 0;

SELECT * FROM db_im.t_message;

delete from t_message where t_message_user_id = '1';

begin;
select * from t_message where t_message_user_id = '1' order by t_message_index desc;
commit;

start transaction;	
select max(t_message_index) from t_message where t_message_user_id = "1";
commit;