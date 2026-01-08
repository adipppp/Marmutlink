package resolver

import "adipppp/Marmutlink/internal/resolver/src/enums"

type LoadResult interface {
	getLoadResultType() enums.LoadResultType
}

type Track struct {
	Encoded string    // The base64 encoded track data
	Info    TrackInfo // Info about the track
}

type TrackInfo struct {
	Identifier string // The track identifier
	Author     string // The track author
	Length     int64  // The track length in milliseconds
	Title      string // The track title
	URI        string // The track URI
	ArtworkURL string // The track artwork URL
}

type Playlist struct {
	Info   PlaylistInfo
	Tracks []Track
}

type PlaylistInfo struct {
	Identifier string // The playlist identifier
	Author     string // The playlist author
	Name       string // The name of the playlist
	URI        string // The playlist URI
	ArtworkURL string // The playlist artwork URL
}

type TrackLoadResult struct {
	LoadType enums.LoadResultType // The type of the result
	Data     Track                // The data of the result
}

func (t TrackLoadResult) getLoadResultType() enums.LoadResultType {
	return t.LoadType
}

type PlaylistLoadResult struct {
	LoadType enums.LoadResultType // The type of the result
	Data     Playlist             // The data of the result
}

func (p PlaylistLoadResult) getLoadResultType() enums.LoadResultType {
	return p.LoadType
}

type SearchLoadResult struct {
	LoadType enums.LoadResultType // The type of the result
	Data     []Track              // The data of the result
}

func (s SearchLoadResult) getLoadResultType() enums.LoadResultType {
	return s.LoadType
}

type EmptyLoadResult struct {
	LoadType enums.LoadResultType // The type of the result
}

func (e EmptyLoadResult) getLoadResultType() enums.LoadResultType {
	return e.LoadType
}

type ErrorLoadResult struct {
	LoadType enums.LoadResultType // The type of the result
	Message  string               // The error message
}

func (e ErrorLoadResult) getLoadResultType() enums.LoadResultType {
	return e.LoadType
}
