package hooks

import (
	"notes-app-demo-backend/hooks/notes"
	"notes-app-demo-backend/hooks/users"

	"github.com/pocketbase/pocketbase/core"
)

func Bind(app core.App) {
	app.OnRecordCreateRequest("notes").BindFunc(notes.CheckContentProfanity)
	app.OnRecordCreateRequest("notes").BindFunc(notes.CheckContentGibberish)
	app.OnRecordUpdateRequest("users").BindFunc(users.CheckNameProfanity)
}
