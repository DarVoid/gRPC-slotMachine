package gameVitaeDashboard

import (
	"encoding/json"
	"fmt"

	"github.com/darvoid/gRPC-slotMachine/connectDB"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) RetrieveSessionData(ctx context.Context, req *SessionParameterRequest) (*SessionParameterReply, error) {
	database, err := connectDB.Connect()
	if err != nil {
		fmt.Println(err)
	}
	var sessions []sessionRecord

	orderBy := req.GetOrderBy() //to do: concatenar com req get asc
	query := database.Table("Common.Session").Select("").Order(orderBy)
	//query_limited := query.Offset(int(req.GetPageIndex())).Limit(int(req.GetPageSize()))
	query.Scan(&sessions)
	//b, err := json.Marshal(sessions)
	if err != nil {
		fmt.Println(err)
	}
	num := query.RowsAffected
	result := make([]*SessionRecordSet, 0)
	for i, val := range sessions {
		json.Marshal(num)
		json.Marshal(i)
		fmt.Println(i)
		fmt.Println(num)
		b := &SessionRecordSet{
			Id:       val.Id,
			User:     val.User,
			GameId:   val.GameId,
			DeviceId: val.DeviceId,
			Date:     val.Date,
			ClinicId: val.ClinicId}

		result = append(result, b)
	}
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
