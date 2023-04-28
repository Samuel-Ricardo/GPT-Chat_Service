package configs

type conf struct {
	DBDriver           string   `mapstructure:"DB_DRIVER"`
	DBHost             string   `mapstructure:"DB_HOST"`
	DBPort             string   `mapstructure:"DB_PORT"`
	DBUser             string   `mapstructure:"DB_USER"`
	DBPassword         string   `mapstructure:"DB_PASSWORD"`
	DBName             string   `mapstructure:"DB_NAME"`
	WebServerPort      string   `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort     string   `mapstructure:"GRPC_SERVER_PORT"`
	InitialChatMessage string   `mapstructure:"INITIAL_CHAT_MESSAGE"`
	OpenAIApiKey       string   `mapstructure:"OPENAI_API_KEY"`
	Model              string   `mapstructure:"MODEL"`
	ModelMaxTokens     int      `mapstructure:"MODEL_MAX_TOKENS"`
	Temperature        float64  `mapstructure:"TEMPERATURE"`
	TopP               float64  `mapstructure:"TOP_P"`
	N                  int      `mapstructure:"N"`
	Stop               []string `mapstructure:"STOP"`
	MaxTokens          int      `mapstructure:"MAX_TOKENS"`
	AuthToken          string   `mapstructure:"AUTH_TOKEN"`
}

