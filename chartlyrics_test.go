package chartlyrics

import (
    "testing"
)

func TestHttp(t *testing.T) {

    var err error
    var r ChartLyricsResult

    s := Search{Artist: "Aesop Rock", Song: "None Shall Pass"}

    r, err = s.SearchLyricDirect(60)

    t.Errorf("%+v", r)

    r, err = s.SearchLyric(60)

    if err != nil {
        t.Fatal(err)
    }

    if len(r.SearchLyricResult) != 1 {
        t.Errorf("there should have been one result. %+v ", r.SearchLyricResult)
    }

    getLyric, err := GetLyric(r.SearchLyricResult[0].LyricId, r.SearchLyricResult[0].LyricChecksum)

    if err != nil {
        t.Fatal(err)
    }

    t.Errorf("%+v", getLyric)
}

func TestValidateEmptyStr(t *testing.T) {

    var err error

    s := Search{Artist: "ff", Song: ""}

    _, err = s.SearchLyric()

    if err == nil {
        t.Fatal("should have failed")
    }

    _, err = s.GetLyric()

    if err == nil {
        t.Fatal("should have failed")
    }
}
