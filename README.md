# ChatGPTで作られたソースコードでChatGPTにリクエストしてChatGPTのレスポンスをもらう全身ChatGPT人間


<img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px"><img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px"><img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px"><img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px">

## 環境設定

### ~/.envrc

OpenAI公式サイトに登録し、APIキーを発行する必要があります。Googleアカウントと電話番号があれば登録可能です。
その後、以下のURLにアクセスしAPIキーを発行し.envrc配下に入力してください。
https://platform.openai.com/account/api-keys

```
export API_KEY="xxxx"
```

## Usage

```
go mod tidy
```

```
go run main.go ハンバーグを作るのに必要な材料を教えて
```
![image](https://user-images.githubusercontent.com/18649842/223382653-699f546d-853e-45a8-a5ab-717d2ffa1f58.png)


<img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px"><img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px"><img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px"><img src="https://user-images.githubusercontent.com/18649842/223380302-fba6fb13-13e5-437f-80fc-cb34f2e231ea.png" width="200px">
