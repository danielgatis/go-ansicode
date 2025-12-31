package ansicode

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

// Hyperlink is used to represent a hyperlink.
type Hyperlink struct {
	// ID is the hyperlink ID.
	ID string

	// URI is the hyperlink URI.
	URI string
}

// OscDispatch is used to handle osc operations.
func (p *Performer) OscDispatch(params [][]byte, bellTerminated bool) {
	if len(params) == 0 || len(params[0]) == 0 {
		return
	}

	terminator := "\x1b\\"
	if bellTerminated {
		terminator = "\x07"
	}

	switch string(params[0]) {
	case "0", "2":
		if len(params) >= 2 {
			var buff bytes.Buffer
			for _, p := range params[1:] {
				buff.WriteString(string(p))
			}

			title := buff.String()
			title = strings.TrimSpace(title)
			p.handler.SetTitle(title)
			return
		}

		log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)

	case "7":
		// OSC 7 ; URI ST - Set working directory
		// Format: file://hostname/path/to/dir
		if len(params) >= 2 {
			var buff bytes.Buffer
			for i, param := range params[1:] {
				if i > 0 {
					buff.WriteString(";")
				}
				buff.Write(param)
			}
			uri := buff.String()
			p.handler.SetWorkingDirectory(uri)
			return
		}
		log.Debugf("Unhandled OSC 7 params=%v bellTerminated=%v", params, bellTerminated)

	case "4":
		if len(params) <= 1 || len(params)%2 == 0 {
			log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
			return
		}

		for i := 1; i < len(params); i += 2 {
			ps, ok := parseNumber(params[i])
			if !ok {
				log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
				continue
			}

			pt := params[i+1]

			color, ok := parseXColor(pt)
			if ok {
				p.handler.SetColor(int(ps), color)
			} else if string(pt) == "?" {
				prefix := fmt.Sprintf("4;%d", ps)
				p.handler.SetDynamicColor(prefix, int(ps), terminator)
			} else {
				log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
			}
		}

	case "8":
		if len(params) > 2 {
			linkParams := params[1]
			uri := string(params[2])

			for i := 3; i < len(params); i++ {
				uri += ";" + string(params[i])
			}

			if uri == "" {
				p.handler.SetHyperlink(nil)
				return
			}

			id := ""
			kvPairs := strings.Split(string(linkParams), ":")
			for _, kvPair := range kvPairs {
				if strings.HasPrefix(kvPair, "id=") {
					id = kvPair[3:]
					break
				}
			}

			p.handler.SetHyperlink(&Hyperlink{id, uri})
		}

	case "10", "11", "12":
		if len(params) < 2 {
			log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
			return
		}

		dynamicCode, ok := parseNumber(params[0])
		if !ok {
			return
		}

		for _, param := range params[1:] {
			offset := dynamicCode - 10
			index := int(NamedColorForeground) + offset

			if index > int(NamedColorCursor) {
				log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
				break
			}

			color, ok := parseXColor(param)
			if ok {
				p.handler.SetColor(int(index), color)
			} else if string(param) == "?" {
				p.handler.SetDynamicColor(strconv.Itoa(dynamicCode), index, terminator)
			} else {
				log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
			}

			dynamicCode++
		}

	case "104":
		if len(params) == 1 || len(params[1]) == 0 {
			for i := 0; i < 256; i++ {
				p.handler.ResetColor(i)
			}

			return
		}

		for _, param := range params[1:] {
			i, err := strconv.Atoi(string(param))
			if err != nil {
				continue
			}

			p.handler.ResetColor(i)
		}

	case "110":
		p.handler.ResetColor(int(NamedColorForeground))

	case "111":
		p.handler.ResetColor(int(NamedColorBackground))

	case "112":
		p.handler.ResetColor(int(NamedColorCursor))

	case "133":
		// Shell Integration (FinalTerm/iTerm2 style)
		// OSC 133 ; A ST - Prompt start
		// OSC 133 ; B ST - Command start (after prompt)
		// OSC 133 ; C ST - Command executed
		// OSC 133 ; D [; exitcode] ST - Command finished
		if len(params) < 2 {
			log.Debugf("Unhandled OSC 133 params=%v bellTerminated=%v", params, bellTerminated)
			return
		}

		cmd := string(params[1])
		switch cmd {
		case "A":
			p.handler.ShellIntegrationMark(PromptStart, -1)
		case "B":
			p.handler.ShellIntegrationMark(CommandStart, -1)
		case "C":
			p.handler.ShellIntegrationMark(CommandExecuted, -1)
		case "D":
			exitCode := -1
			if len(params) >= 3 {
				if code, ok := parseNumber(params[2]); ok {
					exitCode = code
				}
			}
			p.handler.ShellIntegrationMark(CommandFinished, exitCode)
		default:
			log.Debugf("Unhandled OSC 133 command=%s params=%v", cmd, params)
		}

	case "99":
		// Desktop Notifications (Kitty protocol)
		// Format: OSC 99 ; metadata ; payload ST
		// See: https://sw.kovidgoyal.net/kitty/desktop-notifications/
		payload := parseOSC99(params)
		if payload != nil {
			p.handler.DesktopNotification(payload)
		}

	case "1337":
		// iTerm2/WezTerm proprietary sequences
		// OSC 1337 ; SetUserVar=NAME=BASE64_VALUE ST
		if len(params) >= 2 {
			data := string(params[1])
			if strings.HasPrefix(data, "SetUserVar=") {
				// Parse NAME=BASE64_VALUE
				rest := data[len("SetUserVar="):]
				eqIdx := strings.Index(rest, "=")
				if eqIdx > 0 {
					name := rest[:eqIdx]
					encodedValue := rest[eqIdx+1:]
					// Decode base64 value
					decoded, err := base64.StdEncoding.DecodeString(encodedValue)
					if err == nil {
						p.handler.SetUserVar(name, string(decoded))
					} else {
						// Try without padding
						decoded, err = base64.RawStdEncoding.DecodeString(encodedValue)
						if err == nil {
							p.handler.SetUserVar(name, string(decoded))
						}
					}
				}
			}
		}

	default:
		log.Debugf("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
	}
}

func parseXColor(bytes []byte) (color.Color, bool) {
	if len(bytes) == 0 {
		return color.RGBA{}, false
	}

	rgb := make([]uint8, 0)

	if len(bytes) > 0 && string(bytes[0]) == "#" {
		colors := string(bytes[1:])
		colorLen := len(colors) / 3

		for i := 0; i < 3; i++ {
			c := colors[i*colorLen : (i+1)*colorLen]
			max, err := strconv.ParseUint(strings.Repeat("F", len(c)), 16, 0)
			if err != nil {
				break
			}

			value, err := strconv.ParseUint(c, 16, 0)
			if err != nil {
				break
			}

			scaled := uint8(255 * value / max)
			rgb = append(rgb, scaled)
		}
	}

	if len(bytes) >= 4 && string(bytes[:4]) == "rgb:" {
		colors := strings.Split(string(bytes[4:]), "/")

		if len(colors) != 3 {
			return color.RGBA{}, false
		}

		for _, c := range colors {
			max, err := strconv.ParseUint(strings.Repeat("F", len(c)), 16, 0)
			if err != nil {
				break
			}

			value, err := strconv.ParseUint(c, 16, 0)
			if err != nil {
				break
			}

			scaled := uint8(255 * value / max)
			rgb = append(rgb, scaled)
		}
	}

	if len(rgb) == 3 {
		return color.RGBA{
			R: rgb[0],
			G: rgb[1],
			B: rgb[2],
			A: 255,
		}, true
	}

	return color.RGBA{}, false
}

func parseNumber(bytes []byte) (int, bool) {
	if len(bytes) == 0 {
		return 0, false
	}

	num := 0
	for _, b := range bytes {
		digit, err := strconv.Atoi(string(b))
		if err != nil {
			return 0, false
		}

		num = num*10 + digit
	}

	return num, true
}

// parseOSC99 parses an OSC 99 desktop notification sequence.
// Format: OSC 99 ; metadata ; payload ST
// metadata is key=value pairs separated by ':'
// payload is the notification data (title, body, icon, etc.)
func parseOSC99(params [][]byte) *NotificationPayload {
	// params[0] = "99"
	// params[1] = metadata (may be empty)
	// params[2] = payload data (may be empty)

	payload := &NotificationPayload{
		Done:    true, // default: complete notification
		Urgency: 1,    // default: normal urgency
		Timeout: -1,   // default: OS decides
	}

	// Parse metadata if present
	if len(params) >= 2 && len(params[1]) > 0 {
		parseOSC99Metadata(params[1], payload)
	}

	// Parse payload data if present
	if len(params) >= 3 && len(params[2]) > 0 {
		data := params[2]

		// Decode base64 if encoding is set
		if payload.Encoding == "1" {
			decoded, err := base64.StdEncoding.DecodeString(string(data))
			if err == nil {
				payload.Data = decoded
			} else {
				// Try without padding
				decoded, err = base64.RawStdEncoding.DecodeString(string(data))
				if err == nil {
					payload.Data = decoded
				} else {
					payload.Data = data
				}
			}
		} else {
			payload.Data = data
		}
	}

	return payload
}

// parseOSC99Metadata parses the metadata portion of an OSC 99 sequence.
// Format: key=value:key=value:...
func parseOSC99Metadata(meta []byte, p *NotificationPayload) {
	pairs := bytes.Split(meta, []byte(":"))
	for _, pair := range pairs {
		idx := bytes.IndexByte(pair, '=')
		if idx == -1 {
			continue
		}

		key := string(pair[:idx])
		value := string(pair[idx+1:])

		switch key {
		case "i":
			p.ID = value
		case "d":
			p.Done = value != "0"
		case "p":
			p.PayloadType = value
		case "e":
			p.Encoding = value
		case "a":
			p.Actions = strings.Split(value, ",")
		case "c":
			p.TrackClose = value == "1"
		case "w":
			if timeout, err := strconv.Atoi(value); err == nil {
				p.Timeout = timeout
			}
		case "f":
			// AppName is base64 encoded
			if decoded, err := base64.StdEncoding.DecodeString(value); err == nil {
				p.AppName = string(decoded)
			} else {
				p.AppName = value
			}
		case "t":
			// Type is base64 encoded
			if decoded, err := base64.StdEncoding.DecodeString(value); err == nil {
				p.Type = string(decoded)
			} else {
				p.Type = value
			}
		case "n":
			p.IconName = value
		case "g":
			p.IconCacheID = value
		case "s":
			p.Sound = value
		case "u":
			if urgency, err := strconv.Atoi(value); err == nil && urgency >= 0 && urgency <= 2 {
				p.Urgency = urgency
			}
		case "o":
			p.Occasion = value
		}
	}
}
