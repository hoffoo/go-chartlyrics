package chartlyrics

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// first arg is function, second is url encoded params
const (
	addr                 = "http://api.chartlyrics.com/apiv1.asmx/%s?%s"
	_SEARCH_LYRIC        = "SearchLyric"
	_SEARCH_LYRIC_DIRECT = "SearchLyricDirect"
	_GET_LYRIC           = "GetLyric"

	// TODO not implemented
	_SEARCH_LYRIC_TEXT = "SearchLyricText"
	_ADD_LYRIC         = "AddLyric"
)

var throttle chan (int)

func get(function string, query string, t ...int) (clr ChartLyricsResult, err error) {

    var throttle_enabled = len(t) > 0 && t[0] > 0

	if throttle_enabled {
        if throttle == nil {
            throttle = make(chan int)
        } else {
            <-throttle
        }
    }

	resp, err := http.Get(fmt.Sprintf(addr, function, query))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = xml.Unmarshal(b, &clr)
	if err != nil {
		return
	}

    if throttle_enabled {
        go func() {
            time.Sleep(time.Duration(int64(t[0]) * int64(time.Second)))
            throttle <- 1
        }()
    }

	cleanupSearchLyricResult(&clr.SearchLyricResult)

	return
}

func HttpSearchLyric(args string, t ...int) (ChartLyricsResult, error) {
	return get(_SEARCH_LYRIC, args, t...)
}

func HttpSearchLyricDirect(args string, t ...int) (ChartLyricsResult, error) {
	return get(_SEARCH_LYRIC_DIRECT, args, t...)
}

func HttpGetLyric(args string, t ...int) (ChartLyricsResult, error) {
	return get(_GET_LYRIC, args, t...)
}
