use db_im;

SET SQL_SAFE_UPDATES = 0;

select * from t_message where t_message_user_id = 1 order by t_message_index desc;


drop procedure if exists t_message_create_message_test;
DELIMITER //
create procedure t_message_create_message_test(in user_id varchar(40),in session_id bigint(20), in type int(4) ,in context varchar(20000),in count int(4))
begin
declare num int;
declare oldindex bigint;
set num = 1;
while num <= count do

	start transaction;
	
    select max(t_message_index) into oldindex from t_message where t_message_user_id = user_id for update;
	if oldindex is NULL then 
		set oldindex = 0;
	end if;
	set oldindex = oldindex+1;

	insert into t_message(t_message_id,t_message_user_id,t_message_session_id,t_message_index,t_message_type,t_message_content) values(replace(uuid(),'-',''),user_id,session_id,oldindex,type,context);
	set num = num +1;
    
    commit;
end while;


end;
//
DELIMITER ;

delete from t_message;
select max(t_message_index) from t_message where t_message_user_id = "1";
call t_message_create_message_test("1",33,1,"a messge for test 1",100000);

start transaction;	
select max(t_message_index) from t_message where t_message_user_id = "1" for update;

commit;

