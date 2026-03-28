# tool-kit

## Build & Run

```bash
go build -o bin/tool .
go test ./...
go mod tidy
```

## Conventions

- CLI: `github.com/urfave/cli/v2`，命令通过 `init()` 自注册到 `rootCmd.Commands`
- `str/json/ts` 无状态，跳过 config；`dts/docx` 需要 `config.ParseConfig(path)`
- `SetLogger()` 必须在 `ParseConfig` 之后调用（依赖 `config.C.ENV` 和 `config.C.LogLevel`）
- Logger: `go.uber.org/zap`（全局 `zap.L()`）
