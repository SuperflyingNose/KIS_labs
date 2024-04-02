package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/ServerReport/{serverId}", HanldleServerReport)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func HanldleServerReport(w http.ResponseWriter, r *http.Request) {

	GetReport(r.PathValue("serverId"))
	f, err := os.Open("ServerReport.csv")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	w.Header().Set("Content-Disposition", "attachement; filename=ServerReport.csv")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	w.Header().Set("Access-Control-Allow-Origin", "*")

	io.Copy(w, reader)
}

func GetReport(serverId string) {
	rc := NewReportCreator("ServerReport.csv")
	rc.generalInfo(serverId)
	rc.MostPopularServer()
	rc.MostActiveMembers()
	rc.Close()
}

type ReportCreator struct {
	db     *Database
	report *os.File
	writer *csv.Writer
}

func NewReportCreator(name string) *ReportCreator {
	f, err := os.Create(name)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	db, err := NewDatabase()

	if err != nil {
		log.Fatalln("failed to create database connection", err)
	}

	w := csv.NewWriter(f)

	return &ReportCreator{db: db, report: f, writer: w}
}

func (rc *ReportCreator) Close() {
	rc.writer.Flush()
	rc.report.Close()
	rc.db.Close()
}

func (rc *ReportCreator) generalInfo(serverId string) error {
	info := [][]string{
		{"Всего чатов", "Всего участников", "Всего сообщений"},
	}
	info = append(info, make([]string, 3))
	chatcount, err := rc.db.GetChatCountByServerId(serverId)
	if err != nil {
		return err
	}
	info[1][0] = strconv.Itoa(chatcount)

	membercount, err := rc.db.GetMemberCountByServerId(serverId)
	if err != nil {
		return err
	}
	info[1][1] = strconv.Itoa(membercount)

	messagecount, err := rc.db.GetMessageCountByServerId(serverId)
	if err != nil {
		return err
	}

	name, err := rc.db.GetServerName(serverId)
	if err != nil {
		return err
	}
	info[1][2] = strconv.Itoa(messagecount)

	rc.writer.Write([]string{fmt.Sprintf("Отчет о сервере %s", name)})
	for _, record := range info {
		if err := rc.writer.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return nil
}

func (rc *ReportCreator) MostPopularServer() error {
	info := make([][]string, 6)
	rows, err := rc.db.GetMostPopularServers()
	if err != nil {
		return nil
	}
	info[0], err = rows.Columns()
	if err != nil {
		return nil
	}
	for i := 1; i < 6; i++ {
		rows.Next()
		info[i] = make([]string, 5)
		rows.Scan(&info[i][0], &info[i][1], &info[i][2], &info[i][3], &info[i][4])
	}
	rc.writer.Write([]string{"Самые популярные каналы"})
	for _, record := range info {
		if err := rc.writer.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return nil
}

func (rc *ReportCreator) MostActiveMembers() error {
	info := make([][]string, 6)
	rows, err := rc.db.GetMostPopularUsers()
	if err != nil {
		return nil
	}
	info[0], err = rows.Columns()
	if err != nil {
		return nil
	}
	for i := 1; i < 6; i++ {
		rows.Next()
		info[i] = make([]string, 4)
		rows.Scan(&info[i][0], &info[i][1], &info[i][2], &info[i][3])
	}
	rc.writer.Write([]string{"Самые активные пользователи"})
	for _, record := range info {
		if err := rc.writer.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return nil
}
