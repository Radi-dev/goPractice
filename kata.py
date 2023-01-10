from math import isqrt,sqrt

def buddy(start, limit):
    for n in range(start,limit+1):
        m = sum_divisors(n)-n-1
        if sum_divisors(m)-m==n+1 and m>n:
            return [n,m]
    return "Nothing"

def sum_divisors(n):
    result = 1
    div = 1
    while True:
        n=abs(n)
        for div in range(div + 1, isqrt(n) + 1):
            if not n % div:
                mul = 1
                while not n % div:
                    n //= div
                    mul = 1 + mul * div
                result *= mul
                break
        else:
            if n > 1:
                result *= 1 + n
            return result

# print(sum_divisors(10))
print(buddy(1071625, 1103735))