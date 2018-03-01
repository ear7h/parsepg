run: gen
	go run *.go

gen: pg.peg
	go generate


