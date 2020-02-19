# chapter4 cutコマンドをgolangで実装！
- cutコマンドをgolangで実装し、go-cutというコマンドを実装してみましょう
- cutコマンドのオプションや挙動を全て詳しく実装する必要はありません

## cutコマンドの実装

### 注意点
- sample.csvをgo-cutコマンドで読み込みコンソールに結果を表示させる
- main.goにコードを書いていきましょう
- ビルドはMakefileを使用してビルドすること
    - Makefileに予めgo-cutというバイナリ名でビルドされるように設定してあります
- 区切り文字は`,`とする
    - cutコマンドでは-dオプションで区切り文字を指定することが出来ます 
    - cutコマンドの-dオプションをgolangで実装しましょう
- フィールドは2番目を指定すること
    - cutコマンドでは-fオプションでフィールドの何番目を指定するかを決めることが出来ます 
    - cutコマンドの-fオプションをgolangで実装しましょう
    
### 作ったgo-cutコマンドを実行
- 成功すると下記のような出力になるはずです。GoGo!
```shell script
% ./go-cut -d "," -f 2 sample.csv
GoodAfternoon 
Hi
GoodMorning
Hello
GoodEvening 
```

