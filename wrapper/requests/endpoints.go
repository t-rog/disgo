package requests

// Discord API Endpoints
const (
	EndpointBaseURL                                = "https://discord.com/api/v9/"
	EndpointGetGlobalApplicationCommands           = "applications/{application.id}/commands"
	EndpointCreateGlobalApplicationCommand         = "applications/{application.id}/commands"
	EndpointGetGlobalApplicationCommand            = "applications/{application.id}/commands/{command.id}"
	EndpointEditGlobalApplicationCommand           = "applications/{application.id}/commands/{command.id}"
	EndpointDeleteGlobalApplicationCommand         = "applications/{application.id}/commands/{command.id}"
	EndpointBulkOverwriteGlobalApplicationCommands = "applications/{application.id}/commands"
	EndpointGetGuildApplicationCommands            = "applications/{application.id}/guilds/{guild.id}/commands"
	EndpointCreateGuildApplicationCommand          = "applications/{application.id}/guilds/{guild.id}/commands"
	EndpointGetGuildApplicationCommand             = "applications/{application.id}/guilds/{guild.id}/commands/{command.id}"
	EndpointEditGuildApplicationCommand            = "applications/{application.id}/guilds/{guild.id}/commands/{command.id}"
	EndpointDeleteGuildApplicationCommand          = "applications/{application.id}/guilds/{guild.id}/commands/{command.id}"
	EndpointBulkOverwriteGuildApplicationCommands  = "applications/{application.id}/guilds/{guild.id}/commands"
	EndpointGetGuildApplicationCommandPermissions  = "applications/{application.id}/guilds/{guild.id}/commands/permissions"
	EndpointGetApplicationCommandPermissions       = "applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions"
	EndpointEditApplicationCommandPermissions      = "applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions"
	EndpointBatchEditApplicationCommandPermissions = "applications/{application.id}/guilds/{guild.id}/commands/permissions"
	EndpointCreateInteractionResponse              = "interactions/{interaction.id}/{interaction.token}/callback"
	EndpointGetOriginalInteractionResponse         = "webhooks/{application.id}/{interaction.token}/messages/@original"
	EndpointEditOriginalInteractionResponse        = "webhooks/{application.id}/{interaction.token}/messages/@original"
	EndpointDeleteOriginalInteractionResponse      = "webhooks/{application.id}/{interaction.token}/messages/@original"
	EndpointCreateFollowupMessage                  = "webhooks/{application.id}/{interaction.token}"
	EndpointGetFollowupMessage                     = "webhooks/{application.id}/{interaction.token}/messages/{message.id}"
	EndpointEditFollowupMessage                    = "webhooks/{application.id}/{interaction.token}/messages/{message.id}"
	EndpointDeleteFollowupMessage                  = "webhooks/{application.id}/{interaction.token}/messages/{message.id}"
	EndpointGetGuildAuditLog                       = "guilds/{guild.id}/audit-logs"
	EndpointGetChannel                             = "channels/{channel.id}"
	EndpointModifyChannel                          = "channels/{channel.id}"
	EndpointDeleteCloseChannel                     = "channels/{channel.id}"
	EndpointGetChannelMessages                     = "channels/{channel.id}/messages"
	EndpointGetChannelMessage                      = "channels/{channel.id}/messages/{message.id}"
	EndpointCreateMessage                          = "channels/{channel.id}/messages"
	EndpointCrosspostMessage                       = "channels/{channel.id}/messages/{message.id}/crosspost"
	EndpointCreateReaction                         = "channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me"
	EndpointDeleteOwnReaction                      = "channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me"
	EndpointDeleteUserReaction                     = "channels/{channel.id}/messages/{message.id}/reactions/{emoji}/{user.id}"
	EndpointGetReactions                           = "channels/{channel.id}/messages/{message.id}/reactions/{emoji}"
	EndpointDeleteAllReactions                     = "channels/{channel.id}/messages/{message.id}/reactions"
	EndpointDeleteAllReactionsforEmoji             = "channels/{channel.id}/messages/{message.id}/reactions/{emoji}"
	EndpointEditMessage                            = "channels/{channel.id}/messages/{message.id}"
	EndpointDeleteMessage                          = "channels/{channel.id}/messages/{message.id}"
	EndpointBulkDeleteMessages                     = "channels/{channel.id}/messages/bulk-delete"
	EndpointEditChannelPermissions                 = "channels/{channel.id}/permissions/{overwrite.id}"
	EndpointGetChannelInvites                      = "channels/{channel.id}/invites"
	EndpointCreateChannelInvite                    = "channels/{channel.id}/invites"
	EndpointDeleteChannelPermission                = "channels/{channel.id}/permissions/{overwrite.id}"
	EndpointFollowNewsChannel                      = "channels/{channel.id}/followers"
	EndpointTriggerTypingIndicator                 = "channels/{channel.id}/typing"
	EndpointGetPinnedMessages                      = "channels/{channel.id}/pins"
	EndpointPinMessage                             = "channels/{channel.id}/pins/{message.id}"
	EndpointUnpinMessage                           = "channels/{channel.id}/pins/{message.id}"
	EndpointGroupDMAddRecipient                    = "channels/{channel.id}/recipients/{user.id}"
	EndpointGroupDMRemoveRecipient                 = "channels/{channel.id}/recipients/{user.id}"
	EndpointStartThreadfromMessage                 = "channels/{channel.id}/messages/{message.id}/threads"
	EndpointStartThreadwithoutMessage              = "channels/{channel.id}/threads"
	EndpointStartThreadinForumChannel              = "channels/{channel.id}/threads"
	EndpointJoinThread                             = "channels/{channel.id}/thread-members/@me"
	EndpointAddThreadMember                        = "channels/{channel.id}/thread-members/{user.id}"
	EndpointLeaveThread                            = "channels/{channel.id}/thread-members/@me"
	EndpointRemoveThreadMember                     = "channels/{channel.id}/thread-members/{user.id}"
	EndpointGetThreadMember                        = "channels/{channel.id}/thread-members/{user.id}"
	EndpointListThreadMembers                      = "channels/{channel.id}/thread-members"
	EndpointListActiveChannelThreads               = "channels/{channel.id}/threads/active"
	EndpointListPublicArchivedThreads              = "channels/{channel.id}/threads/archived/public"
	EndpointListPrivateArchivedThreads             = "channels/{channel.id}/threads/archived/private"
	EndpointListJoinedPrivateArchivedThreads       = "channels/{channel.id}/users/@me/threads/archived/private"
	EndpointListGuildEmojis                        = "guilds/{guild.id}/emojis"
	EndpointGetGuildEmoji                          = "guilds/{guild.id}/emojis/{emoji.id}"
	EndpointCreateGuildEmoji                       = "guilds/{guild.id}/emojis"
	EndpointModifyGuildEmoji                       = "guilds/{guild.id}/emojis/{emoji.id}"
	EndpointDeleteGuildEmoji                       = "guilds/{guild.id}/emojis/{emoji.id}"
	EndpointListScheduledEventsforGuild            = "guilds/{guild.id}/scheduled-events"
	EndpointCreateGuildScheduledEvent              = "guilds/{guild.id}/scheduled-events"
	EndpointGetGuildScheduledEvent                 = "guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}"
	EndpointModifyGuildScheduledEvent              = "guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}"
	EndpointDeleteGuildScheduledEvent              = "guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}"
	EndpointGetGuildScheduledEventUsers            = "guilds/{guild.id}/scheduled-events/{guild_scheduled_event.id}/users"
	EndpointGetGuildTemplate                       = "guilds/templates/{template.code}"
	EndpointCreateGuildfromGuildTemplate           = "guilds/templates/{template.code}"
	EndpointGetGuildTemplates                      = "guilds/{guild.id}/templates"
	EndpointCreateGuildTemplate                    = "guilds/{guild.id}/templates"
	EndpointSyncGuildTemplate                      = "guilds/{guild.id}/templates/{template.code}"
	EndpointModifyGuildTemplate                    = "guilds/{guild.id}/templates/{template.code}"
	EndpointDeleteGuildTemplate                    = "guilds/{guild.id}/templates/{template.code}"
	EndpointGetGuild                               = "guilds/{guild.id}"
	EndpointGetGuildPreview                        = "guilds/{guild.id}/preview"
	EndpointModifyGuild                            = "guilds/{guild.id}"
	EndpointDeleteGuild                            = "guilds/{guild.id}"
	EndpointGetGuildChannels                       = "guilds/{guild.id}/channels"
	EndpointCreateGuildChannel                     = "guilds/{guild.id}/channels"
	EndpointModifyGuildChannelPositions            = "guilds/{guild.id}/channels"
	EndpointListActiveGuildThreads                 = "guilds/{guild.id}/threads/active"
	EndpointGetGuildMember                         = "guilds/{guild.id}/members/{user.id}"
	EndpointListGuildMembers                       = "guilds/{guild.id}/members"
	EndpointSearchGuildMembers                     = "guilds/{guild.id}/members/search"
	EndpointAddGuildMember                         = "guilds/{guild.id}/members/{user.id}"
	EndpointModifyGuildMember                      = "guilds/{guild.id}/members/{user.id}"
	EndpointModifyCurrentMember                    = "guilds/{guild.id}/members/@me"
	EndpointModifyCurrentUserNick                  = "guilds/{guild.id}/members/@me/nick"
	EndpointAddGuildMemberRole                     = "guilds/{guild.id}/members/{user.id}/roles/{role.id}"
	EndpointRemoveGuildMemberRole                  = "guilds/{guild.id}/members/{user.id}/roles/{role.id}"
	EndpointRemoveGuildMember                      = "guilds/{guild.id}/members/{user.id}"
	EndpointGetGuildBans                           = "guilds/{guild.id}/bans"
	EndpointGetGuildBan                            = "guilds/{guild.id}/bans/{user.id}"
	EndpointCreateGuildBan                         = "guilds/{guild.id}/bans/{user.id}"
	EndpointRemoveGuildBan                         = "guilds/{guild.id}/bans/{user.id}"
	EndpointGetGuildRoles                          = "guilds/{guild.id}/roles"
	EndpointCreateGuildRole                        = "guilds/{guild.id}/roles"
	EndpointModifyGuildRolePositions               = "guilds/{guild.id}/roles"
	EndpointModifyGuildRole                        = "guilds/{guild.id}/roles/{role.id}"
	EndpointDeleteGuildRole                        = "guilds/{guild.id}/roles/{role.id}"
	EndpointGetGuildPruneCount                     = "guilds/{guild.id}/prune"
	EndpointBeginGuildPrune                        = "guilds/{guild.id}/prune"
	EndpointGetGuildVoiceRegions                   = "guilds/{guild.id}/regions"
	EndpointGetGuildInvites                        = "guilds/{guild.id}/invites"
	EndpointGetGuildIntegrations                   = "guilds/{guild.id}/integrations"
	EndpointDeleteGuildIntegration                 = "guilds/{guild.id}/integrations/{integration.id}"
	EndpointGetGuildWidgetSettings                 = "guilds/{guild.id}/widget"
	EndpointModifyGuildWidget                      = "guilds/{guild.id}/widget"
	EndpointGetGuildWidget                         = "guilds/{guild.id}/widget.json"
	EndpointGetGuildVanityURL                      = "guilds/{guild.id}/vanity-url"
	EndpointGetGuildWidgetImage                    = "guilds/{guild.id}/widget.png"
	EndpointGetGuildWelcomeScreen                  = "guilds/{guild.id}/welcome-screen"
	EndpointModifyGuildWelcomeScreen               = "guilds/{guild.id}/welcome-screen"
	EndpointModifyCurrentUserVoiceState            = "guilds/{guild.id}/voice-states/@me"
	EndpointModifyUserVoiceState                   = "guilds/{guild.id}/voice-states/{user.id}"
	EndpointGetInvite                              = "invites/{invite.code}"
	EndpointDeleteInvite                           = "invites/{invite.code}"
	EndpointGetSticker                             = "stickers/{sticker.id}"
	EndpointListGuildStickers                      = "guilds/{guild.id}/stickers"
	EndpointGetGuildSticker                        = "guilds/{guild.id}/stickers/{sticker.id}"
	EndpointCreateGuildSticker                     = "guilds/{guild.id}/stickers"
	EndpointModifyGuildSticker                     = "guilds/{guild.id}/stickers/{sticker.id}"
	EndpointDeleteGuildSticker                     = "guilds/{guild.id}/stickers/{sticker.id}"
	EndpointGetUser                                = "users/{user.id}"
	EndpointGetCurrentUserGuildMember              = "users/@me/guilds/{guild.id}/member"
	EndpointLeaveGuild                             = "users/@me/guilds/{guild.id}"
	EndpointCreateWebhook                          = "channels/{channel.id}/webhooks"
	EndpointGetChannelWebhooks                     = "channels/{channel.id}/webhooks"
	EndpointGetGuildWebhooks                       = "guilds/{guild.id}/webhooks"
	EndpointGetWebhook                             = "webhooks/{webhook.id}"
	EndpointGetWebhookwithToken                    = "webhooks/{webhook.id}/{webhook.token}"
	EndpointModifyWebhook                          = "webhooks/{webhook.id}"
	EndpointModifyWebhookwithToken                 = "webhooks/{webhook.id}/{webhook.token}"
	EndpointDeleteWebhook                          = "webhooks/{webhook.id}"
	EndpointDeleteWebhookwithToken                 = "webhooks/{webhook.id}/{webhook.token}"
	EndpointExecuteWebhook                         = "webhooks/{webhook.id}/{webhook.token}"
	EndpointExecuteSlackCompatibleWebhook          = "webhooks/{webhook.id}/{webhook.token}/slack"
	EndpointExecuteGitHubCompatibleWebhook         = "webhooks/{webhook.id}/{webhook.token}/github"
	EndpointGetWebhookMessage                      = "webhooks/{webhook.id}/{webhook.token}/messages/{message.id}"
	EndpointEditWebhookMessage                     = "webhooks/{webhook.id}/{webhook.token}/messages/{message.id}"
	EndpointDeleteWebhookMessage                   = "webhooks/{webhook.id}/{webhook.token}/messages/{message.id}"
)
