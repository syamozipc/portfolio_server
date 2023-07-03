サンプルアプリ

次やること
- nullableなカラムを作成してnull〇〇を作る
- usersテーブルを作る
- dbインスタンスの使い回しはcontextにする？（apiリクエストごとに別goroutineが起動するので、グローバル変数ではダメ？）
- migrateは別sqlインスタンスで良いのか？
- interface導入
- テスト実装
- error handling
- sqlのコネクションはmiddlewareあたりでやる？（repository単位でやらない）
- 時刻の扱い（いちいちtime.Localを設定しない）
