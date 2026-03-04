# tool-kit — CLAUDE.md

## Build & Run

```bash
# Build
go build -o bin/tool .

# Run
./bin/tool <command> [subcommand] [flags]

# Tests
go test ./...

# Tidy dependencies
go mod tidy
```

## CLI Commands

| Command | Subcommand | Description |
|---------|-----------|-------------|
| `dts` | — | DB 表 → Go struct |
| `docx` | — | DB 表 → Docx 文档 |
| `str` | `hash -t md5\|sha256\|sha1 <text>` | 哈希计算 |
| `str` | `b64 encode\|decode <text>` | Base64 编解码 |
| `str` | `url encode\|decode <text>` | URL 编解码 |
| `str` | `uuid` | 生成 UUID v4 |
| `str` | `rand [-n 32] [-t alpha\|num\|mix]` | 随机字符串 |
| `str` | `case snake\|camel\|pascal <text>` | 命名风格转换 |
| `json` | `fmt [-f file]` | 格式化 JSON |
| `json` | `mini [-f file]` | 压缩 JSON |
| `json` | `valid [-f file]` | 校验 JSON |
| `ts` | `now` | 当前时间戳 + 多格式 |
| `ts` | `to <timestamp>` | Unix → 可读时间 |
| `ts` | `from <datetime>` | 可读时间 → Unix |

## Directory Structure

```
tool-kit/
├── main.go
├── CLAUDE.md
├── Makefile
├── Dockerfile
├── cmd/
│   ├── root.go             # cli.App, Execute()
│   ├── const.go            # ENV constants
│   ├── setup.go            # SetLogger()
│   ├── table_to_struct.go  # dts command
│   ├── table_to_docx.go    # docx command
│   ├── str.go              # str command
│   ├── json.go             # json command
│   └── ts.go               # ts command
├── config/
│   └── config.go           # Config struct (ENV, Debug, LogLevel, Database)
└── pkg/
    ├── db/                 # MySQL/Postgres client + struct/docx generators
    ├── document/           # docx.go (gooxml wrapper)
    ├── strutil/            # strutil.go (hash, b64, url, uuid, rand, case)
    └── jsonutil/           # jsonutil.go (fmt, mini, valid)
```

## Key Conventions

- CLI framework: `github.com/urfave/cli/v2`
- Each command file registers itself via `func init()` appending to `rootCmd.Commands`
- Config is parsed per-command using `config.ParseConfig(configPath)`; stateless tools (str, json, ts) skip config entirely
- Logger: `go.uber.org/zap` (global logger via `zap.L()`)
- `json` and `ts` commands use stdlib only; `str` uses `github.com/google/uuid`

## Config Format (config.yaml)

```yaml
env: dev          # dev | prod
debug: false
log_level: info   # debug | info | warn | error
database:
  driver: mysql   # mysql | postgres
  host: localhost
  port: 3306
  user: root
  password: ""
  name: mydb
```

## Notes

- `json fmt/mini/valid` accept input from `-f <file>` or stdin (pipe-friendly)
- `ts to` auto-detects second vs millisecond timestamps (> 1e12 → ms)
- `ts from` expects format `2006-01-02 15:04:05` (Go reference time)
- `str rand` defaults to length 32, mix charset
