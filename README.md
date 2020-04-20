# kohinataBot
小日向美穂ちゃんのDiscordBotリポジトリ
  
## 開発へ参加する方法
1. GitHubアカウントを作成してOrganizationへの招待を依頼してください。(個人リポジトリへのフォークでも可能ですがOrganizationに入ってた方が楽です。)
2. `git clone git@github.com:MishiroProduction/kohinataBot.git` を実行してリポジトリをコンピュータのローカルへクローンします。
3. `git checkout -b {YOUR_NAME}/develop` を実行します( `{YOUR_NAME}` の箇所にはあなたのハンドルネーム等を入れてください。例えば私であれば `git checkout -b kilocalorie/develop` のようになります)。
4. コードをよく読み、編集をしてください。コンテナー化されているため、Dockerがインストールされていれば `.env` ファイルへキーなどを書くことで簡単にローカルでテストをすることも可能です。コマンドは `docker-compose up --build` です。
5. Git操作については後述します。完成したものをリモート(GitHub)の個人のブランチへPushし終えたあと、GitHubのページ上にある [Compare & pull request] ボタンを押し、修正内容についてのコメントを記述後 [Create pull request] ボタンを押します。
6. プルリクは気がつき次第レビューとマージをこちらで実施します。マージ後自動でBotでその機能が利用できるようになります！
  
## 簡単なGit操作(WIP)
1. あとで追加します。
  
## License
The MIT License.