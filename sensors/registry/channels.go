package registry

import "sensors/data"

//Channels has all channels
type Channels struct {
	ChannelNewUsers data.ChanNewUsers
}

//NewChannels creates a Channels struct
func NewChannels() *Channels {
	chanNewUsers := make(chan data.UserContract, 10)

	return &Channels{
		ChannelNewUsers: data.ChanNewUsers{
			Channel: chanNewUsers,
			Reader:  chanNewUsers,
			Writer:  chanNewUsers,
		},
	}
}
