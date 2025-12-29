# go-ansicode

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-ansicode?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-ansicode)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-ansicode/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-ansicode)

go-ansicode is a package that interprets ANSI codes and allows you to register a handler to deal with the operations.

## Supported Sequences

### C0

| Operation | Description                      |
|-----------|----------------------------------|
| NUL       | Null (NUL) character.            |
| SOH       | Start of Heading (SOH).          |
| STX       | Start of Text (STX).             |
| ETX       | End of Text (ETX).               |
| EOT       | End of Transmission (EOT).       |
| ENQ       | Enquiry (ENQ).                   |
| ACK       | Acknowledge (ACK).               |
| BEL       | Bell (BEL).                      |
| BS        | Backspace (BS).                  |
| HT        | Horizontal Tab (HT).             |
| LF        | Line Feed (LF).                  |
| VT        | Vertical Tab (VT).               |
| FF        | Form Feed (FF).                  |
| CR        | Carriage Return (CR).            |
| SO        | Shift Out (SO).                  |
| SI        | Shift In (SI).                   |
| DLE       | Data Link Escape (DLE).          |
| XON       | Device Control 1 (XON).          |
| DC2       | Device Control 2 (DC2).          |
| XOFF      | Device Control 3 (XOFF).         |
| DC4       | Device Control 4 (DC4).          |
| NAK       | Negative Acknowledge (NAK).      |
| SYN       | Synchronous Idle (SYN).          |
| ETB       | End of Transmission Block (ETB). |
| CAN       | Cancel (CAN).                    |
| EM        | End of Medium (EM).              |
| SUB       | Substitute (SUB).                |
| ESC       | Escape (ESC).                    |
| FS        | File Separator (FS).             |
| GS        | Group Separator (GS).            |
| RS        | Record Separator (RS).           |


### C1

| Operation | Description                                               |
|-----------|-----------------------------------------------------------|
| IND       | Index - Move to the next line (Line Feed)                 |
| NEL       | Next Line - Line Feed followed by Carriage Return         |
| HTS       | Horizontal Tab Set                                        |
| RI        | Reverse Index - Move to the previous line                 |
| DECID     | Identify Terminal - Request for terminal identification   |
| RIS       | Reset to Initial State - Reset all settings to defaults   |
| DECSC     | Save Cursor - Save current cursor position and attributes |
| DECRC     | Restore Cursor - Restore cursor position and attributes   |
| DECALN    | Alignment Test - Perform an alignment test                |
| DECKPAM   | Keypad Application Mode - Enable Keypad Application Mode  |
| DECKPNM   | Numeric Mode - Disable Keypad Application Mode            |


### CSI

| CSI Sequence      | Description                                          |
|-------------------|------------------------------------------------------|
| `CSI Ps '`        | Single-character tabulation set (HTS)                |
| `CSI Ps @`        | Character tabulation with justification (CHA)        |
| `CSI Ps A`        | Line tabulation set (VTS)                            |
| `CSI Ps B`        | Partial line forward (incomplete)                    |
| `CSI Ps b`        | Repeat character (REP)                               |
| `CSI Ps C`        | Cursor forward (CUF)                                 |
| `CSI Ps c`        | Send device attributes (DA)                          |
| `CSI Ps a`        | Character position forward (CHA)                     |
| `CSI Ps D`        | Cursor backward (CUB)                                |
| `CSI Ps d`        | Vertical position absolute (VPA)                     |
| `CSI Ps E`        | Next line (NEL)                                      |
| `CSI Ps e`        | Vertical position relative (VPR)                     |
| `CSI Ps F`        | Previous line (RI)                                   |
| `CSI Ps G`        | Horizontal tab set (HTS)                             |
| `CSI Ps g`        | Tab clear (TBC)                                      |
| `CSI Ps ; Ps H`   | Cursor position (CUP)                                |
| `CSI Ps ; Ps f`   | Horizontal and vertical position (HVP)               |
| `CSI Ps I`        | Forward tabulation (HT)                              |
| `CSI Ps J`        | Erase in display (ED)                                |
| `CSI Ps K`        | Erase in line (EL)                                   |
| `CSI Ps L`        | Insert line (IL)                                     |
| `CSI Ps l`        | Reset mode (DEC)                                     |
| `CSI ? Ps l`      | DEC private mode reset                               |
| `CSI Ps M`        | Delete line (DL)                                     |
| `CSI Pm m`        | Character attribute (SGR)                            |
| `CSI > Pp ; Pv m` | Select character protection attribute (DECSCPP)      |
| `CSI ? Pp m`      | Select media character set and invoke macro (DECSEL) |
| `CSI Ps n`        | Device status report (DSR)                           |
| `CSI Ps P`        | Delete character (DCH)                               |
| `CSI Ps SP q`     | Select modifier and use bit combination (SMRM)       |
| `CSI Ps ; Ps r`   | Set top and bottom margins (DECSTBM)                 |
| `CSI Ps S`        | Scroll up (SU)                                       |
| `CSI s`           | Save cursor position (SCP)                           |
| `CSI Ps T`        | Scroll down (SD)                                     |
| `CSI Ps t`        | Window manipulation (DECSWT)                         |
| `CSI u`           | Restore cursor position (RCP)                        |
| `CSI ? u`         | DEC private mode reset                               |
| `CSI = Ps ; Ps u` | Set conformance level (DECSCL)                       |
| `CSI > Ps u`      | Set ANSI conformance level (DECSASD)                 |
| `CSI < Ps u`      | Set conformance level (DECSCL)                       |
| `CSI Ps X`        | Erase character (ECH)                                |
| `CSI Ps Z`        | Cursor back tabulation (CBT)                         |


### OSC

| OSC Sequence                    | Description                                |
|---------------------------------|--------------------------------------------|
| `OSC 0 ; Pt BEL`                | Set icon name and window title             |
| `OSC 2 ; Pt BEL`                | Set window title                           |
| `OSC 4 ; c ; spec BEL`          | Change color in palette (8/16 colors)      |
| `OSC 7 ; URI BEL`               | Set working directory                      |
| `OSC 8 ; params ; uri BEL`      | Set hyperlinks                             |
| `OSC 10 ; Ps BEL`               | Set foreground text color                  |
| `OSC 11 ; Ps BEL`               | Set background text color                  |
| `OSC 12 ; Ps BEL`               | Set cursor text color                      |
| `OSC 99 ; metadata ; data BEL`  | Desktop notifications (Kitty protocol)     |
| `OSC 104 ; c BEL`               | Reset color in palette (8/16 colors)       |
| `OSC 110 BEL`                   | Reset icon name and window title           |
| `OSC 111 BEL`                   | Reset window title                         |
| `OSC 112 BEL`                   | Reset color in palette (24-bit)            |
| `OSC 133 ; cmd BEL`             | Shell integration marks                    |


### OSC 99 - Desktop Notifications (Kitty Protocol)

Desktop notifications allow terminal applications to send system notifications. The format is:

```
OSC 99 ; metadata ; payload ST
```

**Metadata fields** (colon-separated key=value pairs):

| Key | Description                                              |
|-----|----------------------------------------------------------|
| `i` | Notification ID for tracking/chunking                    |
| `d` | Done flag (0=more chunks coming, 1=complete)             |
| `p` | Payload type: title, body, icon, buttons, close, alive, ?|
| `e` | Encoding (1=base64)                                      |
| `a` | Actions on click (focus, report)                         |
| `c` | Track close events (1=yes)                               |
| `w` | Timeout in ms (-1=OS default, 0=never)                   |
| `f` | Application name (base64 encoded)                        |
| `t` | Notification type for filtering (base64 encoded)         |
| `n` | Icon name: error, warning, info, question                |
| `g` | Icon cache UUID                                          |
| `s` | Sound: system, silent, error, warn, info                 |
| `u` | Urgency: 0=low, 1=normal, 2=critical                     |
| `o` | Occasion: always, unfocused, invisible                   |

**Example:**
```bash
# Simple notification
printf '\033]99;;Hello World\007'

# Notification with title and urgency
printf '\033]99;i=1:p=title;My Title\007'
printf '\033]99;i=1:p=body:u=2;Critical message!\007'
```

See: https://sw.kovidgoyal.net/kitty/desktop-notifications/


### SOS/PM/APC

| Sequence       | Description                                                               |
|----------------|---------------------------------------------------------------------------|
| `ESC X ... ST` | Start of String (SOS) - Application-specific string data                  |
| `ESC ^ ... ST` | Privacy Message (PM) - Private message for terminal                       |
| `ESC _ ... ST` | Application Program Command (APC) - Used by protocols like Kitty Graphics |

These sequences can also be terminated by BEL (0x07) instead of ST (ESC \).


## Install

```bash
go get -u github.com/danielgatis/go-ansicode
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-ansicode"
```

### Example

Please look at: [examples/ansilog/main.go](examples/ansilog/main.go)

```
❯ echo -ne "\033[31;42mThis text is red on a green background\033[0m\nbye" | go run ./examples/ansilog/main.go
SetTerminalCharAttribute {22 0x14000112018 <nil> <nil>}
SetTerminalCharAttribute {23 0x14000112020 <nil> <nil>}
Input 84
Input 104
Input 105
Input 115
Input 32
Input 116
Input 101
Input 120
Input 116
Input 32
Input 105
Input 115
Input 32
Input 114
Input 101
Input 100
Input 32
Input 111
Input 110
Input 32
Input 97
Input 32
Input 103
Input 114
Input 101
Input 101
Input 110
Input 32
Input 98
Input 97
Input 99
Input 107
Input 103
Input 114
Input 111
Input 117
Input 110
Input 100
SetTerminalCharAttribute {0 <nil> <nil> <nil>}
LineFeed
Input 98
Input 121
Input 101
```

## License

Copyright (c) 2023-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)


