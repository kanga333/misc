def binary_search(array, val):
    max = len(array)
    min = 0
    while min <= max:
        i = min + int((max - min) / 2)
        if array[i] == val:
            return True
        elif array[i] < val:
            min = i + 1
        else:
            max = i - 1
    return False

def main():
    _ = input()
    s = [int(v) for v in input().split(" ")]
    _ = input()
    t = [int(v) for v in input().split(" ")]

    count = 0
    for tv in t:
        if binary_search(s, tv):
            count += 1
    print(count)

if __name__ == '__main__':
    main()
