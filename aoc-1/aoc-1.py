# Advent of Code Day 1
if __name__ == "__main__":

    sequence_sum = 0
    left_seqs = []
    right_seqs = []
    counts = {}
    with open("input.txt", "r") as f:
        for line in f:
            data = line.strip("\n").split("   ")    
            left, right = data
            left, right = int(left), int(right)
            counts[right] = counts.get(right, 0) + 1
            left_seqs.append(left)
            right_seqs.append(right)

    left_seqs = sorted(left_seqs)
    right_seqs = sorted(right_seqs)

    similarity = 0
    for d1, d2 in zip(left_seqs, right_seqs):
        sequence_sum += abs(d2 - d1)
        similarity += d1 * counts.get(d1, 0)

    print("Sum of differences: ", sequence_sum)
    print("Similarity: ", similarity)