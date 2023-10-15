# Flyioにデプロイする方法
[参考サイト](https://medium.com/data-folks-indonesia/setup-free-postgresql-on-fly-io-and-import-database-3f8f891cbc71)
https://zenn.dev/kaorumori/articles/accfc9bd1ea785

## データベースを用意する方法
1. flyctlをインストールする
2. `flyctl auth login`でログインする
3. `flyctl postgres create`でpostgresqlのプロジェクトを作成する
   1. この時にマシンをdevelopmentを選ばないとお金がかかるので注意
4. `flyctl postgres connect -a <プロジェクト名>`で接続する
5. 4でつないだ状態で`CREATE DATABASE <データベース名>`でデータベースを作成する
6. flyctl proxy 5432 -a <プロジェクト名>でプロキシを立てる
7. `psql -h localhost -U root`でプロキシ経由でデータベースに接続する
   1. この時にdumpファイルとかを使うと一気にデータとかを復元できるらしいが、docker上で動かしているデータベースに対してdumpファイルを作る方法がわからないので、一旦手動でデータを入れることにした
   2. ./init/init.sqlの内容を一つずつコピペで実行する
8. この段階でflyio側にはデータベースとテーブルの用意が出来ているので、データベースのURLを適宜変更しつつsetUp.goのStartSetUp関数を実行する
9. この時点でflyio側にデータベースの用意は終わっているはず

