from collections import deque

# ALDS1_3_C: 連結リスト
def main():
    l = deque()
    n = int(input())
    for i in range(0, n):
        cmd = input()
        if cmd == "deleteFirst":
            _ = l.popleft()
            continue
        if cmd == "deleteLast":
            _ = l.pop()
            continue
        splited = cmd.split(" ")
        target = splited[1]
        if splited[0] == "delete":
            try:
                l.remove(target)
            except ValueError:
                pass
            continue
        # insert
        l.appendleft(target)
    
    print(" ".join(list(l)))

if __name__ == '__main__':
    main()
