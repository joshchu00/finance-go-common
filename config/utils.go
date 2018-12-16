package config

import (
	"fmt"
	"log"
	"strings"

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

func CrawlerBatchType() string {
	return viper.GetString("crawler.batch.type")
}

func CrawlerBatchStartTime() (int, int, int) {
	start := viper.GetTime("crawler.batch.start")
	return start.Year(), int(start.Month()), start.Day()
}

func CrawlerBatchEndTime() (int, int, int) {
	end := viper.GetTime("crawler.batch.end")
	return end.Year(), int(end.Month()), end.Day()
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
