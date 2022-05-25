## 技術スタック
- フレームワーク: [Goa v3](https://goa.design/)
- O/Rマッパ: [GORM](https://gorm.io/ja_JP/)
- ホットリロード機能: [Air](https://github.com/cosmtrek/air)
- データベース: MySQL
- 環境変数管理: [godotenv](https://github.com/joho/godotenv)

# 環境構築手順
1. goaのファイル自動生成を実行
```bash
goa gen goa-sample/design
goa example goa-sample/design
```

2. dockerの起動
```bash
docker-compose up -d
```

## アクセス
### goa-sample API
http://localhost:8000

### Swagger
http://localhost:8002

### phpmyadmin
http://localhost:9000

### mailcatcher
http://localhost:9000

## APIへの接続について
rest.httpファイルにVSCodeのREST拡張仕様のファイルを用意してます
そちらをお使いください

## マイグレーション
```
docker-compose run todo-migrate up 1
```