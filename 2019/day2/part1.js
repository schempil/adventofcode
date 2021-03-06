var fs = require('fs')

fs.readFile('input.txt', 'utf8', function(err, data) {
    let sequence = data.split(',').map(object => parseInt(object))

    sequence[1] = 12
    sequence[2] = 2

    let index = 0

    while(index >= 0) {
        index = execute(sequence, index)
    }

    console.log('The correct answer is', sequence[0])
})

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