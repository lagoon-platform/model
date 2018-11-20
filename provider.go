package model

import (
	"encoding/json"
	"errors"
)

type (
	Provider struct {
		// The Name of the provider
		Name string
		// The component containing the provider
		Component ComponentRef
		// The provider parameters
		Parameters Parameters
		// The provider environment variables
		EnvVars EnvVars
		// The provider proxy
		Proxy Proxy
	}

	Providers map[string]Provider

	// Reference to a provider
	ProviderRef struct {
		ref        string
		parameters Parameters
		envVars    EnvVars
		proxy      Proxy
		env        *Environment
		location   DescriptorLocation
	}
)

func (r Provider) DescType() string {
	return "Provider"
}

func (r Provider) DescName() string {
	return r.Name
}

func (r Provider) MarshalJSON() ([]byte, error) {
	component, e := r.Component.Resolve()
	if e != nil {
		return nil, e
	}
	return json.Marshal(struct {
		Name       string     `json:",omitempty"`
		Component  string     `json:",omitempty"`
		Parameters Parameters `json:",omitempty"`
		EnvVars    EnvVars    `json:",omitempty"`
		Proxy      Proxy      `json:",omitempty"`
	}{
		Name:       r.Name,
		Component:  component.Id,
		Parameters: r.Parameters,
		EnvVars:    r.EnvVars,
		Proxy:      r.Proxy,
	})
}

func (r Provider) valid(location DescriptorLocation) ValidationErrors {
	return r.Component.validate()
}

func (r Provider) validate() ValidationErrors {
	return r.Component.validate()
}

func (r *Provider) merge(other Provider) error {
	if r.Name != other.Name {
		return errors.New("cannot merge unrelated providers (" + r.Name + " != " + other.Name + ")")
	}
	if err := r.Component.merge(other.Component); err != nil {
		return err
	}
	r.Parameters = r.Parameters.inherits(other.Parameters)
	r.EnvVars = r.EnvVars.inherits(other.EnvVars)
	r.Proxy = r.Proxy.inherits(other.Proxy)
	return nil
}

// createProviders creates all the providers declared into the provided environment
func createProviders(env *Environment, yamlEnv *yamlEnvironment) Providers {
	res := Providers{}
	for name, yamlProvider := range yamlEnv.Providers {
		res[name] = Provider{
			Name:       name,
			Component:  createComponentRef(env, env.location.appendPath("providers."+name), yamlProvider.Component, true),
			Parameters: createParameters(yamlProvider.Params),
			Proxy:      createProxy(yamlProvider.Proxy),
			EnvVars:    createEnvVars(yamlProvider.Env)}
	}
	return res
}

func (r Providers) merge(env *Environment, other Providers) error {
	for id, p := range other {
		if provider, ok := r[id]; ok {
			if err := provider.merge(p); err != nil {
				return err
			}
		} else {
			p.Component.env = env
			r[id] = p
		}
	}
	return nil
}

func (r ProviderRef) Resolve() (Provider, error) {
	validationErrors := r.validate()
	if validationErrors.HasErrors() {
		return Provider{}, validationErrors
	}
	provider := r.env.Providers[r.ref]
	return Provider{
		Name:       provider.Name,
		Component:  provider.Component,
		Parameters: r.parameters.inherits(provider.Parameters),
		EnvVars:    r.envVars.inherits(provider.EnvVars),
		Proxy:      r.proxy.inherits(provider.Proxy)}, nil
}

// createProviderRef creates a reference to the provider declared into the yaml reference
func createProviderRef(env *Environment, location DescriptorLocation, yamlRef yamlProviderRef) ProviderRef {
	return ProviderRef{
		env:        env,
		ref:        yamlRef.Name,
		parameters: createParameters(yamlRef.Params),
		proxy:      createProxy(yamlRef.Proxy),
		envVars:    createEnvVars(yamlRef.Env),
		location:   location}
}

func (r ProviderRef) validate() ValidationErrors {
	validationErrors := ValidationErrors{}
	if len(r.ref) == 0 {
		validationErrors.addError(errors.New("empty provider reference"), r.location)
	} else {
		if _, ok := r.env.Providers[r.ref]; !ok {
			validationErrors.addError(errors.New("reference to unknown provider: "+r.ref), r.location)
		}
	}
	return validationErrors
}

func (r *ProviderRef) merge(other ProviderRef) error {
	if r.ref == "" {
		r.ref = other.ref
	}
	r.parameters = r.parameters.inherits(other.parameters)
	r.envVars = r.envVars.inherits(other.envVars)
	r.proxy = r.proxy.inherits(other.proxy)
	return nil
}
