package config

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func Init() {

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(configYAML)); err != nil {
		log.Fatalln("FATAL", "Open config file error:", err)
	}
}

const (
	EnvironmentNameDev      = "dev"
	EnvironmentNameTest     = "test"
	EnvironmentNameStg      = "stg"
	EnvironmentNameProd     = "prod"
	CrawlerModeBatch        = "batch"
	CrawlerModeDaemon       = "daemon"
	CrawlerBatchKindReal    = "real"
	CrawlerBatchKindVirtual = "virtual"
)

func EnvironmentName() string {
	return viper.GetString("environment.name")
}

func LogDirectory() string {
	return fmt.Sprintf("%s_%s%s", viper.GetString("directory.base.prefix"), EnvironmentName(), viper.GetString("directory.log"))
}

func DataDirectory() string {
	return fmt.Sprintf("%s_%s%s", viper.GetString("directory.base.prefix"), EnvironmentName(), viper.GetString("directory.data"))
}

func CassandraHosts() string {
	return viper.GetString("cassandra.hosts")
}

func CassandraKeyspace() string {
	return fmt.Sprintf("%s_%s", viper.GetString("cassandra.keyspace.prefix"), EnvironmentName())
}

func KafkaBootstrapServers() string {
	return viper.GetString("kafka.bootstrap.servers")
}

func KafkaProcessorTopic() string {
	return fmt.Sprintf("%s_%s", viper.GetString("kafka.topics.processor.prefix"), EnvironmentName())
}

func KafkaAnalyzerTopic() string {
	return fmt.Sprintf("%s_%s", viper.GetString("kafka.topics.analyzer.prefix"), EnvironmentName())
}

func KafkaChooserTopic() string {
	return fmt.Sprintf("%s_%s", viper.GetString("kafka.topics.chooser.prefix"), EnvironmentName())
}

func KafkaNotifierTopic() string {
	return fmt.Sprintf("%s_%s", viper.GetString("kafka.topics.notifier.prefix"), EnvironmentName())
}

func CrawlerMode() string {
	return viper.GetString("crawler.mode")
}

func CrawlerBatchKind() string {
	return viper.GetString("crawler.batch.kind")
}

func CrawlerBatchStartTime() time.Time {
	return viper.GetTime("crawler.batch.start")
}

func CrawlerBatchEndTime() time.Time {
	return viper.GetTime("crawler.batch.end")
}

func TWSEURL() string {
	return viper.GetString("twse.url")
}

func TWSEReferer() string {
	return viper.GetString("twse.referer")
}

func TWSEDataDirectory() string {
	return fmt.Sprintf("%s%s", DataDirectory(), viper.GetString("twse.data"))
}

func TWSECron() string {
	return viper.GetString("twse.cron")
}

func TelegramToken() string {
	return viper.GetString("telegram.token")
}

func TelegramChatID() string {
	return viper.GetString("telegram.chatid")
}

func PorterV1Host() string {
	return viper.GetString("porter.v1.host")
}

func PorterV1Port() string {
	return viper.GetString("porter.v1.port")
}

func ShielderHost() string {
	return viper.GetString("shielder.host")
}

func ShielderPort() string {
	return viper.GetString("shielder.port")
}

func ShielderCORSMethods() []string {
	return viper.GetStringSlice("shielder.cors.methods")
}

func ShielderCORSOrigins() []string {
	return viper.GetStringSlice("shielder.cors.origins")
}
