# log 库（来自go-ethereum）

### 使用

1. 设置全局logger

```api
log.Root().SetHandler(log.NewFileLvlHandler("app.log", 1024*1024, "debug"))
```

2. 使用全局logger

```api
log.Info("some msg here","err",err)
```
