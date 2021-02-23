def printarray(x):
    print(" ".join([str(v) for v in x]))

def insertionsort(x):
    printarray(x)
    for i in range(1, len(x)):
        for j in range(i, 0, -1):
            if x[j] > x[j-1]:
                break
            x[j], x[j-1] = x[j-1], x[j]
        printarray(x)

# ALDS1_1_A: 挿入ソート
def main():
    input() # skip scan n
    x = [int(v) for v in input().split(" ")]
    insertionsort(x)

if __name__ == '__main__':
    main()
