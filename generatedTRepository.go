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
// Implementation(s) of TRepository
// ----------------------------------------------------

//TRepositoryOnRepositoryHolder is the struct containing the Repository in order to implement TRepository
type TRepositoryOnRepositoryHolder struct {
	h Repository
}

//CreateTRepositoryForRepository returns an holder of Repository implementing TRepository
func CreateTRepositoryForRepository(o Repository) TRepositoryOnRepositoryHolder {
	return TRepositoryOnRepositoryHolder{
		h: o,
	}
}

//Scm returns the type of the source control management holding the repository
func (r TRepositoryOnRepositoryHolder) Scm() string {
	return string(r.h.Scm)
}

//URL returns the url where the repository is located
func (r TRepositoryOnRepositoryHolder) URL() TURL {
	return CreateTURLForEkURL(r.h.Url)
}

//Ref returns the reference (tag,branch, ...) to use within the repository
func (r TRepositoryOnRepositoryHolder) Ref() string {
	return r.h.Ref
}

//DescriptorName returns the name of the ekara descriptor for this repository
func (r TRepositoryOnRepositoryHolder) DescriptorName() string {
	return r.h.DescriptorName
}
