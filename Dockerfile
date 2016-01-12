FROM alpine:3.3

ADD mysql-query-metric-statsd /usr/local/bin/mysql-query-metric-statsd

ENTRYPOINT [ "/usr/local/bin/mysql-query-metric-statsd" ]
