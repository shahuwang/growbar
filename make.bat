rm main.exe
go tool yacc -o parser.go -p Grow growbar.y
go build main.go parser.go lex.go struct.go dev.go interpreter.go error.go create.go eval.go utils.go string_pool.go native.go
main.exe
