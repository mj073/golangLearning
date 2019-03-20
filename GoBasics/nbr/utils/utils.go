package utils

import (
	"time"
)

type CacheStruct struct{
	Id	int
	IngressIp	string
	IngressPort	string
	EgressIp	string
	EgressPort 	string
	HttpTimestamp	time.Time
	Host	string
	HttpMethod	string
	Url 	string
	HttpVersion	string
	HttpStatusCode	int
	ContentLength	int
	CacheCode	string
}

type UUIDStruct struct{
	Id 	int
	LocationId	string
	UUID	string
	UserIp	string
	UserPort	string
	NATIp	string
	NATPort	string
	DestinationIp	string
	DestinationPort	string
	UUIDTimestamp	time.Time
}

type FirewallStruct struct{
	Id	int
	IngressIp	string
	IngressPort	string
	PublicIp	string
	PublicPort	string
	LogTimestamp	time.Time
}