---
name: merge-feature
version: 1.0.0
description: 合并 worktree 分支到主分支，验证后清理
disable-model-invocation: true
---

合并功能分支：$ARGUMENTS

如果 $ARGUMENTS 为空，运行 `git worktree list` 列出所有 worktree 让用户选择。

步骤：
1. 确保在项目根目录：`cd` 到主 worktree
2. 确保主分支干净：`git status` 检查无未提交的变更
3. 合并分支：`git merge $ARGUMENTS`
4. 验证：`go build ./... && go test ./...`
5. 如果验证失败：`git merge --abort`，告诉用户失败原因，停止
6. 如果验证通过：清理 worktree 和分支
   ```bash
   git worktree remove .trees/$ARGUMENTS
   git branch -d $ARGUMENTS
   ```
7. 报告合并结果