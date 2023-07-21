# 簡易的な memo コマンド

## 説明

ターミナル開いている時、どうしても即座に何かをメモしたい衝動に駆られたら使えるコマンドです

## セットアップ

1. git clone する

2. memo を一時的に保存するためのファイルのパスを環境変数に登録します

`export MEMO_COMMAND_STORE_PATH=[your-path.bin]`

3. .zprofile などに以下のような記述を書く

```
memo() {
    go run [クローンしたmain.goのパス] $1 $2
}
```

4. source コマンドなどする

## 使い方

- memo -w [メモしたいこと]

メモを書くためのコマンド

ex) memo -w メモメモ

- memo -r

メモした内容を出力するコマンド

- memo -d

メモした内容を全て削除するコマンド
