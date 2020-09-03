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
