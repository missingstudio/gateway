package cmd

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	USAGE     = "Usage"
	CORECMD   = "Core commands"
	HELPCMD   = "Help topics"
	OTHERCMD  = "Other commands"
	FLAGS     = "Flags"
	IFLAGS    = "Inherited flags"
	ARGUMENTS = "Arguments"
	EXAMPLES  = "Examples"
	ENVS      = "Environment variables"
	LEARN     = "Learn more"
	FEEDBACK  = "Feedback"
)

// SetHelp sets a custom help and usage function.
// It allows to group commands in different sections
// based on cobra commands annotations.
func SetHelp(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		rootHelpFunc(cmd, args)
	})
	cmd.SetUsageFunc(rootUsageFunc)
	cmd.SetFlagErrorFunc(rootFlagErrorFunc)
}

func rootUsageFunc(command *cobra.Command) error {
	command.Printf("Usage:  %s", command.UseLine())

	subcommands := command.Commands()
	if len(subcommands) > 0 {
		command.Print("\n\nAvailable commands:\n")
		for _, c := range subcommands {
			if c.Hidden {
				continue
			}
			command.Printf("  %s\n", c.Name())
		}
		return nil
	}

	flagUsages := command.LocalFlags().FlagUsages()
	if flagUsages != "" {
		command.Println("\n\nFlags:")
		command.Print(indent(dedent(flagUsages), "  "))
	}
	return nil
}

func rootFlagErrorFunc(cmd *cobra.Command, err error) error {
	if err == pflag.ErrHelp {
		return err
	}
	return err
}

func rootHelpFunc(command *cobra.Command, args []string) {
	if isRootCmd(command.Parent()) && len(args) >= 2 && args[1] != "--help" && args[1] != "-h" {
		nestedSuggestFunc(command, args[1])
		return
	}

	coreCommands := []string{}
	groupCommands := map[string][]string{}
	helpCommands := []string{}
	otherCommands := []string{}

	for _, c := range command.Commands() {
		if c.Short == "" || c.Hidden {
			continue
		}
		s := rpad(c.Name(), c.NamePadding()+3) + c.Short

		g, ok := c.Annotations["group"]
		if ok && g == "core" {
			coreCommands = append(coreCommands, s)
		} else if ok && g == "help" {
			helpCommands = append(helpCommands, s)
		} else if ok && g != "" {
			groupCommands[g] = append(groupCommands[g], s)
		} else {
			otherCommands = append(otherCommands, s)
		}
	}

	// If there are no core and other commands, assume everything is a core command
	if len(coreCommands) == 0 && len(groupCommands) == 0 {
		coreCommands = otherCommands
		otherCommands = []string{}
	}

	type helpEntry struct {
		Title string
		Body  string
	}

	text := command.Long

	if text == "" {
		text = command.Short
	}

	helpEntries := []helpEntry{}
	if text != "" {
		helpEntries = append(helpEntries, helpEntry{"", text})
	}

	helpEntries = append(helpEntries, helpEntry{USAGE, command.UseLine()})

	if len(coreCommands) > 0 {
		helpEntries = append(helpEntries, helpEntry{CORECMD, strings.Join(coreCommands, "\n")})
	}

	for name, cmds := range groupCommands {
		if len(cmds) > 0 {
			helpEntries = append(helpEntries, helpEntry{fmt.Sprint(toTitle(name) + " commands"), strings.Join(cmds, "\n")})
		}
	}

	if len(otherCommands) > 0 {
		helpEntries = append(helpEntries, helpEntry{OTHERCMD, strings.Join(otherCommands, "\n")})
	}

	if len(helpCommands) > 0 {
		helpEntries = append(helpEntries, helpEntry{HELPCMD, strings.Join(helpCommands, "\n")})
	}

	flagUsages := command.LocalFlags().FlagUsages()
	if flagUsages != "" {
		helpEntries = append(helpEntries, helpEntry{FLAGS, dedent(flagUsages)})
	}

	inheritedFlagUsages := command.InheritedFlags().FlagUsages()
	if inheritedFlagUsages != "" {
		helpEntries = append(helpEntries, helpEntry{IFLAGS, dedent(inheritedFlagUsages)})
	}

	if _, ok := command.Annotations["help:arguments"]; ok {
		helpEntries = append(helpEntries, helpEntry{ARGUMENTS, command.Annotations["help:arguments"]})
	}

	if command.Example != "" {
		helpEntries = append(helpEntries, helpEntry{EXAMPLES, command.Example})
	}

	if _, ok := command.Annotations["help:learn"]; ok {
		helpEntries = append(helpEntries, helpEntry{LEARN, command.Annotations["help:learn"]})
	}

	if _, ok := command.Annotations["help:feedback"]; ok {
		helpEntries = append(helpEntries, helpEntry{FEEDBACK, command.Annotations["help:feedback"]})
	}

	out := command.OutOrStdout()
	for _, e := range helpEntries {
		if e.Title != "" {
			// If there is a title, add indentation to each line in the body
			fmt.Fprintln(out, bold(e.Title))
			fmt.Fprintln(out, indent(strings.Trim(e.Body, "\r\n"), "  "))
		} else {
			// If there is no title print the body as is
			fmt.Println(e.Body)
		}
		fmt.Fprintln(out)
	}
}

// Display helpful error message in case subcommand name was mistyped.
func nestedSuggestFunc(command *cobra.Command, arg string) {
	command.Printf("unknown command %q for %q\n", arg, command.CommandPath())

	var candidates []string
	if arg == "help" {
		candidates = []string{"--help"}
	} else {
		if command.SuggestionsMinimumDistance <= 0 {
			command.SuggestionsMinimumDistance = 2
		}
		candidates = command.SuggestionsFor(arg)
	}

	if len(candidates) > 0 {
		command.Print("\nDid you mean this?\n")
		for _, c := range candidates {
			command.Printf("\t%s\n", c)
		}
	}

	command.Print("\n")
	_ = rootUsageFunc(command)
}

func isRootCmd(command *cobra.Command) bool {
	return command != nil && !command.HasParent()
}

var lineRE = regexp.MustCompile(`(?m)^`)

func indent(s, indent string) string {
	if len(strings.TrimSpace(s)) == 0 {
		return s
	}
	return lineRE.ReplaceAllLiteralString(s, indent)
}

func bold(text string) termenv.Style {
	h := termenv.String(text).Bold()
	return h
}

func toTitle(text string) string {
	heading := cases.Title(language.Und).String(text)
	return heading
}

func dedent(s string) string {
	lines := strings.Split(s, "\n")
	minIndent := -1

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		indent := len(l) - len(strings.TrimLeft(l, " "))
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	if minIndent <= 0 {
		return s
	}

	var buf bytes.Buffer
	for _, l := range lines {
		fmt.Fprintln(&buf, strings.TrimPrefix(l, strings.Repeat(" ", minIndent)))
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds ", padding)
	return fmt.Sprintf(template, s)
}
