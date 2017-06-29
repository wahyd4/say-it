# Notice

    say-it has been moved to Golang platform. Please install it by `go get github.com/wahyd4/say-it`

    For more details, please check at <https://github.com/wahyd4/say-it>

# say-it

[![NPM](https://nodei.co/npm/say-it.png)](https://npmjs.org/package/say-it)

It just like the `say` command in Mac OSX, but born to be much powerful. We used [Baidu Yuyin](http://yuyin.baidu.com/) as the TTS engine. You need to have network access to use this tool. The pronunciation of Chinese is much better than English. Generally this could be a tool to help you to learn Chinese or English.

## Install

    npm install -g say-it

## How to use

  `say-it 'Hello 世界'`

  `say-it -p 3 '春晓 孟浩然  春眠不觉晓，处处闻啼鸟。夜来风雨声，花落知多少。'`

  `say-it -p 4 "Life is like riding a bicycle. To keep your balance, you must keep moving. ― Albert Einstein"`


## Language support limit
  Due to Baidu Yuyin's limit, right now we only support Chinese and English.

## Compatibility

  This is package is developed and tested against under `macOS Sierra`, but it should be worked on all Mac platforms if you can install the latest stable version of Node. For `Linux` and `Windows` users, please make sure you have at least one of the players installed from the list.

  * [`mplayer`](https://www.mplayerhq.hu/)
  * [`afplay`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/afplay.1.html)
  * [`mpg123`](http://www.mpg123.de/)
  * [`mpg321`](http://mpg321.sourceforge.net/)
  * [`play`](http://sox.sourceforge.net/)
  * [`omxplayer`](https://github.com/popcornmix/omxplayer)
  * [`aplay`](https://linux.die.net/man/1/aplay)
  * [`cmdmp3`](https://github.com/jimlawless/cmdmp3)


## Updates

  * 0.1.0 Add basic function to read texts from command line.
  * 0.1.3 Add ability to allow user set different voice.
  * 0.1.7 Allow user to set voice speed and pitch.
## License

  MIT


