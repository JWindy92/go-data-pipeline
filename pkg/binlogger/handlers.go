package binlogger

import (
	"fmt"
	"runtime/debug"

	"github.com/JWindy92/go-data-pipeline/config"
	"github.com/go-mysql-org/go-mysql/canal"
)

type binlogHandler struct {
	canal.DummyEventHandler
	BinlogParser
	Schema string
	Table  string
}

func (h *binlogHandler) OnRow(e *canal.RowsEvent) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r, " ", string(debug.Stack()))
		}
	}()

	// base value for canal.DeleteAction or canal.InsertAction
	var n = 0
	var k = 1

	if e.Action == canal.UpdateAction {
		n = 1
		k = 2
	}

	for i := n; i < len(e.Rows); i += k {

		key := e.Table.Schema + "." + e.Table.Name

		switch key {
		case h.Schema + "." + h.Table:
			dbEvent := DbEvent{}
			h.GetBinLogData(&dbEvent, e, i)
			switch e.Action {
			case canal.UpdateAction:
				// oldUser := DbEvent{}
				// h.GetBinLogData(&oldUser, e, i-1)
				// fmt.Printf("User %d name changed from %s to %s\n", user.Id, oldUser.Name, user.Name)
				fmt.Printf("UPDATE: EventId: %d, Table: %s, RecordId: %d\n", dbEvent.Id, dbEvent.Table, dbEvent.RecordId)
			case canal.InsertAction:
				fmt.Printf("%s: EventId: %d, Table:%s, RecordId:%d\n", dbEvent.EventType, dbEvent.Id, dbEvent.Table, dbEvent.RecordId)
			case canal.DeleteAction:
				fmt.Printf("DELETE: EventId: %d, Table:%s, RecordId:%d\n", dbEvent.Id, dbEvent.Table, dbEvent.RecordId)
			default:
				fmt.Printf("Unknown action")
			}
		}

	}
	return nil
}

func binlogListener() {
	c := config.LoadConfig()
	binLogCanal, err := getDefaultCanal(c.AppDb)
	if err == nil {
		coords, err := binLogCanal.GetMasterPos()
		if err == nil {
			binLogCanal.SetEventHandler(&binlogHandler{Schema: "app_db", Table: "event_audit"})
			binLogCanal.RunFrom(coords)
		}
	}
}

func getDefaultCanal(config config.MySQLConfig) (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	cfg.User = config.User
	cfg.Password = config.Password
	cfg.Flavor = "mysql"

	cfg.Dump.ExecutionPath = ""

	return canal.NewCanal(cfg)
}
