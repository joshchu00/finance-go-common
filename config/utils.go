package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func init() {

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
	return fmt.Sprintf("%s/%s", viper.GetString("directory.base"), viper.GetString("directory.config"))
}

func LogDirectory() string {
	return fmt.Sprintf("%s/%s", viper.GetString("directory.base"), viper.GetString("directory.log"))
}

func DataDirectory() string {
	return fmt.Sprintf("%s/%s", viper.GetString("directory.base"), viper.GetString("directory.data"))
}

func CassandraHosts() string {
	return viper.GetString("cassandra.hosts")
}

func CassandraKeyspace() string {
	return fmt.Sprintf("%s_%s", viper.GetString("cassandra.hosts"), Environment())
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

func CrawlerBatchStartTime() (int, int, int) {
	return viper.GetInt("crawler.batch.start.year"), viper.GetInt("crawler.batch.start.month"), viper.GetInt("crawler.batch.start.day")
}

func CrawlerBatchEndTime() (int, int, int) {
	return viper.GetInt("crawler.batch.end.year"), viper.GetInt("crawler.batch.end.month"), viper.GetInt("crawler.batch.end.day")
}

func CrawlerDaemonCron() string {
	return viper.GetString("crawler.daemon.cron")
}

func TWSEURL() string {
	return viper.GetString("twse.url")
}

func TWSEReferer() string {
	return viper.GetString("twse.referer")
}

func TWSEDataDirectory() string {
	return fmt.Sprintf("%s/%s", DataDirectory(), viper.GetString("directory.data"))
}
