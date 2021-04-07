package repository

import {
	"testing"

}

func TestLoadData(t *testing.T) { 
	t.Run("loading data!", func(t *testing.T) {
		var jobList = repository.LoadData()

		got := buffer.String()
		want := ``

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}