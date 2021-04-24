# ALDS1_3_B: キュー
def main():
    x = [v for v in input().split(" ")]
    n, q = int(x[0]), int(x[1])
    queue = []
    for i in range(0, n):
        x = [v for v in input().split(" ")]
        queue.append({'name': x[0], 'time':int(x[1])})
    
    time = 0
    while len(queue) != 0:
        p = queue.pop(0)
        if p['time'] > q:
            p['time'] -= q
            time += q
            queue.append(p)
        else:
            time += p['time']
            print(p['name'], time)

if __name__ == '__main__':
    main()
