# countname

# 機能

- 20 文字以上の変数、関数、定数名を指摘する

# sample

対象ファイル

```go
package tastcase1

import "log"

var hogehogehogehogehoge, hogehogehogehogehogehoge string // want "name is longer than"

var ( // want "name is longer than"
	fugofugofugofugofugofugo1 = 1
	fugofugofugofugo          = 2
)

const doremifasorashidodoremifa = "doremi" // want "name is longer than"

func hoge() {
	if false {
		var fugafugafugafugafuga int // want "name is longer than"
		log.Println(fugafugafugafugafuga)
	}
}

```

実行結果

```zsh
>>> go vet -vettool=$(which countname) testdata/src/testcase1/testcase1.go
# command-line-arguments
testdata/src/testcase1/testcase1.go:5:1: name is longer than 19
testdata/src/testcase1/testcase1.go:7:1: name is longer than 19
testdata/src/testcase1/testcase1.go:12:1: name is longer than 19
testdata/src/testcase1/testcase1.go:16:3: name is longer than 19
```
