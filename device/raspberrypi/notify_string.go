// Code generated by "stringer -type=notify"; DO NOT EDIT.

package raspberrypi

import "fmt"

const (
	_notify_name_0 = "notifyTxnotifyRx"
	_notify_name_1 = "notifyButtonPressed"
	_notify_name_2 = "notifyButtonRelease"
	_notify_name_3 = "notifyRemotePressed"
	_notify_name_4 = "notifyRemoteRelease"
	_notify_name_5 = "notifyLogicalAddr"
	_notify_name_6 = "notifyTopology"
	_notify_name_7 = "notifyLogicalAddrLost"
)

var (
	_notify_index_0 = [...]uint8{0, 8, 16}
	_notify_index_1 = [...]uint8{0, 19}
	_notify_index_2 = [...]uint8{0, 19}
	_notify_index_3 = [...]uint8{0, 19}
	_notify_index_4 = [...]uint8{0, 19}
	_notify_index_5 = [...]uint8{0, 17}
	_notify_index_6 = [...]uint8{0, 14}
	_notify_index_7 = [...]uint8{0, 21}
)

func (i notify) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _notify_name_0[_notify_index_0[i]:_notify_index_0[i+1]]
	case i == 4:
		return _notify_name_1
	case i == 8:
		return _notify_name_2
	case i == 16:
		return _notify_name_3
	case i == 32:
		return _notify_name_4
	case i == 64:
		return _notify_name_5
	case i == 128:
		return _notify_name_6
	case i == 32768:
		return _notify_name_7
	default:
		return fmt.Sprintf("notify(%d)", i)
	}
}