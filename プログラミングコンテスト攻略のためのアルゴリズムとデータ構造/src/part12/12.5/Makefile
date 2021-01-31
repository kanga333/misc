id = ALDS1_11_D
base = https://judgedat.u-aizu.ac.jp/testcases
no = 1

download:
	mkdir -p cases/$(no)
	curl -o cases/$(no)/input.txt ${base}/$(id)/$(no)/in
	curl -o cases/$(no)/output.txt ${base}/$(id)/$(no)/out

test:
	echo "case $(no)"
	cat cases/$(no)/input.txt | go run main.go | ~/bin/cliassert -v -stdout-contain "`cat cases/$(no)/output.txt`"

open:
	cat main.go | pbcopy
	open http://judge.u-aizu.ac.jp/onlinejudge/submission.jsp#submit/$(id)
