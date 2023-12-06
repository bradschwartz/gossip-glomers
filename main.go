package main

import (
	"encoding/json"
	"fmt"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var (
	id       = 0
	messages = make(map[float64]any)
)

func main() {
	node := maelstrom.NewNode()

	node.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "echo_ok"

		// Echo the original message back with the updated message type.
		return node.Reply(msg, body)
	})

	node.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"
		body["id"] = fmt.Sprintf("%s-%d", node.ID(), id)
		id++

		return node.Reply(msg, body)
	})

	node.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		message := body["message"].(float64)
		messages[message] = true

		respone := map[string]any{"type": "broadcast_ok"}
		return node.Reply(msg, respone)
	})

	node.Handle("read", func(msg maelstrom.Message) error {
		// Don't even bother serializing the message body since we don't use it.
		body := map[string]any{}
		body["type"] = "read_ok"
		keys := make([]float64, 0, len(messages))
		for k := range messages {
			keys = append(keys, k)
		}
		body["messages"] = keys

		return node.Reply(msg, body)
	})

	node.Handle("topology", func(msg maelstrom.Message) error {
		var body = map[string]string{"type": "topology_ok"}

		return node.Reply(msg, body)
	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}
