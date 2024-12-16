package hooks

import (
	"notes-app-demo-backend/hooks/notes"

	"github.com/pocketbase/pocketbase/core"
)

func Bind(app core.App) {
	app.OnRecordCreateRequest("notes").BindFunc(notes.CheckContentProfanity)
}
