ALTER DATABASE di default character set 'utf8';
GRANT ALL ON mysql.proc to 'admin'@'%';
GRANT ALL ON *.* to 'admin'@'%';
SET GLOBAL log_bin_trust_function_creators=1;