package main

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"os"

	"github.com/danielgatis/go-ansicode"
)

var _ ansicode.Handler = (*handler)(nil)

type handler struct{}

func (*handler) Backspace() {
	fmt.Printf("Backspace\n")
}

func (*handler) Bell() {
	fmt.Printf("Bell\n")
}

func (*handler) CarriageReturn() {
	fmt.Printf("CarriageReturn\n")
}

func (*handler) ClearLine(mode ansicode.LineClearMode) {
	fmt.Printf("ClearLine %v\n", mode)
}

func (*handler) ClearScreen(mode ansicode.ClearMode) {
	fmt.Printf("ClearScreen %v\n", mode)
}

func (*handler) ClearTabs(mode ansicode.TabulationClearMode) {
	fmt.Printf("ClearTabs %v\n", mode)
}

func (*handler) ClipboardLoad(clp byte, t string) {
	fmt.Printf("ClipboardLoad %v %v\n", clp, t)
}

func (*handler) ClipboardStore(clp byte, d []byte) {
	fmt.Printf("ClipboardStore %v %v\n", clp, d)
}

func (*handler) ConfigureCharset(index ansicode.CharsetIndex, charset ansicode.Charset) {
	fmt.Printf("ConfigureCharset %v %v\n", index, charset)
}

func (*handler) DeleteChars(n int) {
	fmt.Printf("DeleteChars %v\n", n)
}

func (*handler) DeleteLines(n int) {
	fmt.Printf("DeleteLines %v\n", n)
}

func (*handler) DeviceStatus(n int) {
	fmt.Printf("DeviceStatus %v\n", n)
}

func (*handler) EraseChars(n int) {
	fmt.Printf("EraseChars %v\n", n)
}

func (*handler) Goto(y int, x int) {
	fmt.Printf("Goto %v %v\n", y, x)
}

func (*handler) GotoCol(n int) {
	fmt.Printf("GotoCol %v\n", n)
}

func (*handler) GotoLine(n int) {
	fmt.Printf("GotoLine %v\n", n)
}

func (*handler) Input(r rune) {
	fmt.Printf("Input %v\n", r)
}

func (*handler) InsertBlank(n int) {
	fmt.Printf("InsertBlank %v\n", n)
}

func (*handler) InsertBlankLines(n int) {
	fmt.Printf("InsertBlankLines %v\n", n)
}

func (*handler) LineFeed() {
	fmt.Printf("LineFeed\n")
}

func (*handler) MoveBackward(n int) {
	fmt.Printf("MoveBackward %v\n", n)
}

func (*handler) MoveBackwardTabs(n int) {
	fmt.Printf("MoveBackwardTabs %v\n", n)
}

func (*handler) MoveDown(n int) {
	fmt.Printf("MoveDown %v\n", n)
}

func (*handler) MoveDownCr(n int) {
	fmt.Printf("MoveDownCr %v\n", n)
}

func (*handler) MoveForward(n int) {
	fmt.Printf("MoveForward %v\n", n)
}

func (*handler) MoveForwardTabs(n int) {
	fmt.Printf("MoveForwardTabs %v\n", n)
}

func (*handler) MoveUp(n int) {
	fmt.Printf("MoveUp %v\n", n)
}

func (*handler) MoveUpCr(n int) {
	fmt.Printf("MoveUpCr %v\n", n)
}

func (*handler) PopKeyboardMode(n int) {
	fmt.Printf("PopKeyboardMode %v\n", n)
}

func (*handler) PopTitle() {
	fmt.Printf("PopTitle\n")
}

func (*handler) PushKeyboardMode(mode ansicode.KeyboardMode) {
	fmt.Printf("PushKeyboardMode %v\n", mode)
}

func (*handler) PushTitle() {
	fmt.Printf("PushTitle\n")
}

func (*handler) ReportKeyboardMode() {
	fmt.Printf("ReportKeyboardMode\n")
}

func (*handler) ReportModifyOtherKeys() {
	fmt.Printf("ReportModifyOtherKeys\n")
}

func (*handler) ResetColor(i int) {
	fmt.Printf("ResetColor %v\n", i)
}

func (*handler) RestoreCursorPosition() {
	fmt.Printf("RestoreCursorPosition\n")
}

func (*handler) SaveCursorPosition() {
	fmt.Printf("SaveCursorPosition\n")
}

func (*handler) ScrollDown(n int) {
	fmt.Printf("ScrollDown %v\n", n)
}

func (*handler) ScrollUp(n int) {
	fmt.Printf("ScrollUp %v\n", n)
}

func (*handler) SetActiveCharset(cs int) {
	fmt.Printf("SetActiveCharset %v\n", cs)
}

func (*handler) SetColor(i int, c color.Color) {
	fmt.Printf("SetColor %v %v\n", i, c)
}

func (*handler) SetCursorStyle(style ansicode.CursorStyle) {
	fmt.Printf("SetCursorStyle %v\n", style)
}

func (*handler) SetDynamicColor(p string, i int, t string) {
	fmt.Printf("SetDynamicColor %v %v %v\n", p, i, t)
}

func (*handler) SetHyperlink(hyperlink *ansicode.Hyperlink) {
	fmt.Printf("SetHyperlink %v\n", hyperlink)
}

func (*handler) SetKeyboardMode(mode ansicode.KeyboardMode, behavior ansicode.KeyboardModeBehavior) {
	fmt.Printf("SetKeyboardMode %v %v\n", mode, behavior)
}

func (*handler) SetMode(mode ansicode.TerminalMode) {
	fmt.Printf("SetMode %v\n", mode)
}

func (*handler) SetModifyOtherKeys(modify ansicode.ModifyOtherKeys) {
	fmt.Printf("SetModifyOtherKeys %v\n", modify)
}

func (*handler) SetScrollingRegion(top int, bottom int) {
	fmt.Printf("SetScrollingRegion %v %v\n", top, bottom)
}

func (*handler) SetTerminalCharAttribute(attr ansicode.TerminalCharAttribute) {
	fmt.Printf("SetTerminalCharAttribute %v\n", attr)
}

func (*handler) SetTitle(title string) {
	fmt.Printf("SetTitle %v\n", title)
}

func (*handler) Substitute() {
	fmt.Printf("Substitute\n")
}

func (*handler) Tab(n int) {
	fmt.Printf("Tab %v\n", n)
}

func (*handler) TextAreaSizeChars() {
	fmt.Printf("TextAreaSizeChars\n")
}

func (*handler) TextAreaSizePixels() {
	fmt.Printf("TextAreaSizePixels\n")
}

func (*handler) UnsetMode(mode ansicode.TerminalMode) {
	fmt.Printf("UnsetMode %v\n", mode)
}

func (*handler) UnsetKeypadApplicationMode() {
	fmt.Printf("UnsetKeypadApplicationMode\n")
}

func (*handler) SetKeypadApplicationMode() {
	fmt.Printf("SetKeypadApplicationMode\n")
}

func (*handler) ReverseIndex() {
	fmt.Printf("ReverseIndex\n")
}

func (*handler) ResetState() {
	fmt.Printf("ResetState\n")
}

func (*handler) HorizontalTabSet() {
	fmt.Printf("HorizontalTabSet\n")
}

func (*handler) Decaln() {
	fmt.Printf("Decaln\n")
}

func (*handler) IdentifyTerminal(b byte) {
	fmt.Printf("IdentifyTerminal\n")
}

func main() {
	decoder := ansicode.NewDecoder(&handler{})

	reader := bufio.NewReader(os.Stdin)
	buff := make([]byte, 2048)

	for {
		n, err := reader.Read(buff)

		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Printf("Err %v:", err)
			return
		}

		for _, b := range buff[:n] {
			decoder.WriteByte(b)
		}
	}
}
