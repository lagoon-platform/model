package model

//*****************************************************************************
//
//     _         _           ____                           _           _
//    / \  _   _| |_ ___    / ___| ___ _ __   ___ _ __ __ _| |_ ___  __| |
//   / _ \| | | | __/ _ \  | |  _ / _ \ '_ \ / _ \ '__/ _` | __/ _ \/ _` |
//  / ___ \ |_| | || (_) | | |_| |  __/ | | |  __/ | | (_| | ||  __/ (_| |
// /_/   \_\__,_|\__\___/   \____|\___|_| |_|\___|_|  \__,_|\__\___|\__,_|
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// ----------------------------------------------------
// Implementation(s) of TURL
// ----------------------------------------------------

//TURLOnEkURLHolder is the struct containing the EkURL in order to implement TURL
type TURLOnEkURLHolder struct {
	h EkURL
}

//CreateTURLForEkURL returns an holder of EkURL implementing TURL
func CreateTURLForEkURL(o EkURL) TURLOnEkURLHolder {
	return TURLOnEkURLHolder{
		h: o,
	}
}

//String returns the string representation of the whole url
func (r TURLOnEkURLHolder) String() string {
	return r.h.String()
}

//Scheme returns the url scheme
func (r TURLOnEkURLHolder) Scheme() string {
	return r.h.Scheme()
}

//Path returns the url path
func (r TURLOnEkURLHolder) Path() string {
	return r.h.Path()
}

//AsFilePath returns the path converted as a file path
func (r TURLOnEkURLHolder) AsFilePath() string {
	return r.h.AsFilePath()
}

//Host returns the url host
func (r TURLOnEkURLHolder) Host() string {
	return r.h.Host()
}
