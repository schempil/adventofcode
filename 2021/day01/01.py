file1 = open('input.txt', 'r')
lines = file1.readlines()

def convertLinesToDepths(lines):

    depths = []

    for line in lines:
        strippedLine = line.strip()
        depths.append(int(strippedLine))

    return depths

def countIncreasedDepths(depths):

    countIncreased = 0

    i = 0
    while i < len(depths):

        if i == 0:
            i += 1
            continue

        if depths[i-1] < depths[i]:
            countIncreased += 1

        i += 1

    return countIncreased

depths = convertLinesToDepths(lines)
countIncreased = countIncreasedDepths(depths)

print(countIncreased)
