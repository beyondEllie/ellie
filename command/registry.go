package command

import (
	"flag"
	"fmt"

	actions "github.com/tacheraSasi/ellie/action"
	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/types"
)

var Registry = map[string]Command{
	"run": {
		Handler: actions.Run,
	},
	"::": {
		Usage:   "ellie :: run docker container for me please",
		MinArgs: 1,
		Handler: actions.SmartRun,
	},
	"code": {
		Usage: "ellie code",
		Handler: func(_ []string) {
			actions.StartEllieCode()
		},
	},
	"user-env": {
		Handler: func(s []string) {
			// Create user context
			userCtx := types.NewUserContext()

			// Add system message with instructions and context
			instructions := fmt.Sprintf(`!!!!!!!!!!!!!!!!!!!!!IMPORTANT YOU ARE ELLIE note: %s `, static.Instructions(*userCtx))
			fmt.Println(instructions)
		},
	},
	"focus": {
		PreHook: func() { styles.InfoStyle.Println("Activating focus mode...") },
		Handler: actions.Focus,
	},
	"pwd": {
		Handler: func(_ []string) { actions.Pwd() },
	},
	"size": {
		MinArgs: 1,
		Handler: func(s []string) { actions.Size() },
	},
	"open-explorer": {
		Handler: func(_ []string) { actions.OpenExplorer() },
	},
	"open": {
		Usage:   "open <path>",
		MinArgs: 1,
		Handler: func(args []string) {
			actions.OpenExplorer(args[1])
		},
	},
	"play": {
		MinArgs: 1,
		Usage:   "play <media>",
		PreHook: func() { styles.InfoStyle.Println("Initializing media player...") },
		Handler: actions.Play,
	},
	"setup-git": {
		Handler: func(args []string) {
			actions.GitSetup(configs.GetEnv("PAT"), configs.GetEnv("USERNAME"))
		},
	},
	"sysinfo": {
		Handler: func(_ []string) { actions.SysInfo() },
	},
	"disk": {
		Usage:   "disk [path] - Show disk usage information",
		Handler: actions.Disk,
		SubCommands: map[string]Command{
			"all": {
				Usage:   "disk all - Show all disk partitions",
				Handler: func(_ []string) { actions.DiskAll() },
			},
			"space": {
				Usage:   "disk space - Show disk space summary",
				Handler: func(_ []string) { actions.DiskSpace() },
			},
		},
	},
	"dev-init": {
		Handler: func(args []string) {
			fs := flag.NewFlagSet("dev-init", flag.ExitOnError)
			allFlag := fs.Bool("all", false, "Install all recommended tools")
			fs.Parse(args[1:])
			actions.DevInit(*allFlag)
		},
	},
	"server-init": {
		Handler: func(_ []string) { actions.ServerInit() },
	},
	"install": {
		MinArgs: 1,
		Usage:   "install <package>",
		Handler: func(args []string) { actions.InstallPackage(args[1]) },
	},
	"update": {
		Handler: func(_ []string) { actions.UpdatePackages() },
	},
	"list": {
		MinArgs: 1,
		Usage:   "list <directory>",
		Handler: func(args []string) { actions.ListFiles(args[1]) },
	},
	"create-file": {
		MinArgs: 1,
		Usage:   "create-file <path>",
		Handler: func(args []string) { actions.CreateFile(args[1]) },
	},
	"network-status": {
		Handler: func(_ []string) { actions.NetworkStatus() },
	},
	"connect-wifi": {
		MinArgs: 2,
		Usage:   "connect-wifi <SSID> <password>",
		Handler: func(args []string) { actions.ConnectWiFi(args[1], args[2]) },
	},
	"greet": {
		Handler: func(_ []string) {
			styles.Highlight.Println("Your majesty,", configs.GetEnv("USERNAME"))
		},
	},
	"send-mail": {
		Handler: func(_ []string) { actions.Mailer() },
	},
	"chat": {
		Handler: func(_ []string) { actions.Chat(configs.GetEnv("OPENAI_API_KEY")) },
	},
	"review": {
		Usage:   "review <filename/filepath>",
		MinArgs: 1,
		Handler: func(args []string) { actions.Review(args[1]) },
		// PreHook: ,
	},
	"security-check": {
		Usage:   "security-check <path>",
		MinArgs: 1,
		Handler: func(args []string) { actions.SecurityCheck(args[1]) },
		// PreHook: ,
	},
	"git": {
		SubCommands: map[string]Command{
			// Basic operations
			"status": {Handler: func(_ []string) { actions.GitStatus() }},
			"push":   {Handler: func(_ []string) { actions.GitPush() }},
			"commit": {Handler: func(args []string) { actions.GitConventionalCommit() }},
			"pull":   {Handler: func(_ []string) { actions.GitPull() }},
			"init":   {Handler: func(_ []string) { actions.GitInit() }},
			"clone":  {Handler: func(_ []string) { actions.GitClone() }},
			"fetch":  {Handler: func(_ []string) { actions.GitFetch() }},

			// Branch operations
			"branch-create":      {Handler: func(_ []string) { actions.GitBranchCreate() }},
			"branch-switch":      {Handler: func(_ []string) { actions.GitBranchSwitch() }},
			"branch-delete":      {Handler: func(_ []string) { actions.GitBranchDelete() }},
			"branch-list":        {Handler: func(_ []string) { actions.GitBranchList() }},
			"branch-list-remote": {Handler: func(_ []string) { actions.GitBranchListRemote() }},
			"branch-rename":      {Handler: func(_ []string) { actions.GitBranchRename() }},

			// Stash operations
			"stash-save":  {Handler: func(_ []string) { actions.GitStashSave() }},
			"stash-pop":   {Handler: func(_ []string) { actions.GitStashPop() }},
			"stash-list":  {Handler: func(_ []string) { actions.GitStashList() }},
			"stash-show":  {Handler: func(_ []string) { actions.GitStashShow() }},
			"stash-drop":  {Handler: func(_ []string) { actions.GitStashDrop() }},
			"stash-apply": {Handler: func(_ []string) { actions.GitStashApply() }},

			// Tag operations
			"tag-create": {Handler: func(_ []string) { actions.GitTagCreate() }},
			"tag-list":   {Handler: func(_ []string) { actions.GitTagList() }},
			"tag-delete": {Handler: func(_ []string) { actions.GitTagDelete() }},

			// Log and diff operations
			"log":         {Handler: func(_ []string) { actions.GitLogPretty() }},
			"log-search":  {Handler: func(_ []string) { actions.GitLogSearch() }},
			"log-author":  {Handler: func(_ []string) { actions.GitLogAuthor() }},
			"log-since":   {Handler: func(_ []string) { actions.GitLogSince() }},
			"diff":        {Handler: func(_ []string) { actions.GitDiff() }},
			"diff-staged": {Handler: func(_ []string) { actions.GitDiffStaged() }},
			"diff-branch": {Handler: func(_ []string) { actions.GitDiffBranch() }},

			// Merge and rebase operations
			"merge":       {Handler: func(_ []string) { actions.GitMerge() }},
			"rebase":      {Handler: func(_ []string) { actions.GitRebase() }},
			"cherry-pick": {Handler: func(_ []string) { actions.GitCherryPick() }},
			"reset":       {Handler: func(_ []string) { actions.GitReset() }},
			"revert":      {Handler: func(_ []string) { actions.GitRevert() }},

			// Bisect operations
			"bisect":       {Handler: func(_ []string) { actions.GitBisect() }},
			"bisect-good":  {Handler: func(_ []string) { actions.GitBisectGood() }},
			"bisect-bad":   {Handler: func(_ []string) { actions.GitBisectBad() }},
			"bisect-reset": {Handler: func(_ []string) { actions.GitBisectReset() }},

			// Remote operations
			"remote-list":   {Handler: func(_ []string) { actions.GitRemoteList() }},
			"remote-add":    {Handler: func(_ []string) { actions.GitRemoteAdd() }},
			"remote-remove": {Handler: func(_ []string) { actions.GitRemoteRemove() }},

			// Push operations
			"push-tags":     {Handler: func(_ []string) { actions.GitPushTags() }},
			"push-force":    {Handler: func(_ []string) { actions.GitPushForce() }},
			"push-upstream": {Handler: func(_ []string) { actions.GitPushUpstream() }},

			// Submodule operations
			"submodule-add":    {Handler: func(_ []string) { actions.GitSubmoduleAdd() }},
			"submodule-update": {Handler: func(_ []string) { actions.GitSubmoduleUpdate() }},
			"submodule-status": {Handler: func(_ []string) { actions.GitSubmoduleStatus() }},

			// Configuration operations
			"config-set-user":  {Handler: func(_ []string) { actions.GitConfigSetUser() }},
			"config-list":      {Handler: func(_ []string) { actions.GitConfigList() }},
			"config-set-alias": {Handler: func(_ []string) { actions.GitConfigSetAlias() }},

			// Worktree operations
			"worktree-add":    {Handler: func(_ []string) { actions.GitWorktreeAdd() }},
			"worktree-list":   {Handler: func(_ []string) { actions.GitWorktreeList() }},
			"worktree-remove": {Handler: func(_ []string) { actions.GitWorktreeRemove() }},
			"worktree-prune":  {Handler: func(_ []string) { actions.GitWorktreePrune() }},

			// Maintenance operations
			"reflog": {Handler: func(_ []string) { actions.GitReflog() }},
			"clean":  {Handler: func(_ []string) { actions.GitClean() }},
			"gc":     {Handler: func(_ []string) { actions.GitGC() }},
			"fsck":   {Handler: func(_ []string) { actions.GitFsck() }},

			// Information operations
			"show":    {Handler: func(_ []string) { actions.GitShow() }},
			"blame":   {Handler: func(_ []string) { actions.GitBlame() }},
			"archive": {Handler: func(_ []string) { actions.GitArchive() }},
		},
	},
	"docker": {
		SubCommands: map[string]Command{
			"build": {
				MinArgs: 1,
				Usage:   "docker build <path>",
				Handler: func(args []string) { actions.DockerBuild(args[1:]) },
			},
			"run": {
				MinArgs: 1,
				Usage:   "docker run <image>",
				Handler: func(args []string) { actions.DockerRun(args[1:]) },
			},
			"ps": {
				Handler: func(args []string) { actions.DockerPS(args[1:]) },
			},
			"compose": {
				SubCommands: map[string]Command{
					"up": {
						MinArgs: 0,
						Usage:   "docker compose up",
						Handler: func(args []string) { actions.DockerCompose(args) },
					},
					"down": {
						MinArgs: 0,
						Usage:   "docker compose down",
						Handler: func(args []string) { actions.DockerCompose(args) },
					},
				},
			},
		},
	},
	"start": {
		SubCommands: map[string]Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("start", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("start", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("start", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("start", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"stop": {
		SubCommands: map[string]Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("stop", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("stop", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("stop", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("stop", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"restart": {
		SubCommands: map[string]Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("restart", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("restart", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("restart", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("restart", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"config": {
		Handler: func(_ []string) { configs.Init() },
	},
	"reset-config": {
		Handler: func(_ []string) { configs.ResetConfig() },
	},
	"whoami": {
		Handler: func(_ []string) {
			styles.Highlight.Println("Your majesty,", configs.GetEnv("USERNAME"))
		},
	},
	"alias": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 1,
				Usage:   "alias add <name>=\"<command>\"",
				Handler: actions.AliasAdd,
			},
			"list": {
				Handler: actions.AliasList,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "alias delete <name>",
				Handler: actions.AliasDelete,
			},
		},
	},
	"todo": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 1,
				Usage:   "todo add \"<task>\" [category] [priority]",
				Handler: actions.TodoAdd,
			},
			"list": {
				Handler: actions.TodoList,
			},
			"complete": {
				MinArgs: 1,
				Usage:   "todo complete <id>",
				Handler: actions.TodoComplete,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "todo delete <id>",
				Handler: actions.TodoDelete,
			},
			"edit": {
				MinArgs: 3,
				Usage:   "todo edit <id> <field> <value>",
				Handler: actions.TodoEdit,
			},
		},
	},
	"project": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 2,
				Usage:   "project add <name> <path> [description] [tags...]",
				Handler: actions.ProjectAdd,
			},
			"list": {
				Handler: actions.ProjectList,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "project delete <name>",
				Handler: actions.ProjectDelete,
			},
			"search": {
				MinArgs: 1,
				Usage:   "project search <query>",
				Handler: actions.ProjectSearch,
			},
		},
	},
	"switch": {
		MinArgs: 1,
		Usage:   "switch <project-name>",
		Handler: actions.ProjectSwitch,
	},
	"history": {
		Handler: actions.History,
	},
	"start-day": {
		Handler: actions.StartDay,
	},
	"day-start": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 2,
				Usage:   "day-start add <type> <value>",
				Handler: actions.DayStartConfigAdd,
			},
			"list": {
				Handler: actions.DayStartConfigList,
			},
		},
	},

	//Pending commands
	"weather": {
		Handler: func(args []string) { actions.Weather() },
	},
	"joke": {
		Handler: func(args []string) { actions.Joke() },
	},
	"remind": {
		Handler: func(_ []string) { actions.Remind() },
	},
	"about": {
		Handler: actions.ShowAbout,
	},
	"theme": {
		SubCommands: map[string]Command{
			"set": {
				MinArgs: 1,
				Usage:   "theme set <light|dark|auto>",
				Handler: func(args []string) {
					mode := args[1]
					if mode != "light" && mode != "dark" && mode != "auto" {
						styles.GetErrorStyle().Println("Invalid theme. Use 'light', 'dark', or 'auto'.")
						return
					}
					styles.SetTheme(mode)
					styles.GetSuccessStyle().Printf("Theme set to %s.\n", styles.GetTheme())
				},
			},
			"show": {
				Handler: func(_ []string) {
					styles.GetInfoStyle().Printf("Current theme: %s\n", styles.GetTheme())
				},
			},
		},
	},
	"md": {
		Usage:   "md <filename>",
		MinArgs: 1,
		Handler: actions.MarkdownRender,
	},
}
