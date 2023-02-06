# 立ち上げ方法

## dockerの場合
### ビルド
```
docker-compose build
```

### デプロイ
```
docker-compose up -d
```

## ローカルの場合(Homedrewは予めインストールしておく事)
docker内のGOのバージョンをインストール(1時間位かかる)
```
brew install go@1.18
```

PATHを指定する。
```
echo 'export PATH="/usr/local/opt/go@1.18/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## 参考
定番中の定番:A Tour of Go
https://go-tour-jp.appspot.com/

dockerでGoを始める方法（DBはpostgresqlを使用）
https://zenn.dev/tatsurom/articles/golang-docker-environment
