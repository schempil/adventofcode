var fs = require('fs')
var os = require('os')

fs.readFile('input.txt', 'utf8', function(err, data) {
    let inputs = data.split(os.EOL).map(input => input.split(','))

    let coordinates = []

    inputs.forEach((input, i) => {

        let pointer = { x: 0, y: 0 }

        coordinates[i] = [pointer]

        input.forEach(command => {
            console.log(command, pointer)

            const direction = command.substr(0, 1)
            const amount = command.substr(1, command.length)
            const axis = (direction === 'R' || direction === 'L') ? 'x' : 'y'
            const isPositive = direction === 'R' || direction === 'U'

            const startingMover = coordinates[i][coordinates[i].length - 1][axis]

            for (let mover = startingMover; mover < startingMover + amount; mover++) {
                pointer[axis] = isPositive ? startingMover + mover : startingMover - mover
                coordinates[i].push({x: pointer.x, y: pointer.y})
            }

            console.log(coordinates[i])

        })
    })

})