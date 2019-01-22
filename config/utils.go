package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func Init() {

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("/var/lib/finance/config")
	if err := viper.ReadInConfig(); err != nil { // Find and read the config file and handle errors reading the config file
		log.Fatalln("FATAL", "Open config file error:", err)
	}
}

const (
	EnvironmentDev          = "dev"
	EnvironmentTest         = "test"
	EnvironmentStg          = "stg"
	EnvironmentProd         = "prod"
	CrawlerModeBatch        = "batch"
	CrawlerModeDaemon       = "daemon"
	CrawlerBatchKindReal    = "real"
	CrawlerBatchKindVirtual = "virutal"
)

func Environment() string {
	return viper.GetString("environment")
}

func ConfigDirectory() string {
	return fmt.Sprintf("%s%s", viper.GetString("directory.base"), viper.GetString("directory.config"))
}

func LogDirectory() string {
	return fmt.Sprintf("%s%s", viper.GetString("directory.base"), viper.GetString("directory.log"))
}

func DataDirectory() string {
	return fmt.Sprintf("%s%s", viper.GetString("directory.base"), viper.GetString("directory.data"))
}

func CassandraHosts() string {
	return viper.GetString("cassandra.hosts")
}

func CassandraKeyspace() string {
	return fmt.Sprintf("%s_%s", viper.GetString("cassandra.keyspace"), Environment())
}

func KafkaBootstrapServers() string {
	return viper.GetString("kafka.bootstrap.servers")
}

func KafkaProcessorTopic() string {
	return fmt.Sprintf("%s_%s", viper.GetString("kafka.topics.processor"), Environment())
}

func KafkaAnalyzerTopic() string {
	return fmt.Sprintf("%s_%s", viper.GetString("kafka.topics.analyzer"), Environment())
}

func KafkaChooserTopic() string {
	return fmt.Sprintf("%s_%s", viper.GetString("kafka.topics.chooser"), Environment())
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

func PorterV1Host() string {
	return viper.GetString("porter.v1.host")
}

func PorterV1Port() string {
	return viper.GetString("porter.v1.port")
}
