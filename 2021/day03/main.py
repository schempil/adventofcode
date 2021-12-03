file = open('input.txt', 'r')
lines = file.readlines()


def get_gamma_rate_from_diagnostic_report(incoming_diagnostic_report):
    calculated_gamma_rate = ""

    for i in range(0, len(incoming_diagnostic_report[0])):
        count_zeros = 0
        count_ones = 0

        for diagnose in incoming_diagnostic_report:
            relevant_bit = diagnose[i]

            if relevant_bit == '1':
                count_ones += 1

            if relevant_bit == '0':
                count_zeros += 1

        if count_ones > count_zeros:
            calculated_gamma_rate += "1"
        else:
            calculated_gamma_rate += "0"

    return calculated_gamma_rate


def get_epsilon_rate_from_gamma_rate(incoming_gamma_rate):
    calculated_epsilon_rate = ""

    for char in incoming_gamma_rate:
        calculated_epsilon_rate += str(1 - int(char))

    return calculated_epsilon_rate


def determine_rating(incoming_diagnostic_report, rating):
    updated_list = incoming_diagnostic_report

    for i in range(0, len(incoming_diagnostic_report[0])):
        count_zeros = 0
        count_ones = 0

        for diagnose in updated_list:
            relevant_bit = diagnose[i]

            if relevant_bit == '1':
                count_ones += 1

            if relevant_bit == '0':
                count_zeros += 1

        if rating == 'oxygen':
            if count_ones >= count_zeros:
                updated_list = list(filter(lambda diagnose_entry: diagnose_entry[i] == '1', updated_list))

            else:
                updated_list = list(filter(lambda diagnose_entry: diagnose_entry[i] == '0', updated_list))
        elif rating == 'co2':
            if count_zeros > count_ones:
                updated_list = list(filter(lambda diagnose_entry: diagnose_entry[i] == '1', updated_list))

            else:
                updated_list = list(filter(lambda diagnose_entry: diagnose_entry[i] == '0', updated_list))

        if len(updated_list) == 1:
            return updated_list[0]


def parse_lines_to_diagnostic_report(incoming_lines):
    parsed_diagnostic_report = []

    for line in incoming_lines:
        parsed_diagnostic_report.append(line.strip())

    return parsed_diagnostic_report


diagnostic_report = parse_lines_to_diagnostic_report(lines)

gamma_rate = get_gamma_rate_from_diagnostic_report(diagnostic_report)
epsilon_rate = get_epsilon_rate_from_gamma_rate(gamma_rate)

result_part_one = int(gamma_rate, 2) * int(epsilon_rate, 2)
print("Day03 Part1", result_part_one)

oxygen_generator_rating = determine_rating(diagnostic_report, 'oxygen')
co2_scrubber_rating = determine_rating(diagnostic_report, 'co2')
life_support_rating = int(oxygen_generator_rating, 2) * int(co2_scrubber_rating, 2)

print("Day03 Part2", life_support_rating)
