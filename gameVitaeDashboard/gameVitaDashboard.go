package gameVitaeDashboard

import (
	"fmt"

	"github.com/darvoid/gRPC-slotMachine/conSQLServer"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) RetrieveSessionData(ctx context.Context, req *SessionParameterRequest) (*SessionParameterReply, error) {

	user := "<user>"
	password := "<password>"
	ip := "<ip>"
	port := "<port>"
	database, err := conSQLServer.ConnectToSQLServer(user, password, ip, port)
	if err != nil {
		return nil, err
	}
	var sessions []*sessionRecord

	orderBy := req.GetOrderBy()
	query := database.Table("Common.Session").Select("*").Order(orderBy)
	query_limited := query.Offset(int(req.GetPageIndex())).Limit(int(req.GetPageSize()))
	query_limited.Find(&sessions)

	num := query.RowsAffected
	var result []*SessionRecordSet
	fmt.Printf("%v\n", sessions)

	return &SessionParameterReply{
		Data:      result,
		TotalRows: num,
	}, nil
}

type sessionRecord struct { //table structure
	Id         int64  `gorm:"column:Id";json:"id"`
	User       string `gorm:"column:AspNetUserId";json:"userId"`
	GameId     int64  `gorm:"column:GameId";json:"gameId"`
	DeviceId   int64  `gorm:"column:DeviceId";json:"deviceId"`
	Date       string `gorm:"column:SessionDt";json:"date"`
	ClinicId   int64  `gorm:"column:ClinicHistoryId";json:"clinicId"`
	RowVersion string `gorm:"column:RowVersion";json:"rowVersion"`
}
