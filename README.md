# say-it

<!-- TOC -->

- [say-it](#say-it)
  - [Install](#install)
  - [How to use](#how-to-use)
  - [Language support limit](#language-support-limit)
  - [Compatibility](#compatibility)
  - [How to Add cmdmp3 into PATH (Only for windows user)](#how-to-add-cmdmp3-into-path-only-for-windows-user)
  - [Updates](#updates)
  - [License](#license)

<!-- /TOC -->

It just like the `say` command in Mac OSX, but born to be much powerful. We used [Baidu Yuyin](http://yuyin.baidu.com/) as the TTS engine. You need to have network access to use this tool. The pronunciation of Chinese is much better than English. Generally this could be a tool to help you to learn Chinese or English.

## Install
1. The go way

```bash
    go get -u github.com/wahyd4/say-it
```
2. Or the standalone way, please go to releases page to download the binary application then execute it

## How to use

  `say-it "Hello 世界"`

  `say-it -p 3 "春晓 孟浩然  春眠不觉晓，处处闻啼鸟。夜来风雨声，花落知多少。"`

  `say-it -p 4 "Life is like riding a bicycle. To keep your balance, you must keep moving. ― Albert Einstein"`

  `say-it --help`


## Language support limit
  Due to Baidu Yuyin's limit, right now we only support Chinese and English.

## Compatibility

  This is package is developed and tested against under `macOS Sierra` and `Windows 10`, For windows user please make sure download `cmdmp3` and add it to PATH environment.
  For macOS user we used `afplay` which is already installed by default, so you don't need to anything.
  Currently we don't have plan to support `Linux`


  * [`afplay`](https://developer.apple.com/library/mac/documentation/Darwin/Reference/ManPages/man1/afplay.1.html)
  * [`cmdmp3`](https://github.com/jimlawless/cmdmp3)

## How to Add cmdmp3 into PATH (Only for windows user)
1. Download the binary from: <http://www.mailsend-online.com/wp/cmdmp3new.zip>
2. Extract it to a folder. e.g. `c:/cmdmp3`
3. Go to the Advanced system setting page, and modify the `PATH` environment variable by adding `;c:/cmdmp3` to the end.
4. Restart windows and go to command line and test `cmdmp3`, if there is some outputs contains `cmdmp3 v2.0`. Congratualtions, all done!
5. If you install `say-it` by downloading the binary, then you can make the `say-it` to be globally executable just do some works like this.

## Updates

  * 0.1.0 Add basic function to read texts from command line.
  * 0.1.3 Add ability to allow user set different voice.
  * 0.1.7 Allow user to set voice speed and pitch.
  * 0.2.0 Use go to re-implement the features.
  * 0.3.0 Add support for Windows users.
## License

  MIT


