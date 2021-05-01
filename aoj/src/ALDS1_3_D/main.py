# ALDS1_3_D: 面積計算
def main():
    slopes = []
    pools = []
    sum = 0
    lines = input()

    for x, char in enumerate(lines):
        if char == "\\":
            slopes.append(x)
        if char == "/":
            if len(slopes) == 0:
                continue
            x0 = slopes.pop()
            area = x - x0
            sum += area
            merge_num = 0
            for i in reversed(range(0, len(pools))):
                if pools[i]['x0'] < x0:
                    break
                merge_num += 1
            for _ in range(0, merge_num):
                area += pools.pop()['area']
            pools.append({'x0': x0, 'area': area})
    
    print(sum)
    print(len(pools), end="")
    if len(pools) != 0:
        print(" ", end="")
        print(" ".join([str(p['area']) for p in pools]))

if __name__ == '__main__':
    main()
