package model

type Help struct{ *Entry }

func NewHelp(clock Clock, author, title string, standup Standup) interface{} {
	return Help{NewEntry(clock, author, title, standup)}
}

func (help Help) MakeCreateRequest() (request WhiteboardRequest) {
	request = help.Entry.MakeCreateRequest()
	request.Item.Kind = "Help"
	return
}

func (help Help) MakeUpdateRequest() (request WhiteboardRequest) {
	request = help.Entry.MakeUpdateRequest()
	request.Item.Kind = "Help"
	return
}

func (help Help) GetEntry() *Entry {
	return help.Entry
}