FROM mysql:5.7

COPY init/* /docker-entrypoint-initdb.d/
CMD ["mysqld"]

# MySQLをrootで実行する設定に書き換える
RUN sed -e 's/^ *user *= *mysql$/user = root/' -i /etc/mysql/my.cnf