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
			BG:  "#d28169",
			TXT: "#ffffff",
		},
		NormalTitle: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#f8f3e8",
		},
		// DimmedTitle: TextColorSet{
		// 	BG:  "#ffffff",
		// 	TXT: "#ffffff",
		// },
		ActiveSelectionDescription: TextColorSet{
			BG:  "#9f614f",
			TXT: "#ffffff",
		},
		NormalDescription: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#8d6f62",
		},
		DimmedDescription: TextColorSet{
			// BG:  "transparent",
			TXT: "#8d6f62",
		},
	},

	MainHeader: HeaderColorSet{
		Title: TextColorSet{
			BG:  "#ffffff",
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

		Paginator: PaginatorColorSet{
			Active:   "#ffffff",
			Inactive: "#ffffff",
			Dimmed:   "#ffffff",
		},
	},

	Tab: TabColors{
		ActiveTabBorder:   "#ffb8bc",
		InactiveTabBorder: "#4b3e3b",
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

		Paginator: PaginatorColorSet{
			Active:   "#ffffff",
			Inactive: "#ffffff",
			Dimmed:   "#ffffff",
		},
	},

	SyncBar: SyncBarColors{
		Spinner: "#ffffff",
		InSyncState: TextColorSet{
			BG:  "#a4c847",
			TXT: "#ffffff",
		},
		UnsyncedState: TextColorSet{
			BG:  "#d83f24",
			TXT: "#ffffff",
		},
		StartingState: TextColorSet{
			BG:  "#4b3e3b",
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
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		Timestamp: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		OKStatus: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#a2e3ad",
		},
		NOKStatus: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#f35143",
		},
		WarnStatus: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#ffde6e",
		},
		InfoLevel: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		WarnLevel: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		ErrorLevel: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
		},
		DebugLevel: TextColorSet{
			BG:  "#ffffff",
			TXT: "#ffffff",
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
