from copy import copy

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

def quickSort(a, p, r):
    if p >= r:
        return
    i = partition(a, p, r)
    quickSort(a, p, i - 1)
    quickSort(a, i + 1, r)

def detect_cycles(a, b):
    b_map = index_map(b)
    cycles = []
    visited = {}
    for v in a:
        cycle = []
        next = v
        while(True):
            if visited.get(next):
                if cycle != []:
                    cycles.append(cycle)
                break
            visited[next] = True
            cycle.append(next)
            next = a[b_map[next]]
    return cycles

def index_map(a):
    map = {}
    for i, v in enumerate(a):
        map[v] = i
    return map

def calc_cost(c, min):
    if len(c) == 1:
        return 0
    c_min = c[0]
    left = c_min * (len(c) -1) + sum(c[1:])
    right = min * (len(c) + 1) + c_min + sum(c)
    if left > right:
        return right
    return left

# ALDS1_6_D:最小コストソート
def main():
    n = int(input())
    a = [int(v) for v in input().split(" ") ]
    b = copy(a)
    
    quickSort(a, 0, n-1)
    min = a[0]
    cycles = detect_cycles(a, b)
    sum = 0
    for c in cycles:
        c = calc_cost(c, min)
        sum += c
    print(sum)

if __name__ == '__main__':
    main()
