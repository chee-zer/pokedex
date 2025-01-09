package commands

type pokedexRaw struct {
	Name   string `json:"name"`
	Basexp int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Typee struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Pokemon struct {
	Name   string
	Basexp int
	Height int
	Weight int
	Stats  []string
	Types  []string
}

type Pokedex struct {
	Entry map[string]Pokemon
}

func cookData(dat pokedexRaw) Pokemon {
	var cookedData Pokemon
	cookedData.Name = dat.Name
	cookedData.Basexp = dat.Basexp
	cookedData.Height = dat.Height
	cookedData.Weight = dat.Weight
	for _, stat := range dat.Stats {
		cookedData.Stats = append(cookedData.Stats, stat.Stat.Name)
	}
	for _, ty := range dat.Types {
		cookedData.Stats = append(cookedData.Stats, ty.Typee.Name)
	}

	return cookedData
}

func (p *Pokedex) Add(name string, val Pokemon) error {

	p.Entry[name] = val
	return nil
}
