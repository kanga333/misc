for c in `ls cases`
do
echo "case $c"
cat cases/$c/input.txt | go run main.go | ~/bin/cliassert -v -stdout-contain "`cat cases/$c/output.txt`"
done
