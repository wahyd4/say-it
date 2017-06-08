#!/usr/bin/env node

const chalk = require('chalk')
const program = require('commander')
const request = require('request')
const fs = require('fs')
const player = require('play-sound')(opts = { players: ['afplay', 'mplayer', 'play', 'mpg123', 'mpg321', 'aplay', 'cmdmp3'] })
const charm = require('charm')()
const { TEXT_ERROR, LOADING } = require('./texts')

const showError = (error = TEXT_ERROR) => {
  console.log(chalk.red(error))
}

charm.pipe(process.stdout);

program
  .version('0.0.1')
  .command('', 'read the texts you typed in', { isDefault: true })
  .option('-p, --person [person]', 'choose the voice')

program.on('*', () => {
  const text = process.argv[2]
  if (!text || text.trim() === '') {
    return showError()
  }

  sayIt(text)
})

program.parse(process.argv)

if (!process.argv.slice(2).length) {
  program.outputHelp((text) => {
    console.log(chalk.blue(text))
  })
}

function sayIt(text) {
  const urlAddress = `http://tsn.baidu.com/text2audio?tex=${text}&lan=zh&cuid=${new Date().getTime()}&ctp=1&tok=24.9d61601aef23f1d3497c9c40eb30e7a7.2592000.1499416588.282335-9739014&per=0`
  charm.write(chalk.bgCyan(LOADING))
  request
    .get(encodeURI(urlAddress))
    .on('response', function (response) {
      if ((response.statusCode !== 200) || (response.headers['content-type'] !== 'audio/mp3')) {
        console.log(response.headers['content-type'])
        showError()
        return
      }
    })
    .pipe(fs.createWriteStream('say-it.mp3'))
    .on('finish', () => {
      player.play('say-it.mp3', function (err) {
        if (!!err) {
          showError(err)
        }
      })
      charm.erase('line')
      charm.left(LOADING.length)
      charm.destroy()
    })
}