package prompt

import (
	"bytes"
	"context"
	"html/template"

	"github.com/gofrs/uuid"
	"github.com/missingstudio/studio/backend/models"
)

type Repository interface {
	GetAll(context.Context) ([]models.Prompt, error)
	Upsert(context.Context, models.Prompt) (models.Prompt, error)
	GetByID(context.Context, uuid.UUID) (models.Prompt, error)
	GetByName(context.Context, string) (models.Prompt, error)
	DeleteByID(context.Context, uuid.UUID) error
}

type Prompt struct {
	tmpl *template.Template
	data map[string]any
}

func NewPrompt(text string, data map[string]any) *Prompt {
	return &Prompt{
		tmpl: template.Must(template.New("prompt").Parse(text)),
		data: data,
	}
}

func (p *Prompt) Run() (string, error) {
	var buf bytes.Buffer
	err := p.tmpl.Execute(&buf, p.data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
