// generated by stringer -type=ChannelState; DO NOT EDIT

package tchannel

import "fmt"

const _ChannelState_name = "ChannelClientChannelListeningChannelStartCloseChannelInboundClosedChannelClosed"

var _ChannelState_index = [...]uint8{0, 13, 29, 46, 66, 79}

func (i ChannelState) String() string {
	i -= 1
	if i < 0 || i+1 >= ChannelState(len(_ChannelState_index)) {
		return fmt.Sprintf("ChannelState(%d)", i+1)
	}
	return _ChannelState_name[_ChannelState_index[i]:_ChannelState_index[i+1]]
}
