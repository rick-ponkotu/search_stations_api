mysql -uroot -pmysql --local-infile app -e "LOAD DATA LOCAL INFILE '/docker-entrypoint-initdb.d/station20210312free.csv' INTO TABLE table_station FIELDS TERMINATED BY ',' ENCLOSED BY '\"';"
mysql -uroot -pmysql --local-infile app -e "LOAD DATA LOCAL INFILE '/docker-entrypoint-initdb.d/line20210312free.csv' INTO TABLE table_line FIELDS TERMINATED BY ',' ENCLOSED BY '\"';"