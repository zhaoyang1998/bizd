package utils

import (
	"bizd/metion/global"
	"bizd/metion/model"
)

func TranClientStatus(clients []model.Client) {
	for i, client := range clients {
		switch client.Status {
		case global.UnImplemented:
			clients[i].StatusName = global.ClientAndPointPositionStatusText[global.UnImplemented]
			break
		case global.InImplemented:
			clients[i].StatusName = global.ClientAndPointPositionStatusText[global.InImplemented]
			break
		case global.EndOfImplementation:
			clients[i].StatusName = global.ClientAndPointPositionStatusText[global.EndOfImplementation]
			break
		case global.UnPoc:
			clients[i].StatusName = global.ClientAndPointPositionStatusText[global.UnPoc]
			break
		case global.InPoc:
			clients[i].StatusName = global.ClientAndPointPositionStatusText[global.InPoc]
			break
		case global.EndPoc:
			clients[i].StatusName = global.ClientAndPointPositionStatusText[global.EndPoc]
			break
		default:
			clients[i].StatusName = "-"
		}
	}
}

func TranUserType(users []model.User) {
	for i, user := range users {
		switch user.Type {
		case global.Delivery:
			users[i].TypeName = global.UserTypeText[global.Delivery]
			break
		case global.PM:
			users[i].TypeName = global.UserTypeText[global.PM]
			break
		default:
			users[i].TypeName = "-"
		}
	}
}

func TranPointPositionStatus(pps []model.PointPosition) {
	for i, pp := range pps {
		switch *pp.Status {
		case global.UnResearched:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.UnResearched]
			break
		case global.InResearched:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.InResearched]
			break
		case global.EndOfResearched:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.EndOfResearched]
			break
		case global.UnImplemented:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.UnImplemented]
			break
		case global.InImplemented:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.InImplemented]
			break
		case global.EndOfImplementation:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.EndOfImplementation]
			break
		case global.UnPoc:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.UnPoc]
			break
		case global.InPoc:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.InPoc]
			break
		case global.EndPoc:
			pps[i].StatusName = global.ClientAndPointPositionStatusText[global.EndPoc]
			break
		default:
			pps[i].StatusName = "-"
		}
	}
}
