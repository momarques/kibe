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
	List ListColors `koanf:"list,omitempty"`

	MainHeader HeaderColorSet `koanf:"mainHeader,omitempty"`

	Table TableColors `koanf:"table,omitempty"`
	Tab   TabColors   `koanf:"tab,omitempty"`

	SyncBar   SyncBarColors   `koanf:"syncBar,omitempty"`
	Help      HelpColorSet    `koanf:"help,omitempty"`
	StatusLog StatusLogColors `koanf:"statusLog,omitempty"`
	StatusBar StatusBarColors `koanf:"statusBar,omitempty"`
}

type TextColorSet struct {
	BG  string `koanf:"bg,omitempty"`
	TXT string `koanf:"txt,omitempty"`
}

type HeaderColorSet struct {
	Title        TextColorSet `koanf:"title,omitempty"`
	ItemCount    TextColorSet `koanf:"itemCount,omitempty"`
	FilterPrompt TextColorSet `koanf:"filterPrompt,omitempty"`
	FilterCursor TextColorSet `koanf:"filterCursor,omitempty"`
}

type HelpColorSet struct {
	ShortcutName        TextColorSet `koanf:"shortcutName,omitempty"`
	ShortcutDescription TextColorSet `koanf:"shortcutDescription,omitempty"`
	ShortcutSeparator   TextColorSet `koanf:"shortcutSeparator,omitempty"`
}

type PaginatorColorSet struct {
	Active   string `koanf:"active,omitempty"`
	Inactive string `koanf:"inactive,omitempty"`
	Dimmed   string `koanf:"dimmed,omitempty"`
}

type ListColors struct {
	Header        HeaderColorSet `koanf:"header,omitempty"`
	StatusMessage TextColorSet   `koanf:"statusMessage,omitempty"`

	ActiveSelectionTitle TextColorSet `koanf:"activeSelection_title,omitempty"`
	NormalTitle          TextColorSet `koanf:"normalTitle,omitempty"`
	DimmedTitle          TextColorSet `koanf:"dimmedTitle,omitempty"`

	ActiveSelectionDescription TextColorSet `koanf:"activeSelectionDescription,omitempty"`
	NormalDescription          TextColorSet `koanf:"normalDescription,omitempty"`
	DimmedDescription          TextColorSet `koanf:"dimmedDescription,omitempty"`
}

type TableColors struct {
	ActiveBorder string `koanf:"activeBorder,omitempty"`
	DimmedBorder string `koanf:"dimmedBorder,omitempty"`

	ActiveCell TextColorSet `koanf:"activeCell,omitempty"`
	DimmedCell TextColorSet `koanf:"dimmedCell,omitempty"`

	ActiveHeader TextColorSet `koanf:"activeHeader,omitempty"`
	DimmedHeader TextColorSet `koanf:"dimmedHeader,omitempty"`

	ActiveSelected TextColorSet `koanf:"activeSelected,omitempty"`
	DimmedSelected TextColorSet `koanf:"dimmedSelected,omitempty"`

	Paginator PaginatorColorSet `koanf:"paginator,omitempty"`
}

type TabColors struct {
	ActiveTabBorder         string `koanf:"activeTabBorder,omitempty"`
	InactiveTabBorder       string `koanf:"inactiveTabBorder,omitempty"`
	DimmedActiveTabBorder   string `koanf:"dimmedActiveTabBorder,omitempty"`
	DimmedInactiveTabBorder string `koanf:"dimmedInactiveTabBorder,omitempty"`

	ActiveTab         TextColorSet `koanf:"activeTab,omitempty"`
	InactiveTab       TextColorSet `koanf:"inactiveTab,omitempty"`
	DimmedActiveTab   TextColorSet `koanf:"dimmedActiveTab,omitempty"`
	DimmedInactiveTab TextColorSet `koanf:"dimmedInactiveTab,omitempty"`

	ActiveTabContentKeys       string `koanf:"activeTabContent_keys,omitempty"`
	DimmedActiveTabContentKeys string `koanf:"dimmedActiveTabContentKeys,omitempty"`

	ActiveTabContentValues       string `koanf:"activeTabContentValues,omitempty"`
	DimmedActiveTabContentValues string `koanf:"dimmedActiveTabContentValues,omitempty"`

	Paginator PaginatorColorSet `koanf:"paginator,omitempty"`
}

type SyncBarColors struct {
	Spinner       string       `koanf:"spinner,omitempty"`
	InSyncState   TextColorSet `koanf:"inSyncState,omitempty"`
	UnsyncedState TextColorSet `koanf:"unsyncedState,omitempty"`
	StartingState TextColorSet `koanf:"startingState,omitempty"`
}

type StatusLogColors struct {
	Duration  TextColorSet `koanf:"duration,omitempty"`
	Timestamp TextColorSet `koanf:"timestamp,omitempty"`

	OKStatus   TextColorSet `koanf:"okStatus,omitempty"`
	NOKStatus  TextColorSet `koanf:"nokStatus,omitempty"`
	WarnStatus TextColorSet `koanf:"warnStatus,omitempty"`

	InfoLevel  TextColorSet `koanf:"infoLevel,omitempty"`
	WarnLevel  TextColorSet `koanf:"warnLevel,omitempty"`
	ErrorLevel TextColorSet `koanf:"errorLevel,omitempty"`
	DebugLevel TextColorSet `koanf:"debugLevel,omitempty"`
}

type StatusBarColors struct {
	ResourceSection        TextColorSet `koanf:"resourceSection,omitempty"`
	ResourceDetailsSection TextColorSet `koanf:"resourceDetails_section,omitempty"`
	ContextSection         TextColorSet `koanf:"contextSection,omitempty"`
	NamespaceSection       TextColorSet `koanf:"namespaceSection,omitempty"`
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
