package v1

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestMarshal(t *testing.T) {
	testData := Frame{
		CameraId:  "1",
		Timestamp: uint64(time.Now().Unix()),
		Object: []*Object{{
			ObjectId:     1,
			TrackingId:   1,
			ObjectLabel:  "123",
			DetectConf:   0,
			TrackingConf: 0,
		}},
		ImagePath: "xx",
		JobId:     "1",
		ModelType: "xx",
		NumObj:    0,
	}
	b, err := proto.Marshal(&testData)
	require.NoError(t, err)

	result := base64.StdEncoding.EncodeToString(b)
	fmt.Println(result)
}
