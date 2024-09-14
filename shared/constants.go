package shared

import "time"

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
	AuthMiddlewareContextKey   = "claims"

	MAX_LEN_20  = 20
	MAX_LEN_50  = 50
	MAX_LEN_80  = 80
	MAX_LEN_100 = 100
	MAX_LEN_500 = 500

	TicketStatus_Unknown    = "Unknown"
	TicketStatus_Pending    = "Pending"
	TicketStatus_Referenced = "Referenced"
	TicketStatus_Answered   = "Answered"
	TicketStatus_Closed     = "Closed"
)
