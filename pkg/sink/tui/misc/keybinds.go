/*
Copyright © 2020 The PES Open Source Team pesos@pes.edu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package misc

// HelpKeybindingType is the type of the keybinding that
// a help page will use for a specific command.
type HelpKeybindingType int

const (
	// RootCommand is the keybinding identifier
	// for the "main" command of grofer, i.e. `grofer`.
	RootCommand HelpKeybindingType = iota
	// ProcCommand is the keybinding identifier
	// for the `grofer proc` command.
	ProcCommand
	// PerProcCommand is the keybinding identifier
	// for the `grofer proc -p <pid>` command.
	PerProcCommand
	// ContainerCommand is the keybinding identifier
	// for the `grofer container` command.
	ContainerCommand
	// PerContainerCommand is the keybinding identifier
	// for the `grofer container -c <CID>` command.
	PerContainerCommand
)

// getHelpKeybindingsForCommand returns the help keybinding for a specific command.
func getHelpKeybindingsForCommand(forCommand HelpKeybindingType) []string {
	switch forCommand {
	case RootCommand:
		return getMainCommandKeybindings()
	case ProcCommand:
		return getProcCommandKeybindings()
	case PerProcCommand:
		return getPerProcCommandKeybindings()
	case ContainerCommand:
		return getContainerCommandKeybindings()
	case PerContainerCommand:
		return getPerContainerCommandKeybindings()
	default:
		return getDefaultHelpKeybinding()
	}
}

func getErrorKeybindings() []string {
	return getDefaultHelpKeybinding()
}

func getDefaultHelpKeybinding() []string {
	return []string{
		"",
		"[To close this prompt: <Esc>](fg:white)",
	}
}

func getMainCommandKeybindings() []string {
	return []string{
		"Quit: q or <C-c>",
		"Pause Rendering: s",
		"Table Selection: <Left>/h and <Right>/l",
		"Table Scrolling: <Up>/k and <Down>/j",
		"Enable CPU Table: t",
		"",
		"[To close this prompt: <Esc>](fg:white)",
	}
}

func getProcCommandKeybindings() []string {
	return []string{
		"Quit: q or <C-c>",
		"Pause Rendering: s",
		"",
		"[Process navigation](fg:white)",
		"  - k and <Up>: up",
		"  - j and <Down>: down",
		"  - <C-u>: half page up",
		"  - <C-d>: half page down",
		"  - <C-b>: full page up",
		"  - <C-f>: full page down",
		"  - gg and <Home>: jump to top",
		"  - G and <End>: jump to bottom",
		"",
		"[Sorting](fg:white)",
		"  - Use column number to sort ascending.",
		"  - Use <F-column number> to sort descending.",
		"  - Eg: 1 to sort ascending on 1st Col and F1 for descending",
		"  - 0: Disable Sort",
		"",
		"[Process actions](fg:white)",
		"  - K and <F9>: Open signal selector menu",
		"",
		"[Signal selection](fg:white)",
		"  - K and <F9>: Send SIGTERM to selected process. Kills the process",
		"  - k and <Up>: up",
		"  - j and <Down>: down",
		"  - 0-9: navigate by numeric index",
		"  - <Enter>: send highlighted signal to process",
		"  - <Esc>: close signal selector",
		"",
		"[To close this prompt: <Esc>](fg:white)",
	}
}

func getPerProcCommandKeybindings() []string {
	return []string{
		"Quit: q or <C-c>",
		"Pause Rendering: s",
		"",
		"[To close this prompt: <Esc>](fg:white)",
	}
}

func getContainerCommandKeybindings() []string {
	return []string{
		"Quit: q or <C-c>",
		"Pause Rendering: s",
		"",
		"[Container navigation](fg:white)",
		"  - k and <Up>: up",
		"  - j and <Down>: down",
		"  - <C-u>: half page up",
		"  - <C-d>: half page down",
		"  - <C-b>: full page up",
		"  - <C-f>: full page down",
		"  - gg and <Home>: jump to top",
		"  - G and <End>: jump to bottom",
		"",
		"[Sorting](fg:white)",
		"  - Use column number to sort ascending.",
		"  - Use <F-column number> to sort descending.",
		"  - Eg: 1 to sort ascending on 1st Col and F1 for descending",
		"  - 0: Disable Sort",
		"",
		"[Container actions](fg:white)",
		"  - P: pause a container",
		"  - U: unpause a container",
		"  - R: restart a container",
		"  - S: stop a container",
		"  - K: kill a container",
		"  - X: remove a container (removes links & volumes)",
		"",
		"[To close this prompt: <Esc>](fg:white)",
	}
}

func getPerContainerCommandKeybindings() []string {
	return []string{
		"Quit: q or <C-c>",
		"Pause Rendering: s",
		"",
		"[Table Selection](fg:white)",
		"  - 1: MountTable",
		"  - 2: NetworkTable",
		"  - 3: CPUUsageTable",
		"  - 4: PortMapTable",
		"  - 5: ProcTable",
		"",
		"[Table navigation](fg:white)",
		"  - k and <Up>: up",
		"  - j and <Down>: down",
		"  - <C-u>: half page up",
		"  - <C-d>: half page down",
		"  - <C-b>: full page up",
		"  - <C-f>: full page down",
		"  - gg and <Home>: jump to top",
		"  - G and <End>: jump to bottom",
		"",
		"[To close this prompt: <Esc>](fg:white)",
	}
}