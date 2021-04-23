# ALDS1_3_A: スタック
def main():
    x = [v for v in input().split(" ")]
    stuck = []
    for v in x:
        if v == "+":
            y = stuck.pop()
            x = stuck.pop()
            ret = x + y
            stuck.append(ret)
        elif v == "-":
            y = stuck.pop()
            x = stuck.pop()
            ret = x - y
            stuck.append(ret)
            continue
        elif v == "*":
            y = stuck.pop()
            x = stuck.pop()
            ret = x * y
            stuck.append(ret)
        else:
            stuck.append(int(v))
    print(stuck.pop())

if __name__ == '__main__':
    main()
