package users

import (
	goaway "github.com/TwiN/go-away"
	"github.com/pocketbase/pocketbase/core"
)



func CheckNameProfanity(e *core.RecordRequestEvent) error {
	name := e.Record.GetString("name")

	if (goaway.IsProfane(name)) {
		return e.BadRequestError("Please set an appropriate name", nil)
	}

	return e.Next()
}