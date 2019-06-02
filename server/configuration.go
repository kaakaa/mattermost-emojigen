package main

type configuration struct {
	AccessToken string
}

func (p *EmojigenPlugin) getConfiguration() *configuration {
	p.configurationLock.RLock()
	defer p.configurationLock.RUnlock()

	if p.configuration == nil {
		return &configuration{}
	}
	return p.configuration
}

func (p *EmojigenPlugin) setConfiguration(configuration *configuration) {
	p.configurationLock.Lock()
	defer p.configurationLock.Unlock()
	p.configuration = configuration
}
