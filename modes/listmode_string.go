// generated by stringer -type=ListMode; DO NOT EDIT

package modes

import "fmt"

const _ListMode_name = "ListNoneListBanListQuietListExceptListInvex"

var _ListMode_index = [...]uint8{0, 8, 15, 24, 34, 43}

func (i ListMode) String() string {
	if i < 0 || i+1 >= ListMode(len(_ListMode_index)) {
		return fmt.Sprintf("ListMode(%d)", i)
	}
	return _ListMode_name[_ListMode_index[i]:_ListMode_index[i+1]]
}
