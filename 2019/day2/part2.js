var fs = require('fs')

fs.readFile('input.txt', 'utf8', function(err, data) {
    let sequence = data.split(',').map(object => parseInt(object))

    let solution = solve(JSON.parse(JSON.stringify(sequence)))

    for (let noun = 0; noun < 100; noun++) {
        for (let verb = 0; verb < 100; verb++) {
            sequence[1] = noun
            sequence[2] = verb

            solution = solve(JSON.parse(JSON.stringify(sequence)))

            if (solution === 19690720) {
                console.log(`Solved! Noun is ${noun} and verb is ${verb}, so the answer is ${100 * noun + verb}`)
                return
            }
        }
    }

})

const solve = (sequence) => {
    let index = 0

    while(index >= 0) {
        index = execute(sequence, index)
    }

    return sequence[0]
}

const execute = (sequence, index) => {
    const calculated = calculateOpcode(sequence[index], sequence[sequence[index + 1]], sequence[sequence[index + 2]])

    if (!!calculated || calculated === 0) {
        sequence[sequence[index + 3]] = calculated
        return index + 4
    }

    return -1
}

const calculateOpcode = (opCode, valueOne, valueTwo) => {
    let result = null;

    if (opCode === 1) {
        result = valueOne + valueTwo
    }

    if (opCode === 2) {
        result = valueOne * valueTwo
    }

    return result
}