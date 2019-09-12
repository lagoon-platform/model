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
// Implementation(s) of TStack
// ----------------------------------------------------

//TStackOnStackHolder is the struct containing the Stack in order to implement TStack
type TStackOnStackHolder struct {
	h Stack
}

//CreateTStackForStack returns an holder of Stack implementing TStack
func CreateTStackForStack(o Stack) TStackOnStackHolder {
	return TStackOnStackHolder{
		h: o,
	}
}

//Name returns the name of the stack
func (r TStackOnStackHolder) Name() string {
	return r.h.Name
}

//Parameters returns the stack parameters
func (r TStackOnStackHolder) Parameters() map[string]interface{} {
	return r.h.Parameters
}

//EnvVars returns the stack environment variables
func (r TStackOnStackHolder) EnvVars() map[string]string {
	return r.h.EnvVars
}

//HasHooks returns true if the stack has hooks
func (r TStackOnStackHolder) HasHooks() bool {
	return r.h.Hooks.HasTasks()
}

//Hooks returns the stack hooks
func (r TStackOnStackHolder) Hooks() TStackHooks {
	return CreateTStackHooksForStackHook(r.h.Hooks)
}

//HasDeployHooks returns true if the stack has hooks while deploying
func (r TStackOnStackHolder) HasDeployHooks() bool {
	return r.h.Hooks.Deploy.HasTasks()
}

//Dependencies returns the stack dependencies
func (r TStackOnStackHolder) Dependencies() TDependencies {
	return CreateTDependenciesForDependencies(r.h.DependsOn)
}

//HasCopies returns true if the stacks has copies
func (r TStackOnStackHolder) HasCopies() bool {
	return len(r.h.Copies.Content) > 0
}

//Copies returns the stack copies
func (r TStackOnStackHolder) Copies() map[string]TCopy {
	result := make(map[string]TCopy)
	for k, val := range r.h.Copies.Content {
		result[k] = CreateTCopyForCopy(val)
	}
	return result

}
