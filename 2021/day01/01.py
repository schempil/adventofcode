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

def mapMeasurementsToThreeMeasuredSums(depths):
    threeMeasureSums = []

    i = 0
    while i+2 < len(depths):

        threeMeasureSums.append(depths[i] + depths[i+1] + depths[i+2])

        i += 1

    return threeMeasureSums

depths = convertLinesToDepths(lines)
countIncreased = countIncreasedDepths(depths)

print("Day01, Part1:", countIncreased)

threeMeasureSums = mapMeasurementsToThreeMeasuredSums(depths)
countIncreasedFromThreeMeasures = countIncreasedDepths(threeMeasureSums)

print("Day01, Part2:", countIncreasedFromThreeMeasures)
