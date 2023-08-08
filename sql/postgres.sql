-- 修改postgres默认密码
ALTER USER postgres WITH PASSWORD 'xxx';

-- 创建数据库新用户，如 rental
CREATE USER rental WITH PASSWORD 'xxx';

-- 创建用户数据库，如rental
CREATE DATABASE rental OWNER postgres

-- 将rental数据库的所有权限都赋予xxx用户
GRANT ALL PRIVILEGES ON DATABASE rental TO xxx;

-- 将xxx用户设为超级用户
ALTER USER xxx WITH SUPERUSER;