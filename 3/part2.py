
def group_by(file, n):
    g = list()
    i = 0
    for line in file:
        g.append(line.rstrip())
        if i == n - 1:
            yield g
            g = list()
        i = (i + 1) % n
            
with open("input") as file:
    score = 0
    for groups in group_by(file, 3):
        [a, b, c] = [set(g) for g in groups]
        char = list(a & b & c)[0]
        
        char_score = ord(char)
        if char.isupper():
            char_score += -ord('A') + 27
        else:
            char_score += -ord('a') + 1
            
        score += char_score
    print(score)