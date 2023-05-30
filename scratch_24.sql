create user actor identified by 'password';
create user actor;
show grants for actor;
grant all privileges on *.* to actor;
grant grant option on *.* to actor;
revoke all privileges on *.* from actor;
revoke grant option on *.* from actor;
revoke delete on *.* from actor;
alter user actor identified by 'new_password';

drop user superadmin;
create role superadmin;
grant all privileges on *.* to superadmin;
grant grant option on *.* to superadmin;

create user john;
grant 'superadmin' to john;
set default role all to john;
show grants for john;

revoke 'superadmin' from john;
