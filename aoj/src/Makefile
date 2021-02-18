id = GRL_1_C
no = 1
base = https://judgedat.u-aizu.ac.jp/testcases
dir = $(id)/cases/$(no)
in = $(id)/cases/$(no)/input.txt
out = $(id)/cases/$(no)/output.txt

.PHONY: download
download: $(dir) $(in) $(out)

$(dir):
	mkdir -p $(dir)

$(in):
	curl -o $(in) ${base}/$(id)/$(no)/in

$(out):
	curl -o $(out) ${base}/$(id)/$(no)/out

.PHONY: test
test: download
	echo "case $(no)"
	cat $(id)/cases/$(no)/input.txt | go run $(id)/main.go | ~/bin/cliassert -v -stdout-contain "`cat $(id)/cases/$(no)/output.txt`"

.PHONY: open
open:
	cat $(id)/main.go | pbcopy
	open -a Safari http://judge.u-aizu.ac.jp/onlinejudge/submission.jsp#submit/$(id)
