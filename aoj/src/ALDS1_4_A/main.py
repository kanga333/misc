# ALDS1_4_A: 線形探索
def main():
    _ = input()
    s = input().split(" ")
    _ = input()
    q = input().split(" ")

    cnt = 0
    for qv in q:
        for sv in s:
            if qv == sv:
                cnt += 1
                break
    print(cnt)

if __name__ == '__main__':
    main()
