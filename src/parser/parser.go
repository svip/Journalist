package parser

import (
	"bytes"
	"chatlog"
	"log"
	"regexp"
	"strings"
)

func lineTrim(line string) string {
	return strings.TrimRight(strings.Trim(line, " "), ".!")
}

func ParseIRCLogBytes(data []byte) chatlog.Chatlog {
	out := bytes.NewBuffer(data)

	return ParseIRCLog(out.String())
}

func ParseIRCLog(data string) chatlog.Chatlog {
	re, err := regexp.Compile("(?m)^((\\d{4}-\\d{2}-\\d{2} )?(\\d{2}:)?\\d{2}:\\d{2} +<)")
	if err != nil {
		log.Println(err)
	}
	tmplog := re.Split(data, -1)
	var clog chatlog.Chatlog
	nre, _ := regexp.Compile("(?m)[ ^]+")
	rre, _ := regexp.Compile("(?i)^([a-z0-9]+):")
	for _, line := range tmplog {
		if strings.Trim(line, " ") == "" {
			continue
		}
		line = strings.Trim(nre.ReplaceAllString(line, " "), " \n")
		s := strings.SplitN(line, ">", 2)
		nl := lineTrim(s[1])
		person := clog.GetPerson(s[0])
		if person == nil {
			clog = append(clog, chatlog.Line{Person: &chatlog.Person{s[0], s[0]}, Line: nl})
		} else {
			clog = append(clog, chatlog.Line{Person: person, Line: nl})
		}
		clog[len(clog)-1].Type = chatlog.NormalMessage
		if rre.MatchString(nl) {
			ss := rre.FindStringSubmatch(nl)
			l := strings.SplitN(line, ":", 2)
			receiver := clog.GetPerson(ss[1])
			if receiver == nil {
				clog[len(clog)-1].Receiver = &chatlog.Person{ss[1], ss[1]}
			} else {
				clog[len(clog)-1].Receiver = receiver
			}
			nll := lineTrim(l[1])
			clog[len(clog)-1].Line = nll
			clog[len(clog)-1].Type = chatlog.ToMessage
		}
	}
	return clog
}

