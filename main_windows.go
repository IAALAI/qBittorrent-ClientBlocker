//go:build windows
package main

import (
	"github.com/lxn/win"
	"golang.design/x/hotkey"
	"github.com/getlantern/systray"
)

var showWindow = true
var qBCBHotkey = hotkey.New([]hotkey.Modifier { hotkey.ModCtrl, hotkey.ModAlt }, hotkey.KeyB)

func ShowOrHiddenWindow() {
	consoleWindow := win.GetConsoleWindow()
	if showWindow {
		Log("Debug-ShowOrHiddenWindow", GetLangText("Debug-ShowOrHiddenWindow_HideWindow"), false)
		showWindow = false
		win.ShowWindow(consoleWindow, win.SW_HIDE)
	} else {
		Log("Debug-ShowOrHiddenWindow", GetLangText("Debug-ShowOrHiddenWindow_ShowWindow"), false)
		showWindow = true
		win.ShowWindow(consoleWindow, win.SW_SHOW)
	}
}
func RegHotKey() {
	err := qBCBHotkey.Register()
	if err != nil {
		Log("RegHotKey", GetLangText("Error-RegHotkey"), false, err.Error())
		return
	}
	Log("RegHotKey", GetLangText("Success-RegHotkey"), false)

	for range qBCBHotkey.Keydown() {
		ShowOrHiddenWindow()
	}
}
func RegSysTray() {
	systray.Run(func () {
		systray.SetIcon(icon_Windows)
		systray.SetTitle(programName)
		mShow := systray.AddMenuItem("显示/隐藏", "显示/隐藏程序")
		mQuit := systray.AddMenuItem("退出", "退出程序")

		for {
			select {
				case <-mShow.ClickedCh:
					ShowOrHiddenWindow()
				case <-mQuit.ClickedCh:
					systray.Quit()
			}
		}
	}, func () {
		ReqStop()
	})
}
func main() {
	if PrepareEnv() {
		go RegHotKey()
		go RegSysTray()
		RunConsole()
	}
}
func Platform_Stop() {
	qBCBHotkey.Unregister()
	systray.Quit()
}
