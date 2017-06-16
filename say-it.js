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
const updateNotifier = require('update-notifier');
const pkg = require('./package.json');

const DEFAULT_PERSON =  0
const DEFAULT_SPEED = 5
const DEFAULT_PITCH = 5

const showError = (error = TEXTS.error) => {
  console.log(chalk.red(error))
}

const checkInputSource = (input, items) => {
  if (!!input && !_.includes(items, Number.parseInt(input))) {
    return false
  }
  return true
}

var person = DEFAULT_PERSON
var speed = DEFAULT_SPEED
var pitch = DEFAULT_PITCH

charm.pipe(process.stdout);

updateNotifier({
  pkg
}).notify(); //show npm package update if available

program
  .version('0.1.6')
  .command(' ', 'read the texts you typed in', {
    isDefault: true
  })
  .option('-p, --person [person]', "set different voice. 0: female, 1 and 2: male, 3: male with emotion, 4: female with emotion", DEFAULT_PERSON)
  .option('-s, --speed [speed]', "set speed of voice. 0 - 9, and default is 5", DEFAULT_SPEED)
  .option('-t --pitch [pitch]', "set the voice pitch. 0 - 9, and default is 5", DEFAULT_PITCH)

program.on('*', () => {
  const text = program.args[0]
  if (!text || text.trim() === '') {
    return showError()
  }
  var personValid = checkInputSource(program.person, [0, 1, 2, 3, 4])
  var speedValid = checkInputSource(program.speed, [0, 1, 2, 3, 4, 5, 6, 7, 8, 9])
  var pitchValid = checkInputSource(program.pitch, [0, 1, 2, 3, 4, 5, 6, 7, 8, 9])

  if (!personValid || !speedValid || !pitchValid) {
    return showError(TEXTS.errorOption)
  }
  person = Number.parseInt(program.person)
  speed = Number.parseInt(program.speed)
  pitch = Number.parseInt(program.pitch)

  sayIt(text)
})

program.parse(process.argv)

if (!process.argv.slice(2).length) {
  program.outputHelp((text) => {
    console.log(chalk.blue(text))
  })
}

function sayIt(text) {
  const urlAddress = `http://tsn.baidu.com/text2audio?tex=${text}&lan=zh&cuid=${new Date().getTime()}&ctp=1&tok=24.9d61601aef23f1d3497c9c40eb30e7a7.2592000.1499416588.282335-9739014&per=${person}&pit=${pitch}&spd=${speed}`
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
