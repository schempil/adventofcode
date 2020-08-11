var fs = require('fs')
var os = require('os')

fs.readFile('input.txt', 'utf8', function(err, data) {
    const lines = data.split(os.EOL)

    let sum = 0

    lines.forEach((line) => {
        sum += calculateMass(parseInt(line))
    })

    console.log('The sum is', sum)
})

const calculateMass = (mass) => {
    return Math.floor(mass / 3) - 2
}