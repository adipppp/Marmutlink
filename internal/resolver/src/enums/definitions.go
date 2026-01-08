package enums

type LoadResultType int

const (
	Track = iota
	Playlist
	Search
	Empty
	Error
)

func (t LoadResultType) String() string {
	switch t {
	case Track:
		return "TRACK"
	case Playlist:
		return "PLAYLIST"
	case Search:
		return "SEARCH"
	case Empty:
		return "EMPTY"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
