from copy import copy

def printarray(x):
    print(" ".join([str(v) for v in x]))

def bubblesort(x):
    flag = True
    while flag:
        flag = False
        for j in range(len(x)-1, 0, -1):
            if x[j][1] < x[j-1][1]:
                x[j], x[j-1] = x[j-1], x[j]
                flag = True

def selectsort(x):
    for i in range(0, len(x)):
        minj = i
        for j in range(i, len(x)):
            if x[j][1] < x[minj][1]:
                minj = j
        if minj != i:
            x[minj], x[i] = x[i], x[minj]

# ALDS1_2_C: 安定ソート
def main():
    input() # skip scan n
    x = [v for v in input().split(" ")]
    y = copy(x)
    bubblesort(x)
    selectsort(y)
    printarray(x)
    print("Stable")
    printarray(y)
    if x == y:
        print("Stable")
    else:
        print("Not stable")

if __name__ == '__main__':
    main()
