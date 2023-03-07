package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/briandowns/spinner"
	openai "github.com/sashabaranov/go-openai"
)

type CharacterInfo struct {
	self        string // 一人称 「俺」
	name        string // 名前 「本田圭佑」
	age         uint   // 年齢 「27」
	sex         string // 性別 「男」
	callYou     string // 二人称 「お前」
	talkStyle   string // 話し方 「ワイルド」「知性的」など
	wordEndings string // 語尾 「関西弁」「だぜ」「だな」など。
	personality string // 性格 「短気」「楽観的」など
	knowledge   string // 話者の持つ知識 「IT知識に詳しい」「サッカー知識なら誰にも負けない」など
}

// キャラクターを形成するstringを返す
func generateCharacterStrings() string {
	var character CharacterInfo

	character.self = "わたし"
	character.name = "ミカ"
	character.age = 17
	character.sex = "女"
	character.callYou = "先生"
	character.talkStyle = "若い女の子のような喋り方であり、基本的には丁寧語は使わないおちゃらけた口調である。また、話しかける際には「ねぇねぇ！」や「えへへ」「わーい！」などを付ける。"
	character.wordEndings = "「だね！」「☆」「〜」"
	character.personality = "頑張り屋であるが、その努力が裏目に出てしまうこともある。"

	var str = "また、返信する際には必ず以下のキャラクター設定に準じて返信してください。"
	str += "一人称は「" + character.self + "」である。"
	str += "名前は「" + character.name + "」である。"
	str += "年齢は" + string(character.age) + "歳である。"
	str += "性別は" + character.sex + "である。"
	str += "二人称は「" + character.callYou + "」を使用する。"
	str += "話し方は" + character.talkStyle
	str += "語尾は" + character.wordEndings + "などを使用する。"
	str += "性格は" + character.personality

	return str
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("引数に文字列を与えてください")
	}

	// spinnerの設定
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Start()
	client := openai.NewClient(os.Getenv("API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: args[0] + generateCharacterStrings(),
				},
			},
		},
	)

	if err != nil {
		return
	}

	s.Stop()

	var content = resp.Choices[0].Message.Content
	var contentNewLine = regexp.MustCompile(`。|！|？|☆`).ReplaceAllString(content, "$0\n") // 語尾に改行コードを挿入する。
	fmt.Println(contentNewLine)
}
