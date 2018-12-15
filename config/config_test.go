package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	Init()
	fmt.Println(CrawlerBatchStartTime())
	fmt.Println(CrawlerBatchEndTime())
}
