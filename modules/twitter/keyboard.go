package twitter

import (
	"github.com/gdamore/tcell"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.SetKeyboardChar("/", widget.ShowHelp, "Show/hide this help prompt")
	widget.SetKeyboardChar("h", widget.Prev, "Select previous item")
	widget.SetKeyboardChar("l", widget.Next, "Select next item")
	widget.SetKeyboardChar("o", widget.openFile, "Open item")

	widget.SetKeyboardKey(tcell.KeyEnter, widget.openFile, "Open item")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyRight, widget.Next, "Select next item")
}

func (widget *Widget) openFile() {
	src := widget.currentSourceURI()
	wtf.OpenFile(src)
}
