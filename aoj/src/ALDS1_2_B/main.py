count = 0

def printarray(x):
    print(" ".join([str(v) for v in x]))

def selectsort(x):
    global count

    for i in range(0, len(x)):
        minj = i
        for j in range(i, len(x)):
            if x[j] < x[minj]:
                minj = j
        if minj != i:
            x[minj], x[i] = x[i], x[minj]
            count+=1

# ALDS1_2_A: 選択ソート
def main():
    global count
    input() # skip scan n
    x = [int(v) for v in input().split(" ")]
    selectsort(x)
    printarray(x)
    print(count)

if __name__ == '__main__':
    main()
