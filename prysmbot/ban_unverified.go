package main

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

const (
	monitorUnverifiedInterval = time.Minute * 1
)

// Monitors for users that have not completed a captcha after a
// certain period of time and kicks them from the server.
func monitorForUnverifiedUsers(ctx context.Context, dg *discordgo.Session) {
	ticker := time.NewTicker(monitorUnverifiedInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			members, err := dg.GuildMembers(prysmServerId, fetchGuildMembersAfterId, maxUsersRetrieved)
			if err != nil {
				log.WithError(err).Error("Could not get guild members")
				return
			}
			kicked := 0
			for _, mem := range members {
				joinedAt, err := mem.JoinedAt.Parse()
				if err != nil {
					log.WithError(err).Error("Could not check member timestamp")
					continue
				}
				if !shouldBeKicked(mem.Roles, joinedAt) {
					continue
				}
				if err := dg.GuildMemberDeleteWithReason(
					prysmServerId, mem.User.ID, "User did not complete verification within 10 minutes",
				); err != nil {
					log.WithError(err).Error("Could not kick member")
				}
				log.WithFields(logrus.Fields{
					"userId":   mem.User.ID,
					"username": mem.User.Username,
				}).Info("Kicked user for being unverified too long")
				kicked++
			}

			log.WithFields(logrus.Fields{
				"membersChecked": len(members),
				"lastID":         fetchGuildMembersAfterId,
				"kicked":         kicked,
			}).Debug("Checked unverified members to kick")

			if len(members) == maxUsersRetrieved {
				log.Warn("Max users returned from API. There may be new recent members that are not checked!")
			}
		case <-ctx.Done():
			return
		}
	}
}

func shouldBeKicked(roles []string, joinedAt time.Time) bool {
	kickDuration := time.Now().Add(-maxTimeUnverified)
	unverifiedForTooLong := joinedAt.Before(kickDuration)
	return len(roles) == 0 && unverifiedForTooLong
}
