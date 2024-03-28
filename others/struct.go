package others

import (
	"fmt"
	"time"
)

type Video struct {
	creator     string
	channelName string
	created     string
	createdTime string
	duration    time.Duration
}

func (v *Video) getVideoDetails() string {
	videoDetails := Video{
		v.creator,
		v.channelName,
		v.created,
		v.createdTime,
		v.duration,
	}
	return fmt.Sprintf("Video was created by %s by %s %s, the channel name is %s and the duration is %s", videoDetails.creator, videoDetails.createdTime, videoDetails.created, videoDetails.creator, videoDetails.duration)
}

func MainForStruct() {
	myVideo := Video{
		creator:     "Stanley",
		channelName: "StanCode",
		created:     time.Now().Format("2006-01-02"),
		createdTime: time.Now().Format("15:05"),
		duration:    5 * time.Minute,
	}
	fmt.Println(myVideo.getVideoDetails())
}
