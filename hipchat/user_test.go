package hipchat

import (
	"net/http"
	"testing"
	"io/ioutil"
	"os"
)

func TestUserShareFile(t *testing.T) {
	setup()
	defer teardown()

	tempFile, err := ioutil.TempFile(os.TempDir(), "hipfile")
	tempFile.WriteString("go gophers")
	defer os.Remove(tempFile.Name())

	want := "--hipfileboundary\n" +
		"Content-Type: application/json; charset=UTF-8\n" +
		"Content-Disposition: attachment; name=\"metadata\"\n\n" +
		"{\"message\": \"Hello there\"}\n" +
		"--hipfileboundary\n" +
		"Content-Type:  charset=UTF-8\n" +
		"Content-Transfer-Encoding: base64\n" +
		"Content-Disposition: attachment; name=file; filename=hipfile\n\n" +
		"Z28gZ29waGVycw==\n" +
		"--hipfileboundary\n"

	mux.HandleFunc("/user/1/share/file", func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method %s, want %s", r.Method, m)
		}

		body, _ := ioutil.ReadAll(r.Body)

		if string(body) != want {
			t.Errorf("Request body \n%+v\n,want \n\n%+v", string(body), want)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	args := &ShareFileRequest{Path: tempFile.Name(), Message: "Hello there", Filename: "hipfile"}
	_, err = client.User.ShareFile("1", args)
	if err != nil {
		t.Fatalf("User.ShareFile returns an error %v", err)
	}
}

