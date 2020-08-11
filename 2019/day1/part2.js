var fs = require('fs')
var os = require('os')

fs.readFile('input.txt', 'utf8', function(err, data) {
    const lines = data.split(os.EOL)

    let sum = 0

    lines.forEach((line) => {
        sum += calculateMass(parseInt(line), 0)
    })

    console.log('The sum is', sum)
})

const calculateMass = (mass, fuelSum) => {
    const fuel =  Math.floor(mass / 3) - 2

    fuelSum += Math.max(fuel, 0)

    if (fuel <= 0) {
        return fuelSum
    }

    return calculateMass(fuel, fuelSum)
}