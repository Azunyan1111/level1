# sample README

## 動作環境

* Unix系OS
* 言語ごとの環境(後述)

## セットアップ

### 1. 環境構築

- Go
  - http://golang-jp.org/doc/install からGolangの1.8以上をインストール
- Java
  - https://java.com/ja/download/ からJava 1.8以上をインストール
  - https://gradle.org/ からGradle4.1以上をインストール
- JavaScript
  - https://nodejs.org/en/ からNode.jsの6.11以上をインストール
- Python3
  - https://www.python.org/about/gettingstarted/ からPythonの3.6以上をインストール
  - https://pip.pypa.io/en/stable/installing/ からpipのインストール
- Ruby
  - https://www.ruby-lang.org/en/documentation/installation/ からRubyの2.4以上をインストール


### 2. 依存パッケージのインストール

```
$ make setup
# 一部必要ないものもあります。
```

## サンプルアプリケーションサーバの実行

```
$ make run-level1
```

Makefileに記載されているポート番号で起動します。

## サンプルリクエストをサーバに投げる

```
$ make curl/sample1
$ make curl/sample2
```
