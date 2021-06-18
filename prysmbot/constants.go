package main

import "time"

var (
	maxUsersRetrieved = 20000
	maxTimeUnverified = time.Minute * 10

	// Channels and server ids.
	personalTesting = "701148358445760573"
	prysmInternal   = "691473296696410164"
	prysmGeneral    = "476588476393848832"
	prysmRandom     = "696886109589995521"
	prysmGoerli     = "719985056319406182"
	prysmServerId   = "476244492043812875"

	// User id from June 17, 2021 to use as the starting
	// point for fetching guild members.
	fetchGuildMembersAfterId = "854950333309517844"
)
