module main

require (
	message v1.0.0
	util v1.0.0
	process v1.0.0
)

replace message => ../../common/message
replace process => ../process
replace util => ../../common/util

go 1.17
