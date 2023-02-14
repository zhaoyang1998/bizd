package utils

import (
	"bizd/metion/global"
	"bizd/metion/model"
)

func TranClientStatus(clients []model.Client) {
	for i, client := range clients {
		switch client.Status {
		case global.Unimplemented:
			clients[i].StatusName = global.ClientStatusText[global.Unimplemented]
			break
		case global.InImplemented:
			clients[i].StatusName = global.ClientStatusText[global.InImplemented]
			break
		case global.EndOfImplementation:
			clients[i].StatusName = global.ClientStatusText[global.EndOfImplementation]
			break
		case global.UnPoc:
			clients[i].StatusName = global.ClientStatusText[global.UnPoc]
			break
		case global.InPoc:
			clients[i].StatusName = global.ClientStatusText[global.InPoc]
			break
		case global.EndPoc:
			clients[i].StatusName = global.ClientStatusText[global.EndPoc]
			break
		default:
			clients[i].StatusName = "-"
		}
	}
}
