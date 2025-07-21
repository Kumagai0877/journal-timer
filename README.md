# journaltimer

**journaltimer** は、ジャーナリング用のテーマをランダムに選び、一定時間集中して書くための Go 製 CLI タイマーです。  
タイマー終了後には音で通知します。

---

## 🚀 特徴

- テーマは JSON ファイルで管理（自由に追加・編集可能）
- ターミナル上でシンプルに使える
- macOS / Linux 対応の効果音通知付き
- 軽量・クロスプラットフォーム（Go製バイナリ）

---

## 実行方法

ビルド
```
go build -o journaltimer
```

テーマファイルのコピー  
内容は好きなものに変更してください
```
cp themes.json.sample themes.json
```

実行コマンド（タイマーを5分で実行する場合）
```
./journaltimer -t 5
```
