package main

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/chop-dbhi/origins-dispatch/graph"
	"github.com/jmcvetta/neoism"
)

var (
	resourceMergeSubject *template.Template
	resourceMergeMessage *template.Template
)

const (
	resourceSubscribers      = "MATCH (:Resource {`origins:id`: { id }})<-[:subscribedTo]-(n:User) RETURN n"
	resourceMergeSubjectText = "[{{.Resource.Name}}] Changes"
	resourceMergeMessageText = `
{{range $cmd, $models := .}}- {{$cmd}}{{range $model, $count := $models}}
	- {{$model}}: {{$count}}{{end}}
{{end}}
`
)

func init() {
	resourceMergeSubject = compileTemplate("resourceMergeSubject", resourceMergeSubjectText)
	resourceMergeMessage = compileTemplate("resourceMergeMessage", resourceMergeMessageText)
}

// User corresponds to an Origins user account with an email address.
type User struct {
	Email string `json:"email"`
}

// Users is an array of users.
type Users []*User

// Intermediate result for query
type result []struct {
	N *User `json:"n"`
}

func (u *Users) UnmarshalJSON(b []byte) error {
	res := result{}

	err := json.Unmarshal(b, &res)

	if err != nil {
		return err
	}

	if len(res) > 0 {
		t := make([]*User, len(res))

		for i, x := range res {
			t[i] = x.N
		}

		*u = t
	}

	return nil
}

// Resource is an Origins resource.
type Resource struct {
	Name string `json:"prov:label"`
	Id   string `json:"origins:id"`
}

// Subscribers returns an array of users who subscribe to the resource
// for notifications.
func (r *Resource) Subscribers(c *graph.Client) (Users, error) {
	u := Users{}

	q := neoism.CypherQuery{
		Statement: resourceSubscribers,
		Parameters: neoism.Props{
			"id": r.Id,
		},
		Result: &u,
	}

	if err := c.Execute(&q); err != nil {
		return nil, err
	}

	return u, nil
}

type Operation struct {
	Command  string                 `json:"command"`
	Model    string                 `json:"model"`
	Instance map[string]interface{} `json:"instance"`
	Previous map[string]interface{} `json:"previous"`
	Diff     map[string]interface{} `json:"diff"`
}

type ResourceMergePayload struct {
	Resource   *Resource    `json:"resource"`
	Operations []*Operation `json:"operations"`
}

func (r *ResourceMergePayload) OpStats() map[string]map[string]int {
	a := make(map[string]map[string]int)

	for _, op := range r.Operations {
		if _, ok := a[op.Command]; !ok {
			a[op.Command] = make(map[string]int)
		}

		a[op.Command][op.Model] += 1
	}

	return a
}

func compileTemplate(name string, text string) *template.Template {
	return template.Must(template.New(name).Parse(text))
}

func renderTemplate(t *template.Template, d interface{}) ([]byte, error) {
	b := &bytes.Buffer{}

	if err := t.Execute(b, d); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// Handles the resource merge event
func handleResourceMerge(e *EventPayload) error {
	p := ResourceMergePayload{}

	// Unmarshal the remaining bytes
	if err := json.Unmarshal(*e.Data, &p); err != nil {
		return err
	}

	c := &graph.Client{
		Uri: serveNeo4j,
	}

	users, err := p.Resource.Subscribers(c)

	if err != nil {
		return err
	}

	if len(users) == 0 {
		return nil
	}

	to := make([]string, len(users))

	for i, u := range users {
		to[i] = u.Email
	}

	stats := p.OpStats()

	subject, err := renderTemplate(resourceMergeSubject, p)
	msg, err := renderTemplate(resourceMergeMessage, stats)

	return sendEmail(to, string(subject), msg)
}
