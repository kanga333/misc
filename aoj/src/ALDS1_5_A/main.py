ans = {}

def exhaustive_search(current, input):
    global ans
    if len(input) == 0:
        return False
    next = current + input[0]
    ans[next] = True
    exhaustive_search(next, input[1:])
    exhaustive_search(current, input[1:])
    
    

# ALDS1_5_A: 全探索
def main():
    _ = input()
    s = [int(v) for v in input().split(" ")]
    _ = input()
    q = [int(v) for v in input().split(" ")]
    global ans
    ans[0] = True
    exhaustive_search(0, s)
    for qv in q:
        if qv in ans:
            print("yes")
        else:
            print("no")

if __name__ == '__main__':
    main()
