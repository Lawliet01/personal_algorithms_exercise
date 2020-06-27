package stack

import "fmt"

type browserStack struct {
	forwardStack  *arrayStack
	backwardStack *arrayStack
	currentBrowse string
}

func InitBrowserStack() *browserStack {
	return &browserStack{
		forwardStack:  Init(),
		backwardStack: Init(),
	}
}

func (browser *browserStack) browse(website string) {
	if browser.currentBrowse != "" {
		browser.backwardStack.push(browser.currentBrowse)
	}
	browser.currentBrowse = website
}

func (browser *browserStack) backward() {
	if browser.backwardStack.length != 0 {
		browser.forwardStack.push(browser.currentBrowse)
		browser.currentBrowse = fmt.Sprint(browser.backwardStack.pop())
	}
}

func (browser *browserStack) forward() {
	if browser.forwardStack.length != 0 {
		browser.backwardStack.push(browser.currentBrowse)
		browser.currentBrowse = fmt.Sprint(browser.forwardStack.pop())
	}
}

func (browser *browserStack) String() string {
	return fmt.Sprintf("forwardStack:%s\nbackwardStack:%s\ncurrentBrowse:%s\n", browser.forwardStack, browser.backwardStack, browser.currentBrowse)
}
