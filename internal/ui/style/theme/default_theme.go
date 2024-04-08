package theme

var defaultTheme = &Theme{
	ClientConfig: ClientConfigColors{
		Header: HeaderColorSet{
			Title: TextColorSet{
				BG:  "#f7c768",
				TXT: "#4f4156",
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
			TXT: "#c28bfb",
		},
		NormalTitle: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#dbded9",
		},
		// DimmedTitle: TextColorSet{
		// 	BG:  "#ffffff",
		// 	TXT: "#ffffff",
		// },
		ActiveSelectionDescription: TextColorSet{
			// BG:  "#9f614f",
			TXT: "#d8b0f9",
		},
		NormalDescription: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#a1a09c",
		},
		DimmedDescription: TextColorSet{
			// BG:  "transparent",
			TXT: "#665672",
		},
	},

	MainHeader: HeaderColorSet{
		Title: TextColorSet{
			BG:  "#f7c768",
			TXT: "#4f4156",
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
		ActiveBorder: "#c28bfb",
		DimmedBorder: "#40364a",

		ActiveCell: TextColorSet{
			// BG:  "#ffffff",
			TXT: "#ffffff",
		},
		DimmedCell: TextColorSet{
			BG:  "#4b3e3b",
			TXT: "#616161",
		},
		ActiveHeader: TextColorSet{
			BG:  "#986cbd",
			TXT: "#ffffff",
		},
		DimmedHeader: TextColorSet{
			BG:  "#43354a",
			TXT: "#616161",
		},
		ActiveSelected: TextColorSet{
			BG:  "#d8b0f9",
			TXT: "#322223",
		},
		DimmedSelected: TextColorSet{
			BG:  "#616161",
			TXT: "#616161",
		},
	},

	Tab: TabColors{
		ActiveTabBorder:   "#c28bfb",
		InactiveTabBorder: "#c28bfb",
		// DimmedActiveTabBorder:   "#ffffff",
		DimmedInactiveTabBorder: "#40364a",

		ActiveTab: TextColorSet{
			BG:  "#d8b0f9",
			TXT: "#322223",
		},
		InactiveTab: TextColorSet{
			// BG:  "#ffffff",
			// TXT: "#ffffff",
		},
		DimmedActiveTab: TextColorSet{
			BG:  "#40364a",
			TXT: "#a193a9",
		},
		DimmedInactiveTab: TextColorSet{
			// BG:  "#4b3e3b",
			TXT: "#40364a",
		},

		ActiveTabContentKeys:         "#f8c86a",
		DimmedActiveTabContentKeys:   "#ffffff",
		ActiveTabContentValues:       "#ffffff",
		DimmedActiveTabContentValues: "#ffffff",
	},

	Paginator: PaginatorColorSet{
		Active:   "#c28bfb",
		Inactive: "#775497",
		Dimmed:   "#483955",
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

	StatusBar: StatusBarColors{
		ResourceSection: TextColorSet{
			BG:  "#f8c86a",
			TXT: "#4f4156",
		},
		ResourceDetailsSection: TextColorSet{
			BG:  "#303040",
			TXT: "#ffffff",
		},
		ContextSection: TextColorSet{
			BG:  "#3195ef",
			TXT: "#ffffff",
		},
		NamespaceSection: TextColorSet{
			BG:  "#d06bcd",
			TXT: "#ffffff",
		},
	},

	StatusMessage: StatusMessageColors{
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
	},
}
