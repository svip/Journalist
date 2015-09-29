package journalist

import (
	"bytes"
	"log"
	"math/rand"
	"text/template"
	"chatlog"
)

func createToMessageLine(line chatlog.Line) string {
	text := phrases.ToMessage[rand.Intn(len(phrases.ToMessage))]
	t, err := template.New("line").Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	out := bytes.NewBufferString("")
	t.Execute(out, struct {
		Speaker  string
		FullLine string
		Receiver string
	}{
		line.Person.Name,
		line.Line,
		line.Receiver.Name,
	})
	return out.String()
}

func createNormalMessageLine(line chatlog.Line) string {
	text := phrases.NormalMessage[rand.Intn(len(phrases.NormalMessage))]
	t, err := template.New("line").Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	out := bytes.NewBufferString("")
	err = t.Execute(out, struct {
		Speaker        string
		SpeakerPronoun string
		FullLine       string
		Line1          string
		Line2          string
	}{
		line.Person.Name,
		"han",
		line.Line,
		"",
		"",
	})
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func writeArticle(clog chatlog.Chatlog) []string {
	var article []string
	for _, line := range clog {
		var l string
		switch line.Type {
		case chatlog.ToMessage:
			l = createToMessageLine(line)
		default:
			l = createNormalMessageLine(line)
		}
		article = append(article, l)
	}
	return article
}

func GenerateArticle(clog chatlog.Chatlog) []string {
	collectPhrases()

	article := writeArticle(clog)

	return article
}

