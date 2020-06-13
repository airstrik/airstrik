package engine

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/tebeka/selenium"
)

type Browser struct {
	BrowserId    string
	Type         string
	ChildBrowser []Browser
	wd           selenium.WebDriver
}

func NewBrowser(browserType string) *Browser {

	uuid := uuid.New().String()
	if browserType == "" {
		browserType = "chrome"
	}
	caps := selenium.Capabilities{"browserName": browserType}
	wd, _ := selenium.NewRemote(caps, "")

	return &Browser{BrowserId: uuid, Type: browserType, wd: wd}
}

func (b *Browser) Navigate(url string) {
	fmt.Println("Navigating to the URl for the Testing " + url)
	b.wd.Get(url)
}

func (b *Browser) Close() {
	fmt.Print("Tapering Down the Browser")
	defer b.wd.Quit()
}
