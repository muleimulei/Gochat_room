module process
require message v1.0.0 // indirect
require util v1.0.0
require model v1.0.0

replace util => ../../common/util
replace message => ../../common/message
replace model =>../model

go 1.17
