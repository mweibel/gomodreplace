module github.com/mweibel/gomodreplace/moduletwo

go 1.21.6

require (
	github.com/mweibel/gomodreplace/lib v0.0.0
)

replace github.com/mweibel/gomodreplace/lib v0.0.0 => ../lib
