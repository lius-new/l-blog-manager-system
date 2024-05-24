package svc

import (
	"github.com/lius-new/blog-backend/rpc/analyzer/internal/config"
	blockedModel "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/blocked"
	recordModel "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/record"
	settingModel "github.com/lius-new/blog-backend/rpc/analyzer/model/mongo/setting"
)

type ServiceContext struct {
	Config           config.Config
	ModelWithRecord  recordModel.RecordModel
	ModelWithBlocked blockedModel.BlockedModel
	ModelWithSetting settingModel.SettingModel
	// Utiler           utiler.Utiler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ModelWithRecord:  recordModel.NewRecordModel(c.MongoURL, c.DBName, "record", c.Cache),
		ModelWithBlocked: blockedModel.NewBlockedModel(c.MongoURL, c.DBName, "blocked", c.Cache),
		ModelWithSetting: settingModel.NewSettingModel(c.MongoURL, c.DBName, "analyzer_setting", c.Cache),
		// Utiler:           utiler.NewUtiler(zrpc.MustNewClient(c.Utils)),
	}
}
