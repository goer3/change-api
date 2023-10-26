package initialize

import (
    "change-api/common"
    "change-api/pkg/log2"
    "fmt"
    "os"
    "time"

    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// 日志日期格式调整
func ZapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format(common.MsecTimeFormat))
}

// 日志组件格式化
func Logger() {
    // 新建配置
    config := zap.NewProductionEncoderConfig()
    // 设置时间格式
    config.EncodeTime = ZapLocalTimeEncoder
    // 关闭颜色输出
    config.EncodeLevel = zapcore.CapitalLevelEncoder
    // 新建输出
    var ws zapcore.WriteSyncer
    // 判断是否打印到日志文件
    if common.Config.Log.Enabled {
        // 日志文件
        now := time.Now()
        filename := fmt.Sprintf("%s/service.%04d-%02d-%02d.log",
            common.Config.Log.Path,
            now.Year(),
            now.Month(),
            now.Day())

        // 配置日志切割规则
        hook := &lumberjack.Logger{
            Filename:   filename,
            MaxSize:    common.Config.Log.MaxSize,
            MaxAge:     common.Config.Log.MaxAge,
            MaxBackups: common.Config.Log.MaxBackups,
            Compress:   common.Config.Log.Compress,
        }

        // 延时关闭
        defer func(hook *lumberjack.Logger) {
            _ = hook.Close()
        }(hook)

        // 配置输出到控制台和文件
        ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
    } else {
        // 配置只输出到控制台
        ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
    }

    // 整合日志输出信息
    core := zapcore.NewCore(zapcore.NewConsoleEncoder(config), ws, common.Config.Log.Level)
    logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

    // 配置全局 Log 工具，方便后面直接使用
    common.Log = logger.Sugar()
    log2.INFO("logger initialize success")
}
