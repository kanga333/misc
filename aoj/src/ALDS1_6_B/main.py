def swap(a, i, j):
    a[i], a[j] = a[j], a[i]

def partition(a, p, r):
    x = a[r]
    i = p
    for j in range(p, r):
        if a[j] <= x:
            swap(a, j, i)
            i += 1
    swap(a, i, r)
    return i

# ALDS1_6_B: パーティション
def main():
    n = int(input())
    a = [int(v) for v in input().split(" ")]
    i = partition(a, 0, n-1)
    print(" ".join([str(v) if j != i else "[" + str(v) + "]" for j, v in enumerate(a) ]))

if __name__ == '__main__':
    main()
