package style

var defaultThemeConfig = &Theme{
	List: ListColors{
		Header: HeaderColorSet{
			Title: TextColorSet{
				BG:  "#d65f50",
				TXT: "#ffffff",
			},
			ItemCount: TextColorSet{
				// BG:  "#ffffff",
				TXT: "#685e59",
			},
			FilterPrompt: TextColorSet{
				// BG:  "#ffffff",
				TXT: "#ECFD65",
			},
			FilterCursor: TextColorSet{
				// BG:  "#ffffff",
				TXT: "#ECFD65",
			},
		},
		StatusMessage: TextColorSet{
			// BG:  "",
			TXT: "#a2e3ad",
		},

		ActiveSelectionTitle: TextColorSet{
			// BG:  "#d28169",
			TXT: "#f79980",
		},
		NormalTitle: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#a1a09c",
		},
		// DimmedTitle: TextColorSet{
		// 	BG:  "#ffffff",
		// 	TXT: "#ffffff",
		// },
		ActiveSelectionDescription: TextColorSet{
			// BG:  "#9f614f",
			TXT: "#f79980",
		},
		NormalDescription: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#a1a09c",
		},
		DimmedDescription: TextColorSet{
			// BG:  "transparent",
			TXT: "#4b3a34",
		},
	},

	MainHeader: HeaderColorSet{
		Title: TextColorSet{
			BG:  "#d65f50",
			TXT: "#ffffff",
		},
		ItemCount: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#685e59",
		},
		FilterPrompt: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		FilterCursor: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
	},

	Table: TableColors{
		ActiveBorder: "#ffb8bc",
		DimmedBorder: "#4b3e3b",

		ActiveCell: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		DimmedCell: TextColorSet{
			BG:  "#4b3e3b",
			TXT: "#ffffff",
		},
		ActiveHeader: TextColorSet{
			BG:  "#c5636a",
			TXT: "#ffffff",
		},
		DimmedHeader: TextColorSet{
			BG:  "#4b3e3b",
			TXT: "#616161",
		},
		ActiveSelected: TextColorSet{
			BG:  "#ffb1b5",
			TXT: "#322223",
		},
		DimmedSelected: TextColorSet{
			BG:  "#616161",
			TXT: "#222222",
		},
	},

	Tab: TabColors{
		ActiveTabBorder:   "#ffb8bc",
		InactiveTabBorder: "#ffb8bc",
		// DimmedActiveTabBorder:   "#ffffff",
		DimmedInactiveTabBorder: "#4b3e3b",

		ActiveTab: TextColorSet{
			BG:  "#ffb1b5",
			TXT: "#322223",
		},
		InactiveTab: TextColorSet{
			// BG:  "#ffffff",
			// TXT: "#ffffff",
		},
		DimmedActiveTab: TextColorSet{
			BG:  "#4b3e3b",
			TXT: "#aa9890",
		},
		DimmedInactiveTab: TextColorSet{
			// BG:  "#4b3e3b",
			TXT: "#4b3e3b",
		},

		ActiveTabContentKeys:         "#ff9184",
		DimmedActiveTabContentKeys:   "#ffffff",
		ActiveTabContentValues:       "#ffffff",
		DimmedActiveTabContentValues: "#ffffff",
	},

	Paginator: PaginatorColorSet{
		Active:   "#ffb1b5",
		Inactive: "#624548",
		Dimmed:   "#4b3e3b",
	},

	SyncBar: SyncBarColors{
		Spinner: "#ffffff",
		InSyncState: TextColorSet{
			BG:  "#a4c847",
			TXT: "#ffffff",
		},
		NotSyncedState: TextColorSet{
			BG:  "#d83f24",
			TXT: "#ffffff",
		},
		StartingState: TextColorSet{
			BG:  "#4b3e3b",
			TXT: "#ffffff",
		},
		PausedState: TextColorSet{
			BG:  "#b39038",
			TXT: "#ffffff",
		},
	},

	Help: HelpColorSet{
		ShortcutName: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#e4d491",
		},
		ShortcutDescription: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#e4d491",
		},
		ShortcutSeparator: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#e4d491",
		},
	},

	StatusLog: StatusLogColors{
		Duration: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#595959",
		},
		Timestamp: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#ffffff",
		},
		OKStatus: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#a4c847",
		},
		NOKStatus: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#d65f50",
		},
		WarnStatus: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#ffde6e",
		},
		InfoLevel: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#498c69",
		},
		WarnLevel: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#e4d491",
		},
		ErrorLevel: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#d65f50",
		},
		DebugLevel: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#d282c0",
		},
	},

	StatusBar: StatusBarColors{
		ResourceSection: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		ResourceDetailsSection: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		ContextSection: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		NamespaceSection: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
	},
}
