from copy import copy

cnt = 0

def printarray(x):
    print(" ".join([str(v) for v in x]))

def calcInterval(n):
    inrerval = []
    v = 1
    while v <= n:
        inrerval.append(v)
        v = 3*v + 1
    inrerval.reverse()
    return inrerval

def insertSort(x, n, g):
    global cnt
    for i in range(g, n):
        v = x[i]
        j = i - g
        while j >= 0 and x[j] > v:
            x[j+g] = x[j]
            j = j -g
            cnt = cnt + 1
        x[j+g] = v


def shellSort(x, n):
    iv = calcInterval(n)
    for v in iv:
        insertSort(x, n, v)
    return iv

# ALDS1_2_D: シェルソート
def main():
    global cnt
    n = int(input())
    x = []
    for i in range(0, n):
        x.append(int(input()))
    iv = shellSort(x, len(x))
    print(len(iv))
    printarray(iv)
    print(cnt)
    for i in x:
        print(i)

if __name__ == '__main__':
    main()
