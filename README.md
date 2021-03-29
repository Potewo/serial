# serial
## English
Wrapped [tarm/serial](https://github.com/tarm/serial) to make it easy and safe to communicate with Arduino by golang.  

It makes it easy to send a byte of data using this program because this program sends status before sending data.
Arduino sketch is [here](https://github.com/Potewo/serial-receive-write-with-go-test).  
Connect Arduino to your PC and upload this sketch and then run this go program.  

## 日本語(Japanese)
[tarm/serial](https://github.com/tarm/serial)をラップして安全に、簡単にArduinoとシリアル通信をできるようにしました。  

データを送る前後にステータスなどの情報を送ることで、1バイトごとのバイナリデータを送るのが安全で簡単になります。  

[この](https://github.com/Potewo/serial-receive-write-with-go-test)Arduinoスケッチでこのライブラリを使ったやり取りのテストを行うことができます。  
Arduinoをパソコンに接続し、このスケッチをアップロードしてからこのGoプログラムを動かしてください。
