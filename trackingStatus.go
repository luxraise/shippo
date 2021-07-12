package shippo

import (
	"encoding/json"
	"time"
)

type TrackingStatus struct {
	ObjectMeta

	Status        string    `json:"status"`
	StatusDetails string    `json:"status_details"`
	StatusDate    time.Time `json:"status_date"`
}

type MaybeTrackingStatus struct {
	TrackingStatus
}

func (m *MaybeTrackingStatus) UnmarshalJSON(bs []byte) (err error) {
	if isUnknown(bs) {
		return
	}

	var ts TrackingStatus
	if err = json.Unmarshal(bs, &ts); err != nil {
		return
	}

	*m = MaybeTrackingStatus{ts}
	return
}
