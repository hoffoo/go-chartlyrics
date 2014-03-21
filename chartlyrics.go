package chartlyrics

import (
    "fmt"
    "net/url"
    "strings"
)

// ---------------------------------------------------------
// generic response
// ---------------------------------------------------------

type ChartLyricsResult struct {

    // search results go here
    SearchLyricResult []SearchLyricResult

    // direct query (not search) go here
    LyricChecksum string
    LyricId       string
    LyricSong     string
    LyricArtist   string
    LyricUrl      string

    LyricCoverArtUrl string
    Lyric            string
}

type SearchLyricResult struct {
    LyricChecksum string
    LyricId       string
}

// ---------------------------------------------------------
// querying the api data logic
// ---------------------------------------------------------

func GetLyric(lyricId, lyricCheckSum string) (clr ChartLyricsResult, err error) {

    err = searchValid(checkEmptyStrings, lyricId, lyricCheckSum)
    if err != nil {
        return
    }

    return HttpGetLyric(toQuery("lyricId", lyricId, "lyricChecksum", lyricCheckSum))
}

func SearchLyric(artist, song string) (clr ChartLyricsResult, err error) {

    err = searchValid(checkEmptyStrings, artist, song)
    if err != nil {
        return
    }

    return HttpSearchLyric(toQuery("artist", artist, "song", song))
}

func SearchLyricDirect(artist, song string) (clr ChartLyricsResult, err error) {
    err = searchValid(checkEmptyStrings, artist, song)
    if err != nil {
        return
    }

    return HttpSearchLyricDirect(toQuery("artist", artist, "song", song))
}

// ---------------------------------------------------------
// Search datastructure
// ---------------------------------------------------------

type Search struct {
    Artist        string
    Song          string
    LyricChecksum string
    LyricId       string
}

// ---------------------------------------------------------
// Search methods
// ---------------------------------------------------------

func (s *Search) SearchLyric() (clr ChartLyricsResult, err error) {
    return SearchLyric(s.Artist, s.Song)
}

func (s *Search) SearchLyricDirect() (clr ChartLyricsResult, err error) {
    return SearchLyricDirect(s.Artist, s.Song)
}

func (s *Search) GetLyric() (clr ChartLyricsResult, err error) {
    return GetLyric(s.LyricId, s.LyricChecksum)
}

// ---------------------------------------------------------
// generic util
// ---------------------------------------------------------

func toQuery(s ...string) string {
    q := url.Values{}
    for i := 0; i < len(s); i += 2 {
        q.Add(s[i], s[i+1])
    }

    return q.Encode()
}

// ---------------------------------------------------------
// cleanup api responses
// ---------------------------------------------------------

// used to remove junk or empty data from SearchLyricResult
func cleanupSearchLyricResult(sr *[]SearchLyricResult) {
    var newsr []SearchLyricResult

    for _, k := range *sr {
        if len(k.LyricChecksum) <= 1 {
            continue
        }

        newsr = append(newsr, k)
    }

    *sr = newsr
}

// ---------------------------------------------------------
// error catching
// ---------------------------------------------------------
func searchValid(vfn func(...string) bool, s ...string) error {

    if !vfn(s...) {
        return fmt.Errorf("chartlyrics: invalid search %v\n", s)
    }

    return nil
}

func checkEmptyStrings(args ...string) bool {
    for _, f := range args {
        if strings.Trim(f, "") == "" {
            return false
        }
    }
    return true
}
