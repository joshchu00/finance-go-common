package config

var configYAML = []byte(`
# environment.name: "dev", "test", "stg", "prod"
environment.name: "dev"

directory:
  base:
    prefix: "/var/lib/finance"
  log: "/log"
  data: "/data"
  
cassandra:
  hosts: "192.168.33.10:39042"
  keyspace:
    prefix: "finance"

kafka:
  bootstrap.servers: "192.168.33.10:39092"
  topics: 
    processor:
      prefix: "processor"
    analyzer:
      prefix: "analyzer"
    chooser:
      prefix: "chooser"
    notifier:
      prefix: "notifier"

crawler:
  # mode: "batch", "daemon"
  mode: "daemon"

  batch:
    # kind: "real", "virtual"
    kind: "real"
    start: "2018-11-12"
    end: "2018-11-16"

twse:
  url: "http://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=%s&type=ALLBUT0999&_=%d"
  referer: "http://www.twse.com.tw/zh/page/trading/exchange/MI_INDEX.html"
  data: "/twse"
  # cron: use UTC
  cron: "0 5 7-12 * * *"

porter:
  v1:
    host: "192.168.33.10"
    port: "51011"

shielder:
  host: "192.168.33.10"
  port: "58080"
  cors:
    methods:
      - "GET"
    origins:
      - "http://192.168.33.10:50080"
`)
