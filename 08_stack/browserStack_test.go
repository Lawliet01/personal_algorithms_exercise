package stack

import (
	"fmt"
	"testing"
)

func TestIsBrowserStackWork(t *testing.T) {
	browser := InitBrowserStack()
	browser.browse("www.qq.com")
	browser.browse("www.pornhub.com")
	browser.browse("www.google.com")
	browser.backward()
	browser.backward()
	fmt.Println(browser)
}
