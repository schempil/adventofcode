file1 = open('input.txt', 'r')
lines = file1.readlines()


def convert_lines_to_depths(incoming_lines):
    integers = []

    for line in incoming_lines:
        stripped_line = line.strip()
        integers.append(int(stripped_line))

    return integers


def count_increased_depths(incoming_depths):
    count = 0

    i = 0
    while i < len(incoming_depths):

        if i == 0:
            i += 1
            continue

        if incoming_depths[i - 1] < incoming_depths[i]:
            count += 1

        i += 1

    return count


def map_measurements_to_three_measured_sums(incoming_depths):
    sums = []

    i = 0
    while i + 2 < len(incoming_depths):
        sums.append(incoming_depths[i] + incoming_depths[i + 1] + incoming_depths[i + 2])

        i += 1

    return sums


if __name__ == "__main__":
    depths = convert_lines_to_depths(lines)
    countIncreased = count_increased_depths(depths)

    print("Day 01, Part 1:", countIncreased)

    threeMeasureSums = map_measurements_to_three_measured_sums(depths)
    countIncreasedFromThreeMeasures = count_increased_depths(threeMeasureSums)

    print("Day 01, Part 2:", countIncreasedFromThreeMeasures)
