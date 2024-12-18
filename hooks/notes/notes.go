package notes

import (
	"fmt"

	goaway "github.com/TwiN/go-away"
	"github.com/kevincobain2000/go-gibberish/gibberish"
	"github.com/pocketbase/pocketbase/core"
)
	

func CheckContentProfanity(e *core.RecordRequestEvent) error {
    content := e.Record.GetString("content")
    name := e.Auth.GetString("name")

	if goaway.IsProfane(content) {
		profanity := goaway.ExtractProfanity(content)
		return e.BadRequestError(fmt.Sprintf("C'mon %s, let's keep this a safe space! 🙈 Please avoid using the word '%s'.", name, profanity), nil)
	}

    return e.Next()
}


var gib = gibberish.NewGibberish()

func CheckContentGibberish(e *core.RecordRequestEvent) error {
    content := e.Record.GetString("content")
    name := e.Auth.GetString("name")

    if gib.Detect(content).IsGibberish {
        return e.BadRequestError(fmt.Sprintf("Hey %s, your message seems like gibberish. Please make sure you're writing something meaningful! 🤔", name), nil)
    }

    return e.Next()
}

