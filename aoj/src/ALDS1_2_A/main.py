count = 0

def printarray(x):
    print(" ".join([str(v) for v in x]))

def bubblesort(x):
    global count
    flag = True

    while flag:
        flag = False
        for j in range(len(x)-1, 0, -1):
            if x[j] < x[j-1]:
                x[j], x[j-1] = x[j-1], x[j]
                flag = True
                count+=1

# ALDS1_2_A: バブルソート
def main():
    global count
    input() # skip scan n
    x = [int(v) for v in input().split(" ")]
    bubblesort(x)
    printarray(x)
    print(count)

if __name__ == '__main__':
    main()
