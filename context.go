package gram

import "context"

type Context struct {
	context.Context
	ID      string
	ChatID  string
	Cmd     string
	CmdUser string
}

type CmdArgs struct {
}

func (c *CmdArgs) Get(out interface{}) error {
	return nil
}

func GetContext(chatID string) *Context {
	return &Context{}
}

func (c *Context) Cancel() {

}

func (c *Context) Send(s string) error {
	return nil
}

func (c *Context) Get(out interface{}) error {
	return nil
}

func (c *Context) CmdArgs() *CmdArgs {
	return nil
}

func (c *Context) Promp(description string, button Button, otherButtons ...Button) error {
	return nil
}

func (c *Context) SetKeyboard(buttonGroup []ButtonGroup, otherButtonGroups ...[]ButtonGroup) error {
	return nil
}

func (c *Context) ShowNotification(s string) error {
	return nil
}

func (c *Context) Switch(pipeline *Pipeline) {

}
