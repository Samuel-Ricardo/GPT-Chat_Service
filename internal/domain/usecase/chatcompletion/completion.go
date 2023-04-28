package chatcompletion

type ChatCompletionConfigInputDTO struct {
	Model                string
	ModelMaxTokens       int
	Temperature          float32  // 0.0 to 1.0
	TopP                 float32  // 0.0 to 1.0 - to a low value, like 0.1, the model will be very conservative in its word choices, and will tend to generate relatively predictable prompts
	N                    int      // number of messages to generate
	Stop                 []string // list of tokens to stop on
	MaxTokens            int      // number of tokens to generate
	PresencePenalty      float32  // -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	FrequencyPenalty     float32  // -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, increasing the model's likelihood to talk about new topics.
	InitialSystemMessage string
}

type ChatCompletionInputDTO struct {
  ChatID string `json:"chat_id,omitempty"`
  UserID string `json:"user_id"`
  UserMessage string `json:"user_message"`
  Config ChatCompletionConfigInputDTO `json:"Config"`
}
