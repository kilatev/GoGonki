package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
  greeting := "Aloha World"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
  c.Validation.Required(myName).Message("Your Name is required!")
  c.Validation.MinSize(myName, 3).Message("Your name should be at least 3 symbols")
  
  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(App.Index)
  }
  
  return c.Render(myName)
}