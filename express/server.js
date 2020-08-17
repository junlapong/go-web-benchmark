const express = require('express')
const morgan  = require('morgan')

const app  = express()
const port = 8080

app.use(morgan('tiny'))
// app.use(morgan(':date[iso]  :status  :method :url\t:res[content-length] - :response-time ms'))

app.get('/', (req, res) => {
  res.send('Hello, World!')
})

app.get('/ping', (req, res) => {
    res.json({message: 'pong'})
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
