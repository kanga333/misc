from sys import stdin

def main():
    n = int(input())
    d = {}
    for _ in range(0, n):
        cmd = stdin.readline().strip().split(" ")
        val = cmd[1]
        if cmd[0] == "insert":
            d[val] = True
        else:
            if d.get(val, False):
                print("yes")
            else:
                print("no")

if __name__ == '__main__':
    main()
