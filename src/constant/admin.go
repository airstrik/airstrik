package constant

type UserStatus string

const (
	UserActive      UserStatus = "Active"
	UserInActive    UserStatus = "InActive"
	UserDeleted     UserStatus = "Deleted"
	UserNotVerified UserStatus = "NotVerified"
	UserInvited     UserStatus = "Invited"
)

type GroupStatus string

const (
	GroupActive   UserStatus = "Active"
	GroupInActive UserStatus = "InActive"
	GroupDeleted  UserStatus = "Deleted"
)
