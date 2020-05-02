
def solution(line, n, k):
    Max, curr = 0, 0
    m_k = k
    for i in range(n):
        if line[i] == "0":
            if m_k > 0:
                m_k -= 1
                curr += 1
            else:
                m_k = k
                curr = 0
        else:
            curr += 1
        Max = max(Max, curr)

    return Max


_input = input().split(" ")
n, k = int(_input[0]), int(_input[1])
line = input().split(" ")
# line = ["0" for i in range(n)]
m = solution(line, n, k)
print(m)
