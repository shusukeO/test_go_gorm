## Go の ORM 、GORM を試す

### 参考

- https://gorm.io/ja_JP/docs/index.html

### メモ

- context を渡して時間がかかるクエリのタイムアウト時間を設定できる
- WithContext って記述をよく見て、なんだろうって思ってたやつ

```
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

db.WithContext(ctx).Find(&users)
```

- エラーハンドリングは、\*gorm.DB の Error フィールドを確認する

```
if err := db.Where("name = ?", "jinzhu").First(&user).Error; err != nil {
  // ここでエラーハンドリング
}
// または
if result := db.Where("name = ?", "jinzhu").First(&user); result.Error != nil {
  // ここでエラーハンドリング
}
```
