# ALDS1_3_D: 面積計算
def main():
    x = 0
    slopes = []
    pools = []
    sum = 0
    lines = input()

    for c in lines:
        if c == "\\":
            slopes.append(x)
        if c == "/":
            if len(slopes) == 0:
                x += 1
                continue
            x0 = slopes.pop()
            area = x - x0
            sum += area
            merge_num = 0
            for i in range(len(pools)-1, -1, -1):
                if pools[i]['x0'] < x0:
                    break
                merge_num += 1
            for i in range(0, merge_num):
                area += pools.pop()['area']
            pools.append({'x0': x0, 'area': area})
        x += 1
    
    print(sum)
    if len(pools) == 0:
        print(0)
    else:
        print(len(pools)," ".join([str(p['area']) for p in pools]))

if __name__ == '__main__':
    main()
