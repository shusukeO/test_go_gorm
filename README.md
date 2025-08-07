# GORM Test Project

SQLite と MySQL で GORM の動作を比較検証するプロジェクトです。

## セットアップ

### MySQL サーバーの起動

```bash
docker-compose up -d
```

### MySQL サーバーの停止

```bash
docker-compose down
```

## 使用方法

```bash
# MySQLで実行（MySQLが起動していない場合は自動で起動）
make run

# SQLiteで実行
make run DB_TYPE=sqlite

# MySQLサーバーを停止（データは保持）
make stop

# MySQLサーバーとデータを完全削除
make clean

# ヘルプを表示
make help
```

## データベース接続情報

### MySQL

- ホスト: localhost:3306
- データベース: testdb
- ユーザー: testuser
- パスワード: testpass

### SQLite

- ファイル: test.db

## 検証ポイント

1. `RegionID` フィールドの `default:'japan'` タグの動作
2. `AutoMigrate` でカラム追加時の既存レコードへのデフォルト値の反映
3. SQLite と MySQL での挙動の違い

## 参考

- https://gorm.io/ja_JP/docs/index.html

### メモ

- context を渡して時間がかかるクエリのタイムアウト時間を設定できる
- WithContext って記述をよく見て、なんだろうって思ってたやつ

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

db.WithContext(ctx).Find(&users)
```

- エラーハンドリングは、\*gorm.DB の Error フィールドを確認する

```go
if err := db.Where("name = ?", "jinzhu").First(&user).Error; err != nil {
  // ここでエラーハンドリング
}
// または
if result := db.Where("name = ?", "jinzhu").First(&user); result.Error != nil {
  // ここでエラーハンドリング
}
```
