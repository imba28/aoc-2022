def contains(min_a, max_a, min_b, max_b):
    return min_a <= min_b and max_a >= min_b and max_a >= max_b

def fully_contains(min_a, max_a, min_b, max_b):
    return contains(min_a, max_a, min_b, max_b) or contains(min_b, max_b, min_a, max_a)

with open("input") as file:
    c = 0
    for line in file:
        [[min_a, max_a], [min_b, max_b]] = [pair.split("-") for pair in line.rstrip().split(",")]
        
        if fully_contains(int(min_a), int(max_a), int(min_b), int(max_b)):
            c += 1
    print(c)