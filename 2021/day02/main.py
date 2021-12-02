file1 = open('input.txt', 'r')
lines = file1.readlines()


class Instruction:
    def __init__(self, direction, amount):
        self.direction = direction
        self.amount = amount


class SubmarinePosition:
    def __init__(self, x, depth):
        self.x = x
        self.depth = depth


def convert_lines_to_instructions(incoming_lines):
    instructions_from_lines = []

    for line in incoming_lines:
        split_result = line.strip().split(" ")
        instructions_from_lines.append(Instruction(split_result[0], int(split_result[1])))

    return instructions_from_lines


def determine_position_after_instructions(incoming_instructions):

    submarine_position = SubmarinePosition(0, 0)

    for instruction in incoming_instructions:
        if instruction.direction == 'forward':
            submarine_position.x += instruction.amount

        if instruction.direction == 'up':
            submarine_position.depth -= instruction.amount

        if instruction.direction == 'down':
            submarine_position.depth += instruction.amount

    return submarine_position


instructions = convert_lines_to_instructions(lines)

final_submarine_position = determine_position_after_instructions(instructions)
result_part_one = final_submarine_position.x * final_submarine_position.depth
print(result_part_one)