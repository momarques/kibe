package style

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
	List ListColors `yaml:"list"`

	MainHeader HeaderColorSet `yaml:"mainHeader"`

	Table TableColors `yaml:"table"`
	Tab   TabColors   `yaml:"tab"`

	SyncBar   SyncBarColors   `yaml:"syncBar"`
	Help      HelpColorSet    `yaml:"help"`
	StatusLog StatusLogColors `yaml:"statusLog"`
	StatusBar StatusBarColors `yaml:"statusBar"`
}

type TextColorSet struct {
	BG  string `yaml:"bg"`
	TXT string `yaml:"txt"`
}

type HeaderColorSet struct {
	Title        TextColorSet `yaml:"title"`
	ItemCount    TextColorSet `yaml:"itemCount"`
	FilterPrompt TextColorSet `yaml:"filterPrompt"`
	FilterCursor TextColorSet `yaml:"filterCursor"`
}

type HelpColorSet struct {
	ShortcutName        TextColorSet `yaml:"shortcutName"`
	ShortcutDescription TextColorSet `yaml:"shortcutDescription"`
	ShortcutSeparator   TextColorSet `yaml:"shortcutSeparator"`
}

type PaginatorColorSet struct {
	Active   string `yaml:"active"`
	Inactive string `yaml:"inactive"`
	Dimmed   string `yaml:"dimmed"`
}

type ListColors struct {
	Header        HeaderColorSet `yaml:"header"`
	StatusMessage TextColorSet   `yaml:"statusMessage"`

	ActiveSelectionTitle TextColorSet `yaml:"activeSelectionTitle"`
	NormalTitle          TextColorSet `yaml:"normalTitle"`
	DimmedTitle          TextColorSet `yaml:"dimmedTitle"`

	ActiveSelectionDescription TextColorSet `yaml:"activeSelectionDescription"`
	NormalDescription          TextColorSet `yaml:"normalDescription"`
	DimmedDescription          TextColorSet `yaml:"dimmedDescription"`
}

type TableColors struct {
	ActiveBorder string `yaml:"activeBorder"`
	DimmedBorder string `yaml:"dimmedBorder"`

	ActiveCell TextColorSet `yaml:"activeCell"`
	DimmedCell TextColorSet `yaml:"dimmedCell"`

	ActiveHeader TextColorSet `yaml:"activeHeader"`
	DimmedHeader TextColorSet `yaml:"dimmedHeader"`

	ActiveSelected TextColorSet `yaml:"activeSelected"`
	DimmedSelected TextColorSet `yaml:"dimmedSelected"`

	Paginator PaginatorColorSet `yaml:"paginator"`
}

type TabColors struct {
	ActiveTabBorder         string `yaml:"activeTabBorder"`
	InactiveTabBorder       string `yaml:"inactiveTabBorder"`
	DimmedActiveTabBorder   string `yaml:"dimmedActiveTabBorder"`
	DimmedInactiveTabBorder string `yaml:"dimmedInactiveTabBorder"`

	ActiveTab         TextColorSet `yaml:"activeTab"`
	InactiveTab       TextColorSet `yaml:"inactiveTab"`
	DimmedActiveTab   TextColorSet `yaml:"dimmedActiveTab"`
	DimmedInactiveTab TextColorSet `yaml:"dimmedInactiveTab"`

	ActiveTabContentKeys       string `yaml:"activeTabContentKeys"`
	DimmedActiveTabContentKeys string `yaml:"dimmedActiveTabContentKeys"`

	ActiveTabContentValues       string `yaml:"activeTabContentValues"`
	DimmedActiveTabContentValues string `yaml:"dimmedActiveTabContentValues"`

	Paginator PaginatorColorSet `yaml:"paginator"`
}

type SyncBarColors struct {
	Spinner       string       `yaml:"spinner"`
	InSyncState   TextColorSet `yaml:"inSyncState"`
	UnsyncedState TextColorSet `yaml:"unsyncedState"`
	StartingState TextColorSet `yaml:"startingState"`
}

type StatusLogColors struct {
	Duration  TextColorSet `yaml:"duration"`
	Timestamp TextColorSet `yaml:"timestamp"`

	OKStatus   TextColorSet `yaml:"okStatus"`
	NOKStatus  TextColorSet `yaml:"nokStatus"`
	WarnStatus TextColorSet `yaml:"warnStatus"`

	InfoLevel  TextColorSet `yaml:"infoLevel"`
	WarnLevel  TextColorSet `yaml:"warnLevel"`
	ErrorLevel TextColorSet `yaml:"errorLevel"`
	DebugLevel TextColorSet `yaml:"debugLevel"`
}

type StatusBarColors struct {
	ResourceSection        TextColorSet `yaml:"resourceSection"`
	ResourceDetailsSection TextColorSet `yaml:"resourceDetailsSection"`
	ContextSection         TextColorSet `yaml:"contextSection"`
	NamespaceSection       TextColorSet `yaml:"namespaceSection"`
}

var ThemeConfig Theme

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
	_, err = file.Write(defaultThemeConfig.Marshal())
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

	if err := k.Unmarshal("", &ThemeConfig); err != nil {
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
