package filter

import (
	"strings"

	"github.com/mbrt/gmailfilter/pkg/config"
)

// Filters is a list of filters created in Gmail.
type Filters []Filter

func (fs Filters) String() string {
	builder := strings.Builder{}
	for _, f := range fs {
		builder.WriteString(f.String())
		builder.WriteRune('\n')
	}
	return builder.String()
}

// Filter matches 1:1 a filter created on Gmail.
type Filter struct {
	// ID is an optional identifier associated with a filter.
	ID       string
	Action   Action
	Criteria Criteria
}

func (f Filter) String() string {
	builder := strings.Builder{}

	builder.WriteString("- Criteria:\n")
	writeParam(&builder, "from", f.Criteria.From)
	writeParam(&builder, "to", f.Criteria.To)
	writeParam(&builder, "subject", f.Criteria.Subject)
	writeParam(&builder, "query", f.Criteria.Query)

	builder.WriteString("  Actions:\n")
	writeBool(&builder, "archive", f.Action.Archive)
	writeBool(&builder, "delete", f.Action.Delete)
	writeBool(&builder, "mark as important", f.Action.MarkImportant)
	writeBool(&builder, "mark as read", f.Action.MarkImportant)
	writeParam(&builder, "categorize as", string(f.Action.Category))
	writeParam(&builder, "apply label", f.Action.AddLabel)

	return builder.String()
}

func writeParam(b *strings.Builder, name, value string) {
	if value != "" {
		b.WriteString("    ")
		b.WriteString(name)
		b.WriteString(": ")
		b.WriteString(value)
		b.WriteRune('\n')
	}
}

func writeBool(b *strings.Builder, name string, value bool) {
	if value {
		b.WriteString("    ")
		b.WriteString(name)
		b.WriteRune('\n')
	}
}

// Action represents an action associated with a Gmail filter.
type Action struct {
	Archive       bool
	Delete        bool
	MarkImportant bool
	MarkRead      bool
	Category      config.Category
	AddLabel      string
}

// Empty returns true if no action is specified.
func (a Action) Empty() bool {
	return a == Action{}
}

// Criteria represents the filtering criteria associated with a Gmail filter.
type Criteria struct {
	From    string
	To      string
	Subject string
	Query   string
}

// Empty returns true if no criteria is specified.
func (c Criteria) Empty() bool {
	return c == Criteria{}
}

// Label contains information about a Gmail label.
type Label struct {
	ID   string
	Name string
}
