.PHONY: run test

run:
	go run fizzbuzz.go

test:
	bash -c "diff <(go run fizzbuzz.go) results1-30.txt && echo ok"

