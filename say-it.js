const chalk = require('chalk')
const program = require('commander')

program
  .version('0.0.1')
  .option('-p, --person [person]', 'choose the voice')

program.parse(process.argv)

console.log(chalk.blue('Hello world!'))