---
name: commit
version: 1.0.0
description: |
  根据 git diff 自动生成符合 Conventional Commits 规范的提交信息，支持中英文。
  Use when asked to "生成提交信息", "commit message", "生成 commit", "帮我提交", or "/commit".
allowed-tools:
  - Bash
  - AskUserQuestion
---

# /commit — 自动生成提交信息

## 步骤 1：收集 diff 信息

```bash
echo "=== Git Status ==="
git status --short
echo ""
echo "=== Staged Changes ==="
git diff --staged --stat
echo ""
echo "=== Staged Diff (前200行) ==="
git diff --staged | head -200
echo ""
echo "=== Recent Commits (风格参考) ==="
git log --oneline -5
```

## 步骤 2：分析变更

根据上面的 diff 输出，判断：

- **变更类型**：参考最近的提交风格决定用中文还是英文
- **影响范围**：哪个模块/文件/功能受影响
- **变更性质**：新功能、修复、重构、文档、测试等

如果没有 staged 的变更（`git diff --staged` 为空），检查是否有未暂存的改动：

```bash
git diff --stat
```

如果有未暂存改动，用 AskUserQuestion 询问用户是否要先 `git add`，或者指定要 add 哪些文件。

## 步骤 3：生成提交信息

根据 [Conventional Commits](https://www.conventionalcommits.org/) 规范生成信息：

```
<type>(<scope>): <subject>

[可选的 body]

[可选的 footer]
```

**type 对照表：**
| type | 适用场景 |
|------|---------|
| `feat` | 新功能 |
| `fix` | Bug 修复 |
| `refactor` | 重构（不影响功能） |
| `docs` | 文档变更 |
| `test` | 测试相关 |
| `chore` | 构建/工具/依赖等杂项 |
| `perf` | 性能优化 |
| `style` | 格式/代码风格（不影响逻辑） |
| `ci` | CI/CD 配置 |
| `build` | 构建系统变更 |

**scope 选取规则（必填）：**
- scope 始终用英文小写
- 从变更文件路径中提取最能代表本次修改范围的模块名
- 优先使用顶层目录名：如 `cmd`、`pkg`、`config`、`.claude` → `claude`
- 如果变更跨多个目录，使用最主要的模块；若无法归类则用项目名
- 示例：
  - 修改 `cmd/str.go` → scope: `str`
  - 修改 `pkg/db/mysql.go` → scope: `db`
  - 修改 `CLAUDE.md` + `.claude/skills/` → scope: `claude`
  - 修改 `Makefile` + `Dockerfile` → scope: `build`

**语言规则：**
- 观察最近5条提交的语言风格
- 如果项目主要用中文提交，subject 用中文
- 如果主要用英文，subject 用英文

## 步骤 4：展示并确认

将生成的提交信息展示给用户，格式如下：

```
生成的提交信息：

  <type>(<scope>): <subject>

是否使用此信息提交？
  [y] 直接提交
  [e] 编辑后提交
  [c] 取消
```

使用 AskUserQuestion 询问用户选择。

## 步骤 5：执行提交

- 如果用户选 **y**：直接运行 `git commit -m "<message>"`
- 如果用户选 **e**：让用户说出修改内容，更新后再确认提交
- 如果用户选 **c**：取消，不做任何操作

提交完成后显示 `git log --oneline -1` 确认。

## 注意事项

- 不要自动 `git push`，除非用户明确要求
- 如果 staged 区域为空，不要自动 `git add .`，先询问用户
- subject 长度建议不超过 72 个字符
- 不要在 commit message 里加 "Co-Authored-By: Claude" 等 AI 署名
