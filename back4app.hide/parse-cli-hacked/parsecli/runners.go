package parsecli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type cobraRun func(cmd *cobra.Command, args []string)

// RunNoArgs wraps a run function that shouldn't get any arguments.
func RunNoArgs(e *Env, f func(*Env) error) cobraRun {
	return func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Fprintf(e.Err, "unexpected arguments:%+v\n\n", args)
			cmd.Help()
			e.Exit(1)
		}
		if err := f(e); err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}
	}
}

// RunWithArgs wraps a run function that can access arguments to cobraCmd
func RunWithArgs(e *Env, f func(*Env, []string) error) cobraRun {
	return func(cmd *cobra.Command, args []string) {
		if err := f(e, args); err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}
	}
}

// RunWithClient wraps a run function that should get an app, when the default is
// picked from the config in the current working directory.
func RunWithClient(e *Env, f func(*Env, *Context) error) cobraRun {
	return func(cmd *cobra.Command, args []string) {
		app := DefaultKey
		if len(args) > 1 {
			fmt.Fprintf(
				e.Err,
				"unexpected arguments, only an optional app name is expected:%+v\n\n",
				args,
			)
			cmd.Help()
			e.Exit(1)
		}
		if len(args) == 1 {
			app = args[0]
		}
		cl, err := newContext(e, app)
		if err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}
		if err := f(e, cl); err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}
	}
}

// RunWithClientConfirm wraps a run function that takes an app name
// or asks for confirmation to use the default app name instead
func RunWithClientConfirm(e *Env, f func(*Env, *Context) error) cobraRun {
	return func(cmd *cobra.Command, args []string) {
		app := DefaultKey
		if len(args) > 1 {
			fmt.Fprintf(
				e.Err,
				"unexpected arguments, only an optional app name is expected:%+v\n\n",
				args,
			)
			cmd.Help()
			e.Exit(1)
		}
		if len(args) == 1 {
			app = args[0]
		}
		cl, err := newContext(e, app)
		if err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}

		if app == DefaultKey {
			fmt.Fprintf(
				e.Out,
				`Did not provide app name as an argument to the command.
Please enter an app name to execute this command on
or press ENTER to use the default app %q: `,
				cl.Config.GetDefaultApp(),
			)
			var appName string
			fmt.Fscanf(e.In, "%s\n", &appName)
			appName = strings.TrimSpace(appName)
			if appName != "" {
				cl, err = newContext(e, appName)
				if err != nil {
					fmt.Fprintln(e.Err, ErrorString(e, err))
					e.Exit(1)
				}
			}
		}

		if err := f(e, cl); err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}
	}
}

// RunWithArgsClient wraps a run function that should get an app, whee the default is
// picked from the config in the current working directory. It also passes args to the
// runner function
func RunWithArgsClient(e *Env, f func(*Env, *Context, []string) error) cobraRun {
	return func(cmd *cobra.Command, args []string) {
		app := DefaultKey
		if len(args) > 1 {
			app = args[0]
			args = args[1:]
		}
		cl, err := newContext(e, app)
		if err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}
		if err := f(e, cl, args); err != nil {
			fmt.Fprintln(e.Err, ErrorString(e, err))
			e.Exit(1)
		}
	}
}
