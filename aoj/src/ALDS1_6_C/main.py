from copy import copy

class Card:
    def __init__(self, input):
        suit, number = input.split(" ")
        self.suit = suit
        self.number = int(number)
    
    def __str__(self):
        return str(self.suit) + " " + str(self.number)

def swap(a, i, j):
    a[i], a[j] = a[j], a[i]

def partition(a, p, r):
    x = a[r].number
    i = p
    for j in range(p, r):
        if a[j].number <= x:
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

def merge(a, left, mid, right):
    l = copy(a[left:mid])
    r = copy(a[mid:right])
    l.append(Card("c 1000000000"))
    r.append(Card("c 1000000000"))

    li, ri = 0, 0
    for i in range(left, right):
        if l[li].number <= r[ri].number:
            a[i] = l[li]
            li += 1
        else:
            a[i] = r[ri]
            ri += 1

def merge_sort(a, left, right):
    if left+1 >= right:
        return
    mid = int((left + right) / 2)
    merge_sort(a, left, mid)
    merge_sort(a, mid, right)
    merge(a, left, mid, right)
    return

def main():
    n = int(input())
    a = []
    b = []
    for _ in range(n):
        card = Card(input())
        a.append(card)
        b.append(card)
    quickSort(a, 0, n-1)
    merge_sort(b, 0, n)
    aa = [ str(v) for v in a ]
    ba = [ str(v) for v in b ]
    if aa == ba:
        print("Stable")
    else:
        print("Not stable")
    for c in a:
        print(c)


if __name__ == '__main__':
    main()
