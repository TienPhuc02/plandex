package shared

import (
	openai "github.com/sashabaranov/go-openai"
)

type CurrentPlanFiles struct {
	Files map[string]string `json:"files"`
}

type ConversationMessage struct {
	Message       openai.ChatCompletionMessage `json:"message"`
	Tokens        int                          `json:"tokens"`
	Summarized    bool                         `json:"summarized"`
	Summary       string                       `json:"summary"`
	SummaryTokens int                          `json:"summaryTokens"`
}

type PromptRequest struct {
	Prompt           string                `json:"prompt"`
	ModelContext     ModelContext          `json:"modelContext"`
	CurrentPlan      CurrentPlanFiles      `json:"currentPlan"`
	Conversation     []ConversationMessage `json:"conversation"`
	ParentProposalId string                `json:"parentProposalId"`
}

type SummarizeRequest struct {
	Text string `json:"text"`
}

type SummarizeResponse struct {
	Name     string `json:"name"`
	Summary  string `json:"summary"`
	FileName string `json:"fileName"`
}

type ModelContextPart struct {
	Name      string `json:"name"`
	Summary   string `json:"summary"`
	Body      string `json:"body"`
	Url       string `json:"url"`
	FilePath  string `json:"filePath"`
	Sha       string `json:"sha"`
	NumTokens int    `json:"numTokens"`
	UpdatedAt string `json:"updatedAt"`
}
type ModelContext []ModelContextPart

type ModelContextState struct {
	NumTokens    int    `json:"numTokens"`
	Counter      uint32 `json:"counter"`
	ActiveTokens int    `json:"activeTokens"`
	ChatFlexPct  int    `json:"chatFlexPct"`
	PlanFlexPct  int    `json:"planFlexPct"`
}

type PlanSettings struct {
	Name string `json:"name"`
}

type PlanChunk struct {
	Type     string `json:"type"`
	FilePath string `json:"filePath"`
	Content  string `json:"content"`
}

type PlanDescription struct {
	MadePlan  bool     `json:"madePlan"`
	CommitMsg string   `json:"commitMsg"`
	Files     []string `json:"files"`
}