package driver

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

func NewScheduler() *scheduler.Client {
	// SDK の設定をロードする
	// 環境変数、共有資格情報、共有設定ファイルから追加の設定と資格情報値をロードする
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("Failed to load SDK config, %v", err)
	}

	return scheduler.NewFromConfig(cfg)
}
