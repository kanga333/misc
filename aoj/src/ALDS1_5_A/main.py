def exhaustive_search(target, input):
    if target == 0:
        return True
    if len(input) == 0:
        return False
    if input[0] <= target:
        return exhaustive_search(target - input[0], input[1:]) or exhaustive_search(target, input[1:])
    return exhaustive_search(target, input[1:])
    
    

# ALDS1_5_A: 全探索
def main():
    _ = input()
    s = [int(v) for v in input().split(" ")]
    _ = input()
    q = [int(v) for v in input().split(" ")]
    for qv in q:
        if exhaustive_search(qv, s):
            print("yes")
        else:
            print("no")

if __name__ == '__main__':
    main()
