var fs = require('fs')

fs.readFile('input.txt', 'utf8', function(err, data) {
    let sequence = data.split(',')

    console.log('sequence', sequence)

})