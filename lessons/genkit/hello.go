// Hello Genkit: initialize the Google AI plugin, define a flow, and generate text.
//
// Needs a Gemini API key in the environment:
//
//	export GEMINI_API_KEY=your-key
//	go run ./lessons/genkit
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

func main() {
	ctx := context.Background()

	// Init registers plugin models. A bad config (for example a missing API key)
	// panics here instead of failing on the first request.
	g := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-flash-latest"),
	)

	// A flow is a traced Go function with typed input and output.
	jokeFlow := genkit.DefineFlow(g, "jokeFlow",
		func(ctx context.Context, topic string) (string, error) {
			resp, err := genkit.Generate(ctx, g,
				ai.WithSystem("You are a stand-up comedian. One line, no preamble."),
				ai.WithPrompt("Tell a joke about %s.", topic),
			)
			if err != nil {
				return "", err
			}
			return resp.Text(), nil
		})

	joke, err := jokeFlow.Run(ctx, "gophers")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(joke)
}
