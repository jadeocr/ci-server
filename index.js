const express = require('express')
const morgan = require('morgan')
const { exec } = require('child_process')
const bodyParser = require('body-parser');

const app = express()
const port = process.env.PORT || 4000
const token = process.env.TOKEN

if (!token) {
  process.exit(1)
}

app.use(morgan(':method :url :status :response-time ms - :res[content-length] [:date]'))
app.use(bodyParser.urlencoded({ extended: true }))

app.get('/', (req, res, next) => {
  res.send('Hello from the <a href=\'https://next.jadeocr.com\'>jadeocr</a> CI Server')
})

app.post('/github', (req, res) => {
  if (req.body.payload == token) {
    res.sendStatus(200)
    exec('sh deploy.sh', (err, stdout, stderr) => {
      let dateTime = new Date().toLocaleString()
      if (err) {
        console.log('Error', err, dateTime)
      } else {
        console.log('Successfully deployed', dateTime)
      }
    })
  } else {
    res.status(400).send('Unauthorized token')
  }
})

app.listen(port, () => {
  console.log(`CI server listening at http://localhost:${port}`)
})

