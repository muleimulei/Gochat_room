module main

require util v1.0.0

require (
	github.com/gomodule/redigo v1.8.8
	message v1.0.0
	process v1.0.0
	model v1.0.0
)

replace message => ../../common/message

replace util => ../../common/util

replace process => ../process
replace model => ../model

go 1.17
