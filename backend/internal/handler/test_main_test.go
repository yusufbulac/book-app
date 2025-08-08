package handler_test

import (
	"os"
	"testing"

	"github.com/yusufbulac/byfood-case/backend/pkg/logger"
)

func TestMain(m *testing.M) {
	// init global zap logger used by middlewares
	logger.InitLogger()
	code := m.Run()
	_ = logger.Log.Sync()
	os.Exit(code)
}
