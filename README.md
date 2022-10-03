# ✏IO-Stack

これはシェルの標準入出力をスタックとして保存・出力できるコマンドラインツールです。     
シェル芸等にお使いください。

## 🔨Examples
```bash
>echo A|IO-Stack
>IO-Stack
A
```
Aをスタックに積み、積んだAを出力しています。

```bash
>echo A|IO-Stack
>echo B|IO-Stack
>IO-Stack
B
>IO-Stack
A
```
A,Bの順にスタックを積み、積んだ順から順番に出力しています。


```bash
>echo A|IO-Stack
>echo B|IO-Stack
>IO-Stack --nd
B
>IO-Stack --nd
B
```
デフォルト設定だと、取り出したスタックは消去されます。取り出す際にndオプションをつけると消去されません。

```bash
>echo A|IO-Stack
>echo B|IO-Stack
>IO-Stack --nd 0
A
>IO-Stack --nd 1
B
```
コマンドライン引数に数字を入れると指定インデックスのスタックを取り出すことができます。


```bash
>echo A|IO-Stack
>echo B|IO-Stack
>IO-Stack
B
>IO-Stack
A
>IO-Stack
```
スタックが空の場合何も出力されません。

## 💪Examples(将来)
将来的にできるようになる機能です。    
~~Done is better than perfect.~~    
近いうちに対応します。    
```bash
>echo A|IO-Stack
>echo B|IO-Stack
>IO-Stack --nd -1
B
>IO-Stack --nd -2
A
```
マイナスのインデックスによるスタックの取り出し。
```bash
>echo A|IO-Stack
>echo B|IO-Stack 0
>IO-Stack --nd 0
B
>IO-Stack --nd 1
A
```
インデックスを指定したデータの途中挿入。
```bash
>echo A|IO-Stack
>echo B|IO-Stack 0
>IO-Stack --flush
A
B
>IO-Stack
```
スタックをすべて出力。   


## 🎫LICENSE

[MIT](./LICENSE)

## ✍Author

[PenguinCabinet](https://github.com/PenguinCabinet)


