package sound

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dbatbold/beep"
)

// DownloadVoiceFiles downloads natural voice files
func DownloadVoiceFiles(music *beep.Music, writer io.Writer, names []string) error {
	dir := filepath.Join(beep.HomeDir(), "voices")
	if len(names) == 0 {
		names = []string{"piano", "violin"}
	}
	for _, name := range names {
		if !strings.HasSuffix(name, ".zip") {
			name += ".zip"
		}
		url := "http://bmrust.com/dl/beep/voices/" + name

		// locate file
		resp, err := http.Head(url)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return err
		}

		// fetch file
		resp, err = http.Get(url)
		fmt.Println("start Download ...", name)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return err
		}
		defer resp.Body.Close()

		// read file
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println("end Download ...", name)
		if err != nil {
			return err
		}

		// save file
		os.MkdirAll(dir, 0755)
		filename := filepath.Join(dir, name)
		err = ioutil.WriteFile(filename, body, 0644)
		if err != nil {
			return err
		}

		fmt.Fprintf(writer, "files Saving %s\n", filename)
	}

	return nil
}
