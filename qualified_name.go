package model

import (
	"errors"
	"regexp"
)

// QualifiedName The Qualified name of an environment.
// This name can be used to identify, using for example Tags or Labels, all the
// content created relatively to the environment on the infrastructure of the desired
// cloud provider.
type QualifiedName struct {
	name string
	r    Environment
}

//String returns the qualified name as string
func (qn QualifiedName) String() string {
	return qn.name
}

// IsAValidQualifier is the regular expression used to validate the qualified name
// of an environment
//
// The qualified name is the concatenation if the environment name and its qualifier
// separated by a "_"
//
// The name and the qualifier can contain only alphanumerical characters
var IsAValidQualifier = regexp.MustCompile(`^[a-zA-Z0-9_a-zA-Z0-90-9]+$`).MatchString

func (qn QualifiedName) validate() ValidationErrors {
	vErrs := ValidationErrors{}
	if !qn.ValidQualifiedName() {
		vErrs.addError(errors.New("the environment name or the qualifier contains a non alphanumeric character"), qn.r.location.appendPath("name|qualifier"))
	}
	return vErrs
}

//ValidQualifiedName returns true if the qualified name is valid
func (qn QualifiedName) ValidQualifiedName() bool {
	return IsAValidQualifier(qn.String())
}

// QualifiedName returns the concatenation of the environment name and qualifier
// separated by a "_".
// If the environment qualifier is not defined it will return just the name
func (r Environment) QualifiedName() QualifiedName {
	return qualify(r)
}

// QualifiedName returns the concatenation of the environment name and qualifier
// separated by a "_".
// If the environment qualifier is not defined it will return just the name
func (r yamlEnvironment) QualifiedName() QualifiedName {
	if len(r.Qualifier) == 0 {
		return QualifiedName{name: r.Name}
	}
	return QualifiedName{name: r.Name + "_" + r.Qualifier}

}

func qualify(r Environment) QualifiedName {

	if len(r.Qualifier) == 0 {
		return QualifiedName{r: r, name: r.Name}
	}
	return QualifiedName{r: r, name: r.Name + "_" + r.Qualifier}

}
