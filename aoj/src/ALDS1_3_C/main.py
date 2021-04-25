# ALDS1_3_C: 連結リスト
def main():
    l = []
    n = int(input())
    for i in range(0, n):
        cmd = input()
        if cmd == "deleteFirst":
            _ = l.pop(0)
            continue
        if cmd == "deleteLast":
            _ = l.pop()
            continue
        splited = cmd.split(" ")
        target = splited[1]
        if splited[0] == "delete":
            for i, v in enumerate(l):
                if target == v:
                    _ = l.pop(i)
                    break
            continue
        # insert
        l.insert(0, target)
    
    print(" ".join(l))

if __name__ == '__main__':
    main()
