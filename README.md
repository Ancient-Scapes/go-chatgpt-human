# ChatGPTで作られたソースコードでChatGPTにリクエストしてChatGPTのレスポンスをもらう全身ChatGPT人間

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
go run main.go 本田圭佑とじゃんけんをして勝つ方法
```

![image](https://user-images.githubusercontent.com/18649842/223357350-e8a50b2e-28a5-4c61-aca6-cbad5792b3a4.png)
