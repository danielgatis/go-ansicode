package ansicode

import (
	"bytes"
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

		p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)

	case "4":
		if len(params) <= 1 || len(params)%2 == 0 {
			p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
			return
		}

		for i := 1; i < len(params); i += 2 {
			ps, ok := parseNumber(params[i])
			if !ok {
				p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
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
				p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
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
			p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
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
				p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
				break
			}

			color, ok := parseXColor(param)
			if ok {
				p.handler.SetColor(int(index), color)
			} else if string(param) == "?" {
				p.handler.SetDynamicColor(strconv.Itoa(dynamicCode), index, terminator)
			} else {
				p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
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

	default:
		p.logger.Tracef("Unhandled OSC params=%v bellTerminated=%v", params, bellTerminated)
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
