var fs = require('fs')
var os = require('os')

fs.readFile('input.txt', 'utf8', function(err, data) {
    let sequence = data.split(',').map(object => parseInt(object))

    let index = 0

    while(index >= 0) {
        index = execute(sequence, index)
    }

    console.log(sequence)
})

const execute = (sequence, index) => {

    console.log('### execute index', index)

    const calculated = calculateOpcode(sequence[index], sequence[index + 1], sequence[index + 2])

    console.log('### calculated', calculated)

    if (calculated) {
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