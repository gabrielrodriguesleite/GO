package concorrencia

import "testing"

func mockVerificadorWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestVerificaWebsites(t *testing.T) {

}
