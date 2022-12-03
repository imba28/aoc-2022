with open("input") as file:
    score = 0
    for line in file:
        m = len(line) // 2
        a, b = line[:m], line[m:]
        shared_item_types = set(a) & set(b)
        
        char = list(shared_item_types)[0]
        char_score = ord(char)
        if char.isupper():
            char_score += -ord('A') + 27
        else:
            char_score += -ord('a') + 1
            
        score += char_score
    print(score)