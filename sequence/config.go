package sequence

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)


// Edge represents an Edge on a Sequence Diagram
type Edge struct {
	From  string `yaml:"from"`
	To    string `yaml:"to"`
	Label string `yaml:"label"`
	Type  string `yaml:"type"`
}

type config struct {
	Title        string   `yaml:"title"`
	Partecipants []string `yaml:"partecipants"`
	Edges        []Edge   `yaml:"edges"`
}

// DiagramFromYaml create a Diagram from a YAML file 
func DiagramFromYaml(file string) (*Diagram, error) {
	var cfg config
	name := strings.TrimSuffix(file, ".yaml")

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	d := NewDiagram(name)

	d.SetTitle(cfg.Title)

	for _, partecipant := range cfg.Partecipants {
		d.AddParticipants(partecipant)
	}

	for _, edge := range cfg.Edges {
		switch edge.Type {
		case "->":
			err = d.AddDirectionalEdge(edge.From, edge.To, edge.Label)
		case "-":
			err = d.AddUndirectionalEdge(edge.From, edge.To, edge.Label)

		default:
			err = errors.New("edge type not valid")
		}
	}

	if err != nil {
		return nil, err
	}

	return d, nil
}
