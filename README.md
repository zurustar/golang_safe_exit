#

golang で Web アプリを安全に終了する方法を探ります。

- リクエストを受け付けたら DB にクエリを発行して応答を返すようにしたい。
- Ctrl+C を受け付けたら DB へのクエリが終了（走り切る or キャンセルする)するのを待ってから DB への接続を閉じてちゃんと終了したい。