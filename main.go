package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type page int

const (
	homePage page = iota
	aboutPage
	expPage
	contactPage
)

type menu struct {
	currentPage page
	cursor      int
	choices     []string
	selected    string
	//selected    map[int]struct{}
}

func initialMenu() menu {
	return menu{
		choices: []string{"Home", "About", "Experience", "Contact"},

		// A map which indicates which choices are selected. We're using
		// the map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		//selected: make(map[int]struct{}),
		selected:    "",
		currentPage: homePage,
	}
}

func (m menu) Init() tea.Cmd {
	return tea.SetWindowTitle("A text-based visualization experiment")
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc", "b":
			if m.currentPage != homePage {
				m.currentPage = homePage
				m.cursor = 0
				m.selected = ""
				return m, tea.ClearScreen
			}
		case "enter":
			if m.currentPage == homePage {
				m.selected = m.choices[m.cursor]
				switch m.selected {
				case "":
					m.currentPage = homePage
					return m, tea.ClearScreen
				case "About":
					m.currentPage = aboutPage
					return m, tea.ClearScreen
				case "Experience":
					m.currentPage = expPage
					return m, tea.ClearScreen
				case "Contact":
					m.currentPage = contactPage
					return m, tea.ClearScreen
				}
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else if m.cursor == 0 {
				m.cursor = len(m.choices) - 1
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else if m.cursor == len(m.choices)-1 {
				m.cursor = 0
			}
		}
	}
	return m, nil
}

func (m menu) View() string {
	var s string

	switch m.currentPage {
	case homePage:
		if m.selected == "Home" {
			s += m.renderMenu()
			s += "\nOops! You are already at the Home Page!"
		} else {
			s += m.renderMenu()
		}
	case aboutPage:
		s = ""
		s += m.renderAbout()
		s += "\nPress 'ESC' or 'b' to return to the home page."
	case expPage:
		s = ""
		s += m.renderExp()
		s += "\nPress 'ESC' or 'b' to return to the home page."
	case contactPage:
		s = ""
		s += m.renderContact()
		s += "\nPress 'ESC' or 'b' to return to the home page."
	}

	s += "\nPress 'q' to quit.\n"

	return s
}

func (m menu) renderMenu() string {
	s := "Welcome to this speck of sand in the universe, pick an option below to learn more about the stardust that composes me.\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "->"
		}

		s += fmt.Sprintf("%s [%s]\n", cursor, choice)
	}

	return s
}

func (m menu) renderAbout() string {
	s := "[About Me]\n\n"

	s += "## Who ##\n"
	s += "	A curious mind, a forever learner, a person built by empathy, and a swiss army knife technologist.\n"
	s += "## Why ##\n"
	s += "	Only child, early parent loss, isolation, a thirst for knowledge, a desire for belonging, a dedication to finding success.\n"
	s += "## How ##\n"
	s += "	Persistence.\n"

	return s
}

func (m menu) renderExp() string {
	s := "[Experience]\n\n"

	s += "## Education ##\n"
	s += "	Associates of Science - Computer Science - 2015\n"
	s += "	Bachleors of Science - Information Technology - Network and Information Security - 2019\n\n"
	s += "## Work ##\n"
	s += "	IT & Computer Programmer - Clinical Trial and Business Development Agency - 2018 - 2019\n"
	s += "	  -Data Systems Management - Salesforce, SharpSpring\n"
	s += "	  -Python Development - Web Scraping, Data Analysis\n"
	s += "	  -Network Administration - SOHO Routing, Wireless, Domain Admin - GoDaddy\n"
	s += "	  -M365 Administration - Email, Office365, M365 Admin Center\n"
	s += "	  -IT Support - Laptop, VoIP, Printer, TV/Digital Display - Break/Fix\n"
	s += "	Field Service Engineer - Government Contractor - 2019 - 2022\n"
	s += "	  -IT Support - Laptop, VoIP, Printer, TV/Digital Display - Break/Fix\n"
	s += "	  -Networking - Enterprise Switch and Router maintainence\n"
	s += "	  -Windows Administration - Enterprise Windows Server 2012-2022 maintainence, Windows 10-11 maintainence\n"
	s += "	  -Powershell Automation - Automated tasks, for workstation builds, file share management, large scale break/fix solutions\n"
	s += "	  -Workstation Deployment - Large scale enterprise scheduling, building and deployment of workstations\n"
	s += "	Network Deployment Engineer - Government Contractor - 2022\n"
	s += "	  -Networking - Cisco Enterprise Routers and Switches configuration, remediation, patching\n"
	s += "	  -Inventory Management - Cisco Vendor quotes, build of materials, purchase, and deployment of networking equipment\n"
	s += "	  -Field Deployment - Planning, Scheduling, Execution of network deployment for new sites, refreshed sites, emergency response sites\n"
	s += "	  -Protocols - BGP, EIGRP, SDWAN, Layer2 Vlan Routing, LACP Trunks, SFP, qSFP Fiber and Ethernet\n"
	s += "	IT Special in InfoSec - Federal Employee - 2022 - Present\n"
	s += "	  -Team Lead - Active Directory, Windows Systems Engineering, Platform Engineering\n"
	s += "	  -Planning, Scheduling, Execution of business requirements\n"
	s += "	  -Configuration Management - STIG Compliance, Golden Image, Cloud Pipeline Configuration, DevSecOps\n"
	s += "	  -Vulnerability Management - Enterprise remediation efforts via powershell automation and configuration manager device management\n"
	s += "	  -Budgetary - Inquire, Trial, Purchase new products to improve enterprise infrastructure management and engineering\n"
	s += "	  -Swiss Army Knife, Jack of All Trades\n"

	return s
}

func (m menu) renderContact() string {
	s := "[Contact]\n\n"
	s += "E-Mail -> maintainentropicprivacy@gmail.com\n"
	return s
}

func main() {
	p := tea.NewProgram(initialMenu(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
