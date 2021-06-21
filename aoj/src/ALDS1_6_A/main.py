K = 10_000

def printarray(x):
    print(" ".join(str(v) for v in x))

def countingSort(a, n, k):
    c = [0 for _ in range(0, k)]
    for j in range(0, n):
        c[a[j]] += 1
    for i in range(1, k):
        c[i] += c[i-1]
    
    b = [ 0 for _ in range(0, n)]
    for i in range(n, 0, -1):
        b[c[a[i-1]] - 1] = a[i-1]
        c[a[i-1]] -= 1
    
    return b

def main():
    n = int(input())
    a = [int(v) for v in input().split(" ")]
    b = countingSort(a, n, K)
    printarray(b)
    

if __name__ == '__main__':
    main()
