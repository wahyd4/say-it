#!/usr/bin/env node

const chalk = require('chalk')
const program = require('commander')
const request = require('request')
const fs = require('fs')
const player = require('play-sound')(opts = {
  players: ['afplay', 'mplayer', 'play', 'mpg123', 'mpg321', 'aplay', 'cmdmp3']
})
const charm = require('charm')()
const TEXTS = require('./texts')
const _ = require('lodash')

const showError = (error = TEXTS.error) => {
  console.log(chalk.red(error))
}

var person = 0

charm.pipe(process.stdout);

program
  .version('0.0.1')
  .command(' ', 'read the texts you typed in', {
    isDefault: true
  })
  .option('-p, --person [person]', "set different voice. 0: female, 1 and 2: male, 3: male with emotion, 4: female with emotion")

program.on('*', () => {
  const text = program.args[0]
  if (!text || text.trim() === '') {
    return showError()
  }
  if (!_.includes([0, 1, 2, 3, 4], Number.parseInt(program.person))) {
    return showError(TEXTS.invalidPerson)
  }
  person = Number.parseInt(program.person)
  sayIt(text)
})

program.parse(process.argv)

if (!process.argv.slice(2).length) {
  program.outputHelp((text) => {
    console.log(chalk.blue(text))
  })
}

function sayIt(text) {
  const urlAddress = `http://tsn.baidu.com/text2audio?tex=${text}&lan=zh&cuid=${new Date().getTime()}&ctp=1&tok=24.9d61601aef23f1d3497c9c40eb30e7a7.2592000.1499416588.282335-9739014&per=${person}`
  charm.write(chalk.bgCyan(TEXTS.loading))
  request
    .get(encodeURI(urlAddress))
    .on('response', function (response) {
      if ((response.statusCode !== 200) || (response.headers['content-type'] !== 'audio/mp3')) {
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
      charm.left(TEXTS.loading.length)
      charm.destroy()
    })
}
