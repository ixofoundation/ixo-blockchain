const sov = require('sovrin-did')
console.log(JSON.stringify(sov.gen()).replace('"did":"', '"did":"did:ixo:'))
