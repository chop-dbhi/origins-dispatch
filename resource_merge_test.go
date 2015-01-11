package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/chop-dbhi/origins-dispatch/graph"
)

const neo4jTestUri = "http://localhost:7500/db/data/"

const resourceMergeData = `{
	"resource": {"origins:id": "test", "prov:label": "Test"},
	"operations": [
		{"command": "add", "model": "Entity"},
		{"command": "add", "model": "Entity"},
		{"command": "update", "model": "Entity"},
		{"command": "remove", "model": "Entity"},
		{"command": "add", "model": "Agent"},
		{"command": "add", "model": "Activity"}
	]
}`

var c *graph.Client

func init() {
	c = &graph.Client{
		Uri: neo4jTestUri,
	}
}

// Populate graph with test data
func populate(c *graph.Client) {
	c.ExecuteStatement("CREATE (r:Resource {`origins:id`: 'test', `prov:label`: 'Test'}), (u1:User {email: 'user1@example.com'}), (u2:User {email: 'user2@example.com'}), (:User {email: 'user3@example.com'}), (u1)-[:subscribedTo]->(r), (u2)-[:subscribedTo]->(r)")
}

func createResourceMergePayload(t *testing.T) *ResourceMergePayload {
	raw := json.RawMessage(resourceMergeData)

	e := &EventPayload{
		Event: "resource-merge",
		Data:  &raw,
	}

	p := ResourceMergePayload{}

	// Unmarshal the remaining bytes
	if err := json.Unmarshal(*e.Data, &p); err != nil {
		t.Error(err)
	}

	return &p
}

func TestResource(t *testing.T) {
	// Delete data first
	c.ExecuteStatement("MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE r, n")

	populate(c)

	r := &Resource{
		Id:   "test",
		Name: "Test",
	}

	users, _ := r.Subscribers(c)

	assert.Equal(t, len(users), 2)
}

func TestHandleResourceMergePayload(t *testing.T) {
	p := createResourceMergePayload(t)

	assert.Equal(t, len(p.Operations), 6)

	stats := p.OpStats()

	if _, ok := stats["add"]; !ok {
		t.Error("op stats failed")
	}
}

func TestResourceMergeEmailMessage(t *testing.T) {
	p := createResourceMergePayload(t)
	stats := p.OpStats()

	subject, _ := renderTemplate(resourceMergeSubject, p)
	message, _ := renderTemplate(resourceMergeMessage, stats)

	assert.Equal(t, string(subject), "[Test] Changes")
	assert.Equal(t, strings.TrimSpace(string(message)), `- add
	- Activity: 1
	- Agent: 1
	- Entity: 2
- remove
	- Entity: 1
- update
	- Entity: 1`)
}
