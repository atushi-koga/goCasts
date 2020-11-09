- パッケージ
パッケージは 'Executable package' と 'Reusable package' の2種類ある。
Executable package = mainパッケージ であり、コンパイル、実行可能なパッケージ。main関数を必ず持たなければならない。
mainパッケージの main関数からプログラムの実行が開始される。
Reusableパッケージは、ヘルパーコードとして利用されるパッケージで、単独でコンパイル、実行不可。

- テストコード
テストファイル名は `_test.go` で終わる名前にする。`deck.go` のテストであれば `deck_test.go` になる。
パッケージ内の全てのテストを実行するには、 `go test` コマンドを実行する。