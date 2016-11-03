go tool yacc -o parser.go -p Grow growbar.y
go build main.go parser.go lex.go struct.go dev.go interpreter.go error.go create.go
main.exe
