package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	db      *sql.DB
	ctx     context.Context
	timeout time.Duration
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:SuperN0se@localhost:5432/discordminus?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db, ctx: context.Background(), timeout: time.Duration(2) * time.Hour}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetChatCountByServerId(serverId string) (int, error) {
	var chatCount int
	query := "SELECT count(*) FROM public.\"Channel\" WHERE \"Channel\".\"serverId\" = $1"

	c, cancel := context.WithTimeout(d.ctx, d.timeout)
	defer cancel()
	err := d.db.QueryRowContext(c, query, serverId).Scan(&chatCount)
	if err != nil {
		return 0, err
	}

	return chatCount, nil
}

func (d *Database) GetMemberCountByServerId(serverId string) (int, error) {
	var memberCount int
	query := "SELECT count(*) FROM public.\"Member\" WHERE \"Member\".\"serverId\" = $1"

	c, cancel := context.WithTimeout(d.ctx, d.timeout)
	defer cancel()
	err := d.db.QueryRowContext(c, query, serverId).Scan(&memberCount)
	if err != nil {
		return 0, err
	}

	return memberCount, nil
}

func (d *Database) GetMessageCountByServerId(serverId string) (int, error) {
	var chatCount int
	query := "SELECT count(*) FROM public.\"Message\" JOIN public.\"Channel\" on \"Message\".\"channelId\" = \"Channel\".\"id\" WHERE \"Channel\".\"serverId\" = $1"

	c, cancel := context.WithTimeout(d.ctx, d.timeout)
	defer cancel()
	err := d.db.QueryRowContext(c, query, serverId).Scan(&chatCount)
	if err != nil {
		return 0, err
	}

	return chatCount, nil
}

func (d *Database) GetServerName(serverId string) (string, error) {
	var name string
	query := `SELECT "Server"."name" FROM public."Server" WHERE "Server"."id" = $1`
	err := d.db.QueryRow(query, serverId).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (d *Database) GetMostPopularServers() (*sql.Rows, error) {
	query := `SELECT  "channelCount"."serverName" AS "Название сервера", "memberCount"."count" AS "Кол-во участников", "channelCount"."count" AS "Кол-во каналов", "messageCount"."count" AS "Кол-во сообщений", "creatorName"."creator" AS "Создатель"	FROM (SELECT "Server"."name" AS "serverName", "Profile"."name" AS "creator" FROM "Server" JOIN ("Member" JOIN "Profile" ON "Member"."profileId" = "Profile"."id") ON "Member"."profileId" = "Server"."profileId" GROUP BY "Server"."name", "Profile"."name") AS "creatorName" JOIN ((SELECT "Server"."name" AS "serverName", count("Member"."id") AS "count" FROM "Server" JOIN "Member" ON "Member"."serverId" = "Server"."id" GROUP BY "Server"."name") AS "memberCount" JOIN	((SELECT "Server"."name" AS "serverName", count(*) AS "count" FROM "Channel" JOIN "Server" ON "Server"."id" = "Channel"."serverId" GROUP BY "Server"."name") AS "channelCount" JOIN	(SELECT "Server"."name" AS "serverName", count("Message"."content") AS "count" FROM "Server" JOIN "Channel" ON "Channel"."serverId" = "Server"."id"	JOIN "Message" ON "Message"."channelId" = "Channel"."id" GROUP BY  "Server"."name") AS "messageCount" ON "messageCount"."serverName" = "channelCount"."serverName")	ON "memberCount"."serverName" = "channelCount"."serverName") ON "creatorName"."serverName" = "channelCount"."serverName" ORDER BY "memberCount"."count" desc, "messageCount"."count" desc LIMIT 5`
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (d *Database) GetMostPopularUsers() (*sql.Rows, error) {
	query := `SELECT "Profile"."name", count("Message"."content"), string_agg(distinct "Member"."role"::text, ', '),
membersServer."name"
FROM "Profile" 
 JOIN 
 ("Member" 
  JOIN "Message" 
  ON "Member"."id" = "Message"."memberId") 
 ON "Profile"."id" = "Member"."profileId"
JOIN
(SELECT "Profile"."id" as "id", string_agg(distinct "Server"."name"::text, ', ') as "name" FROM 
 "Profile" JOIN (
"Member" JOIN "Server" ON "Member"."serverId" = "Server"."id")
ON "Profile"."id" = "Member"."profileId" GROUP BY "Profile"."id") AS membersServer
 ON "Profile"."id" = membersServer."id"
 GROUP BY membersServer."name", "Profile"."name"`
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
