package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rivo/tview"
)

// ファイルの内容を読み込み、値を返す関数
func getFile() string {
	data, _ := ioutil.ReadFile("output.txt")
	return string(data)
}

func main() {
	save := 0

	alltodo := getFile()

	app := tview.NewApplication()

	// 「新規 ToDo リスト」のフォーム
	form := tview.NewForm().
		AddInputField("タイトル", "", 20, nil, nil).
		AddInputField("説明", "", 40, nil, nil).
		AddCheckbox("完了", false, nil).
		AddButton("Save", func() {
			save = 1
		}).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("新規 ToDo リスト").SetTitleAlign(tview.AlignCenter)

	// 「追加された ToDo」のテキストビュー
	textView := tview.NewTextView()
	textView.SetText(alltodo)
	textView.SetTitle("追加された ToDo")
	textView.SetBorder(true)

	// form・textViewのそれぞれのウィジェットの配置をflexによって決める
	flex := tview.NewFlex().
		AddItem(form, 0, 1, false).
		AddItem(textView, 0, 2, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	// タイトル・説明を入力して save ボタンを押しときの処理
	if save == 1 && form.GetFormItem(0).(*tview.InputField).GetText() != "" && form.GetFormItem(1).(*tview.InputField).GetText() != "" {
		// ファイルに書き込み
		file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			//エラー処理
			log.Fatal(err)
		}
		fmt.Fprintln(file, "タイトル: "+form.GetFormItem(0).(*tview.InputField).GetText())
		fmt.Fprintln(file, "  説明  : "+form.GetFormItem(1).(*tview.InputField).GetText())
		fmt.Fprintf(file, "\n")

		defer file.Close()
	}
}
