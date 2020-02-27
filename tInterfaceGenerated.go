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

// TEnvironment is a read only environment
type TEnvironment interface {
	//Name returns the name of the environment
	Name() string
	//Qualifier returns the qualifier of the environment
	Qualifier() string
	//Description returns the description of the environment
	Description() string
	//QualifiedName returns the qualified of the environment
	QualifiedName() string
	//Platform returns the platform used to deploy environment
	Platform() TPlatform
	//HasVars returns true if the environment has defined vars
	HasVars() bool
	//Vars returns the environement vars
	Vars() map[string]interface{}
	//Orchestrator returns the orchestrator managing the environment nodes
	Orchestrator() TOrchestrator
	//HasProviders returns true if the environment has providers
	HasProviders() bool
	//Providers returns the environment providers
	Providers() map[string]TProvider
	//HasNodes returns true if the environment has nodes
	HasNodes() bool
	//Nodes returns the environment providers
	Nodes() map[string]TNodeSet
	//HasStacks returns true if the environment has stacks
	HasStacks() bool
	//Stacks returns the environment stacks
	Stacks() map[string]TStack
	//HasTasks returns true if the environment has tasks
	HasTasks() bool
	//Tasks returns the environment tasks
	Tasks() map[string]TTask
	//HasHooks returns true if the environment has hooks
	HasHooks() bool
	//Hooks returns the environment hooks
	Hooks() TEnvironmentHooks
	//HasInitHooks returns true if the environment has hooks before creating
	HasInitHooks() bool
	//HasCreateHooks returns true if the environment has hooks while creating
	HasCreateHooks() bool
	//HasInstallHooks returns true if the environment has hooks while installing
	HasInstallHooks() bool
	//HasDeployHooks returns true if the environment has hooks while deploying
	HasDeployHooks() bool
	//HasDestroyHooks returns true if the environment has hooks while destroying
	HasDestroyHooks() bool
	//HasTemplates returns true if the environment has defined templates
	HasTemplates() bool
	//Templates returns the environment templates
	Templates() []string
}

// TOrchestrator is a read only orchestrator
type TOrchestrator interface {
	//Parameters returns the orchestrator parameters
	Parameters() map[string]interface{}
	//EnvVars returns the orchestrator environment variables
	EnvVars() map[string]string
	//Component returns the orchestrator component
	Component() (TComponent, error)
}

// TProvider is a read only provider
type TProvider interface {
	//Name returns the name of the provider
	Name() string
	//Parameters returns the provider parameters
	Parameters() map[string]interface{}
	//EnvVars returns the provider environment variables
	EnvVars() map[string]string
	//Proxy returns the proxy definition applied to the provider
	Proxy() TProxy
	//Component returns the provider component
	Component() (TComponent, error)
}

// TProviderRef is a read only reference on a provider
type TProviderRef interface {
	//Provider returns the provider wherein the node should be deployed
	Provider() (TProvider, error)
}

// TNodeSet is a read only node set
type TNodeSet interface {
	//Name returns the name of the node set
	Name() string
	//Instances returns the number of nodes to create for this node set
	Instances() int
	//Provider returns the reference on the provider wherein the node should be deployed
	Provider() TProviderRef
	//HasHooks returns true if the node has hooks
	HasHooks() bool
	//Hooks returns the node hooks
	Hooks() TNodeHook
	//HasCreateHooks returns true if the node has hooks while creating
	HasCreateHooks() bool
	//HasLabels returns true if the node has defined labels
	HasLabels() bool
	//Labels returns the node labels
	Labels() map[string]string
}

// TStack is a read only stack
type TStack interface {
	//Name returns the name of the stack
	Name() string
	//Playbook returns the custom deployment playbook of the stack if any
	Playbook() string
	//Parameters returns the stack parameters
	Parameters() map[string]interface{}
	//EnvVars returns the stack environment variables
	EnvVars() map[string]string
	//HasHooks returns true if the stack has hooks
	HasHooks() bool
	//Hooks returns the stack hooks
	Hooks() TStackHooks
	//HasDeployHooks returns true if the stack has hooks while deploying
	HasDeployHooks() bool
	//Dependencies returns the stack dependencies
	Dependencies() TDependencies
	//HasCopies returns true if the stacks has copies
	HasCopies() bool
	//Copies returns the stack copies
	Copies() map[string]TCopy
}

// TStackRef is a read only reference on a stack
type TStackRef interface {
	//Stack returns the stack corresponding to the ref
	Stack() (TStack, error)
}

// TDependencies is a read only list of stack dependencies
type TDependencies interface {
	//HasDependencies returns true if there is dependencies
	HasDependencies() bool
	//Dependencies returns the references of stacks we depend on
	Dependencies() []TStackRef
}

// TCopy is a read only representation of files to be copied
type TCopy interface {
	//HasLabels returns true if the copy has defined labels
	HasLabels() bool
	//Labels returns the copy labels
	Labels() map[string]string
	//HasSources returns true if the copy has defined sources
	HasSources() bool
	//Sources returns the copy sources
	Sources() []string
}

// TTask is a read only task
type TTask interface {
	//Name returns the name of the task
	Name() string
	//Playbook returns the playbook linked to the task
	Playbook() string
	//Parameters returns the task parameters
	Parameters() map[string]interface{}
	//EnvVars returns the task environment variables
	EnvVars() map[string]string
	//HasHooks returns true if the task has hooks
	HasHooks() bool
	//Hooks returns the task hooks
	Hooks() TTaskHooks
}

// TTaskRef is a read only reference on a task
type TTaskRef interface {
	//Task returns the task corresponding to the ref
	Task() (TTask, error)
}

// TPlatform is a read only platform
type TPlatform interface {
	//Base returns the base location of the platform
	Base() TBase
	//Parent returns the parent used by the platform
	Parent() TComponent
	//HasComponents returns true if the platform has components
	HasComponents() bool
}

// TBase is a read only base location
type TBase interface {
	//URL returns the url where the base refers
	URL() TURL
}

// TURL is a read only ekara url
type TURL interface {
	//String returns the string representation of the whole url
	String() string
	//Scheme returns the url scheme
	Scheme() string
	//Path returns the url path
	Path() string
	//AsFilePath returns the path converted as a file path
	AsFilePath() string
	//Host returns the url host
	Host() string
}

// TProxy is a read only proxy configuration
type TProxy interface {
	//HTTP returns the proxy http definition
	HTTP() string
	//HTTPS returns the proxy https definition
	HTTPS() string
	//NoProxy returns the no proxy definition
	NoProxy() string
}

// TComponent is a read only component
type TComponent interface {
	//ID returns the name of the component
	ID() string
	//Repository returns the repository where the component is located
	Repository() TRepository
	//HasTemplates returns true if the component has defined templates
	HasTemplates() bool
	//Templates returns true if the component templates
	Templates() []string
}

// TRepository is a read only repository
type TRepository interface {
	//Scm returns the type of the source control management holding the repository
	Scm() string
	//URL returns the url where the repository is located
	URL() TURL
	//Ref returns the reference (tag,branch, ...) to use within the repository
	Ref() string
	//DescriptorName returns the name of the ekara descriptor for this repository
	DescriptorName() string
}

// TEnvironmentHooks is a read only representation of the hooks associated to an environment
type TEnvironmentHooks interface {
	//HasInit returns true if the hooks has tasks while before creating
	HasInit() bool
	//Init returns the initialization tasks
	Init() THook
	//HasCreate returns true if the hooks has tasks while creating
	HasCreate() bool
	//Create returns the creating tasks
	Create() THook
	//HasInstall returns true if the hooks has tasks while installing
	HasInstall() bool
	//Install returns the installing tasks
	Install() THook
	//HasDeploy returns true if the hooks has tasks while deploying
	HasDeploy() bool
	//Deploy returns the deploying tasks
	Deploy() THook
	//HasDestroy returns true if the hooks has tasks while destroying
	HasDestroy() bool
	//Destroy returns the destruction tasks
	Destroy() THook
}

// TStackHooks is a read only representation of the hooks associated to a stack
type TStackHooks interface {
	//HasDeploy returns true if the hooks has tasks while deploying
	HasDeploy() bool
	//Deploy returns the deploying tasks
	Deploy() THook
}

// TTaskHooks is a read only representation of the hooks associated to a task
type TTaskHooks interface {
	//HasExecute returns true if the hooks has tasks while executing
	HasExecute() bool
	//Execute returns the executing tasks
	Execute() THook
}

// TNodeHook is a read only representation of the hooks associated to a node
type TNodeHook interface {
	//HasCreate returns true if the hooks has tasks while creating
	HasCreate() bool
	//Create returns the creating tasks
	Create() THook
}

// THook is a read only hooks
type THook interface {
	//After returns the references of tasks to run after the hooked action
	After() []TTaskRef
	//Before returns the references of tasks to run before the hooked action
	Before() []TTaskRef
}
