package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	database, cleanDatabase := createTmpFile(t, `[
		{"Name": "Niels", "Wins": 10},
		{"Name": "Cleo", "Wins": 33}]`)
	defer cleanDatabase()

	store, _ := NewFileSystemPlayerStore(database)

	t.Run("get the whole league", func(t *testing.T) {
		got := store.GetLeague()

		want := []Player{
			{"Cleo", 33},
			{"Niels", 10},
		}

		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Niels")
		want := 10

		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Niels")

		got := store.GetPlayerScore("Niels")
		want := 11

		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTmpFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})
}

func createTmpFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create tmp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
