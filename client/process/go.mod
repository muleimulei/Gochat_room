module process

require message v1.0.0

require util v1.0.0

require model v1.0.0

replace message => ../../common/message

replace util => ../../common/util

replace model => ../model

go 1.17
