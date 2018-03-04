package main

// Repositories is the group of repos that monorepo going to administrate
type Repositories struct {
	Repos []Repo `json:"repositories"`
}

// Repo the configuration with each repo that should be administrate by the monorepo
type Repo struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
}

func (r Repo) isValid() bool {
	if r.Name == "" || r.Repository == "" {
		return false
	}
	return true
}
