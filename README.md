# example-go-cov

Go言語のコードカバレッジ計測サンプルです。

## カバレッジの取得方法

```bash
# テスト実行とカバレッジプロファイル生成
go test -v -coverprofile=coverage.out ./...

# カバレッジパーセンテージの確認
go tool cover -func=coverage.out

# HTML形式のレポート生成
go tool cover -html=coverage.out -o coverage.html
```

## GitHub Actions連携

`.github/workflows/coverage.yml` で以下を自動化:
- テスト実行とカバレッジ計測
- カバレッジ率のレポート生成

## ライセンス

MIT
