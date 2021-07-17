MAX = 1_001

class LCS():
    def __init__(self):
        self.cache = [ [0 for _ in range(MAX)] for _ in range(MAX)]
  
    def do(self, x, y):
        val = 0
        for xi, xv in enumerate(x):
            for yi, yv in enumerate(y):
                if xv == yv:
                    self.cache[xi + 1][yi + 1] = self.cache[xi][yi] + 1
                elif self.cache[xi][yi + 1] > self.cache[xi + 1][yi]:
                    self.cache[xi + 1][yi + 1] = self.cache[xi][yi + 1]
                else:
                    self.cache[xi + 1][yi + 1] = self.cache[xi + 1][yi]
                if val < self.cache[xi + 1][yi + 1]:
                    val = self.cache[xi + 1][yi + 1]
        return val      

# ALDS1_10_C: LCS
def main():
    n = int(input())
    for _ in range(n):
        x = input()
        y = input()
        print(LCS().do(x, y))
    pass


if __name__ == '__main__':
    main()
