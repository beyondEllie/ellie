package actions

import (
	"github.com/tacheraSasi/ellie/styles"
)

func HandleService(action, service string) {
	actionVerb := action + "ing"
	styles.InfoStyle.Printf("%s %s service...\n", actionVerb, service)

	switch action {
	case "start":
		switch service {
		case "apache":
			StartApache()
		case "mysql":
			StartMysql()
		case "postgres":
			StartPostgres()
		case "all":
			StartAll()
		}
	case "stop":
		switch service {
		case "apache":
			StopApache()
		case "mysql":
			StopMysql()
		case "postgres":
			StopPostgres()
		case "all":
			StopAll()
		}
	case "restart":
		switch service {
		case "apache":
			StopApache()
			StartApache()
		case "mysql":
			StopMysql()
			StartMysql()
		case "postgres":
			StopPostgres()
			StartPostgres()
		case "all":
			StopAll()
			StartAll()
		}
	}
}
