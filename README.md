# ponz_goecho_server

GoのWAF(Web Application Framework)であるEchoを使って簡単なWebサーバーを立てる。

## 使い方

* そのまま動かすパターンと, Dockerコンテナとして動かすパターンどちらも可能。
* Dockerで動かすパターンの使い方だけ書きます。

### Dockerで動かす

* 事前にDockerをインストールしてください。
* Dockerで動かす場合はGoをインストールする必要はない。

```
$ docker pull sasaken555/ponz_goecho_server:1.0.0
$ docker run -p 1323:1323 sasaken555/ponz_goecho_server:1.0.0

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.5
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323

```

## 作成者

* sasaken555
  * [All Repository Here!](https://github.com/sasaken555)
