package theme

import (
	"log"
	"os"

	"github.com/adrg/xdg"
	kyaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"gopkg.in/yaml.v2"
)

type Theme struct {
	ClientConfig ClientConfigColors `yaml:"clientConfig,omitempty"`

	MainHeader HeaderColorSet `yaml:"mainHeader,omitempty"`

	Table TableColors `yaml:"table,omitempty"`
	Tab   TabColors   `yaml:"tab,omitempty"`

	Paginator PaginatorColorSet `yaml:"paginator,omitempty"`
	SyncBar   SyncBarColors     `yaml:"syncBar,omitempty"`
	Help      HelpColorSet      `yaml:"help,omitempty"`
	StatusLog StatusLogColors   `yaml:"statusLog,omitempty"`
	StatusBar StatusBarColors   `yaml:"statusBar,omitempty"`
}

type TextColorSet struct {
	BG  string `yaml:"bg,omitempty"`
	TXT string `yaml:"txt,omitempty"`
}

type HeaderColorSet struct {
	Title        TextColorSet `yaml:"title,omitempty"`
	ItemCount    TextColorSet `yaml:"itemCount,omitempty"`
	FilterPrompt TextColorSet `yaml:"filterPrompt,omitempty"`
	FilterCursor TextColorSet `yaml:"filterCursor,omitempty"`
}

type HelpColorSet struct {
	ShortcutName        TextColorSet `yaml:"shortcutName,omitempty"`
	ShortcutDescription TextColorSet `yaml:"shortcutDescription,omitempty"`
	ShortcutSeparator   TextColorSet `yaml:"shortcutSeparator,omitempty"`
}

type PaginatorColorSet struct {
	Active   string `yaml:"active,omitempty"`
	Inactive string `yaml:"inactive,omitempty"`
	Dimmed   string `yaml:"dimmed,omitempty"`
}

type ClientConfigColors struct {
	Header        HeaderColorSet `yaml:"header,omitempty"`
	StatusMessage TextColorSet   `yaml:"statusMessage,omitempty"`

	ActiveSelectionTitle TextColorSet `yaml:"activeSelectionTitle,omitempty"`
	NormalTitle          TextColorSet `yaml:"normalTitle,omitempty"`
	DimmedTitle          TextColorSet `yaml:"dimmedTitle,omitempty"`

	ActiveSelectionDescription TextColorSet `yaml:"activeSelectionDescription,omitempty"`
	NormalDescription          TextColorSet `yaml:"normalDescription,omitempty"`
	DimmedDescription          TextColorSet `yaml:"dimmedDescription,omitempty"`
}

type TableColors struct {
	ActiveBorder string `yaml:"activeBorder,omitempty"`
	DimmedBorder string `yaml:"dimmedBorder,omitempty"`

	ActiveCell TextColorSet `yaml:"activeCell,omitempty"`
	DimmedCell TextColorSet `yaml:"dimmedCell,omitempty"`

	ActiveHeader TextColorSet `yaml:"activeHeader,omitempty"`
	DimmedHeader TextColorSet `yaml:"dimmedHeader,omitempty"`

	ActiveSelected TextColorSet `yaml:"activeSelected,omitempty"`
	DimmedSelected TextColorSet `yaml:"dimmedSelected,omitempty"`
}

type TabColors struct {
	ActiveTabBorder         string `yaml:"activeTabBorder,omitempty"`
	InactiveTabBorder       string `yaml:"inactiveTabBorder,omitempty"`
	DimmedActiveTabBorder   string `yaml:"dimmedActiveTabBorder,omitempty"`
	DimmedInactiveTabBorder string `yaml:"dimmedInactiveTabBorder,omitempty"`

	ActiveTab         TextColorSet `yaml:"activeTab,omitempty"`
	InactiveTab       TextColorSet `yaml:"inactiveTab,omitempty"`
	DimmedActiveTab   TextColorSet `yaml:"dimmedActiveTab,omitempty"`
	DimmedInactiveTab TextColorSet `yaml:"dimmedInactiveTab,omitempty"`

	ActiveTabContentKeys       string `yaml:"activeTabContentKeys,omitempty"`
	DimmedActiveTabContentKeys string `yaml:"dimmedActiveTabContentKeys,omitempty"`

	ActiveTabContentValues       string `yaml:"activeTabContentValues,omitempty"`
	DimmedActiveTabContentValues string `yaml:"dimmedActiveTabContentValues,omitempty"`
}

type SyncBarColors struct {
	Spinner        string       `yaml:"spinner,omitempty"`
	InSyncState    TextColorSet `yaml:"inSyncState,omitempty"`
	NotSyncedState TextColorSet `yaml:"notSyncedState,omitempty"`
	StartingState  TextColorSet `yaml:"startingState,omitempty"`
	PausedState    TextColorSet `yaml:"pausedState,omitempty"`
}

type StatusLogColors struct {
	Duration  TextColorSet `yaml:"duration,omitempty"`
	Timestamp TextColorSet `yaml:"timestamp,omitempty"`

	OKStatus   TextColorSet `yaml:"okStatus,omitempty"`
	NOKStatus  TextColorSet `yaml:"nokStatus,omitempty"`
	WarnStatus TextColorSet `yaml:"warnStatus,omitempty"`

	InfoLevel  TextColorSet `yaml:"infoLevel,omitempty"`
	WarnLevel  TextColorSet `yaml:"warnLevel,omitempty"`
	ErrorLevel TextColorSet `yaml:"errorLevel,omitempty"`
	DebugLevel TextColorSet `yaml:"debugLevel,omitempty"`
}

type StatusBarColors struct {
	ResourceSection        TextColorSet `yaml:"resourceSection,omitempty"`
	ResourceDetailsSection TextColorSet `yaml:"resourceDetailsSection,omitempty"`
	ContextSection         TextColorSet `yaml:"contextSection,omitempty"`
	NamespaceSection       TextColorSet `yaml:"namespaceSection,omitempty"`
}

var Selected Theme

func (t *Theme) Marshal() []byte {
	out, err := yaml.Marshal(t)
	if err != nil {
		log.Fatalf("failed to marshal theme: %s", err)
	}
	return out
}

func createConfig(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	_, err = file.Write(defaultTheme.Marshal())
	return err
}

func loadConfig(filePath string) error {
	var k = koanf.New(".")
	var parser = kyaml.Parser()

	if err := k.Load(file.Provider(filePath), parser); err != nil {
		if err := createConfig(filePath); err != nil {
			return err
		}
		if err := loadConfig(filePath); err != nil {
			return err
		}
	}

	if err := k.Unmarshal("", &Selected); err != nil {
		return err
	}
	return nil
}

func init() {
	themeConfigFilePath, _ := xdg.ConfigFile("kibe/theme.yaml")

	if err := loadConfig(themeConfigFilePath); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}
