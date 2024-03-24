package style

import (
	"log"
	"os"
	"regexp"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Theme struct {
	List ListColors `mapstructure:"list,omitempty"`

	MainHeader HeaderColorSet `mapstructure:"mainHeader,omitempty"`

	Table TableColors `mapstructure:"table,omitempty"`
	Tab   TabColors   `mapstructure:"tab,omitempty"`

	SyncBar   SyncBarColors   `mapstructure:"syncBar,omitempty"`
	Help      HelpColorSet    `mapstructure:"help,omitempty"`
	StatusLog StatusLogColors `mapstructure:"statusLog,omitempty"`
	StatusBar StatusBarColors `mapstructure:"statusBar,omitempty"`
}

type TextColorSet struct {
	BG  string `mapstructure:"bg,omitempty"`
	TXT string `mapstructure:"txt,omitempty"`
}

type HeaderColorSet struct {
	Title        TextColorSet `mapstructure:"title,omitempty"`
	ItemCount    TextColorSet `mapstructure:"itemCount,omitempty"`
	FilterPrompt TextColorSet `mapstructure:"filterPrompt,omitempty"`
	FilterCursor TextColorSet `mapstructure:"filterCursor,omitempty"`
}

type HelpColorSet struct {
	ShortcutName        TextColorSet `mapstructure:"shortcutName,omitempty"`
	ShortcutDescription TextColorSet `mapstructure:"shortcutDescription,omitempty"`
	ShortcutSeparator   TextColorSet `mapstructure:"shortcutSeparator,omitempty"`
}

type PaginatorColorSet struct {
	Active   string `mapstructure:"active,omitempty"`
	Inactive string `mapstructure:"inactive,omitempty"`
	Dimmed   string `mapstructure:"dimmed,omitempty"`
}

type ListColors struct {
	Header        HeaderColorSet `mapstructure:"header,omitempty"`
	StatusMessage TextColorSet   `mapstructure:"statusMessage,omitempty"`

	ActiveSelectionTitle TextColorSet `mapstructure:"activeSelection_title,omitempty"`
	NormalTitle          TextColorSet `mapstructure:"normalTitle,omitempty"`
	DimmedTitle          TextColorSet `mapstructure:"dimmedTitle,omitempty"`

	ActiveSelectionDescription TextColorSet `mapstructure:"activeSelectionDescription,omitempty"`
	NormalDescription          TextColorSet `mapstructure:"normalDescription,omitempty"`
	DimmedDescription          TextColorSet `mapstructure:"dimmedDescription,omitempty"`
}

type TableColors struct {
	ActiveBorder string `mapstructure:"activeBorder,omitempty"`
	DimmedBorder string `mapstructure:"dimmedBorder,omitempty"`

	ActiveCell TextColorSet `mapstructure:"activeCell,omitempty"`
	DimmedCell TextColorSet `mapstructure:"dimmedCell,omitempty"`

	ActiveHeader TextColorSet `mapstructure:"activeHeader,omitempty"`
	DimmedHeader TextColorSet `mapstructure:"dimmedHeader,omitempty"`

	ActiveSelected TextColorSet `mapstructure:"activeSelected,omitempty"`
	DimmedSelected TextColorSet `mapstructure:"dimmedSelected,omitempty"`

	Paginator PaginatorColorSet `mapstructure:"paginator,omitempty"`
}

type TabColors struct {
	ActiveTabBorder         string `mapstructure:"activeTabBorder,omitempty"`
	InactiveTabBorder       string `mapstructure:"inactiveTabBorder,omitempty"`
	DimmedActiveTabBorder   string `mapstructure:"dimmedActiveTabBorder,omitempty"`
	DimmedInactiveTabBorder string `mapstructure:"dimmedInactiveTabBorder,omitempty"`

	ActiveTab         TextColorSet `mapstructure:"activeTab,omitempty"`
	InactiveTab       TextColorSet `mapstructure:"inactiveTab,omitempty"`
	DimmedActiveTab   TextColorSet `mapstructure:"dimmedActiveTab,omitempty"`
	DimmedInactiveTab TextColorSet `mapstructure:"dimmedInactiveTab,omitempty"`

	ActiveTabContentKeys       string `mapstructure:"activeTabContent_keys,omitempty"`
	DimmedActiveTabContentKeys string `mapstructure:"dimmedActiveTabContentKeys,omitempty"`

	ActiveTabContentValues       string `mapstructure:"activeTabContentValues,omitempty"`
	DimmedActiveTabContentValues string `mapstructure:"dimmedActiveTabContentValues,omitempty"`

	Paginator PaginatorColorSet `mapstructure:"paginator,omitempty"`
}

type SyncBarColors struct {
	Spinner       string       `mapstructure:"spinner,omitempty"`
	InSyncState   TextColorSet `mapstructure:"inSyncState,omitempty"`
	UnsyncedState TextColorSet `mapstructure:"unsyncedState,omitempty"`
	StartingState TextColorSet `mapstructure:"startingState,omitempty"`
}

type StatusLogColors struct {
	Duration  TextColorSet `mapstructure:"duration,omitempty"`
	Timestamp TextColorSet `mapstructure:"timestamp,omitempty"`

	OKStatus   TextColorSet `mapstructure:"okStatus,omitempty"`
	NOKStatus  TextColorSet `mapstructure:"nokStatus,omitempty"`
	WarnStatus TextColorSet `mapstructure:"warnStatus,omitempty"`

	InfoLevel  TextColorSet `mapstructure:"infoLevel,omitempty"`
	WarnLevel  TextColorSet `mapstructure:"warnLevel,omitempty"`
	ErrorLevel TextColorSet `mapstructure:"errorLevel,omitempty"`
	DebugLevel TextColorSet `mapstructure:"debugLevel,omitempty"`
}

type StatusBarColors struct {
	ResourceSection        TextColorSet `mapstructure:"resourceSection,omitempty"`
	ResourceDetailsSection TextColorSet `mapstructure:"resourceDetails_section,omitempty"`
	ContextSection         TextColorSet `mapstructure:"contextSection,omitempty"`
	NamespaceSection       TextColorSet `mapstructure:"namespaceSection,omitempty"`
}

var ThemeConfig Theme
var v = viper.New()

func (t *Theme) Marshal() []byte {
	out, err := yaml.Marshal(t)
	if err != nil {
		log.Fatalf("failed to marshal theme: %s", err)
	}
	return out
}

func init() {

	themeConfigFilePath, _ := xdg.ConfigFile("kibe/theme.yaml")
	v.SetConfigFile(themeConfigFilePath)

	err := v.ReadInConfig()
	if err != nil {
		file, err := os.Create(themeConfigFilePath)
		if err != nil {
			log.Fatalf("failed to create config file: %s", err)
		}
		_, err = file.Write(defaultThemeConfig.Marshal())
		if err != nil {
			log.Fatalf("failed to write config file: %s", err)
		}
	}

	err = v.Unmarshal(&ThemeConfig)
	if err != nil {
		log.Fatalf("failed to unmarshal theme config file: %s", err)
	}
}

func GetColor(c string) lipgloss.TerminalColor {
	colorPattern, _ := regexp.Compile(`^#[0-9a-fA-F]{6}$`)
	if colorPattern.MatchString(c) {
		return lipgloss.Color(c)
	} else {
		return lipgloss.NoColor{}
	}
}
