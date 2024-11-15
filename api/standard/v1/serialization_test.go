package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

const testJson = `{
  "version": "2.0",
  "task": {
    "workspace_id": "1022",
    "zone_id": "200",
    "box_id": "302203231522101",
    "model_name": "gas_station",
    "task_id": "902203231529"
  },
  "events": {
    "camera_id": "file:///work/config/../video/baidu_189_Baidu_0139_0.mp4",
    "timestamp": 1678885978209,
    "image": {
      "path": "0_248_1920x1080_1678885978157564000.jpg",
      "width": 1920,
      "height": 1080
    },
    "objects": [
      {
        "track_id": 1,
        "object_label": "person",
        "detect_conf": 0.91007620096206665,
        "track_conf": 1.0,
        "object_box": {
          "x": 0.8203125,
          "y": 0.21666669845581055,
          "width": 0.090625002980232239,
          "height": 0.3722221851348877
        },
        "attribute": [
          {
            "name": "workclothes",
            "value": "yes"
          }
        ]
      },
      {
        "track_id": 2,
        "object_label": "extinguisher",
        "detect_conf": 0.91007620096206665,
        "object_box": {
          "x": 0.8203125,
          "y": 0.21666669845581055,
          "width": 0.090625002980232239,
          "height": 0.3722221851348877
        }
      },
      {
        "track_id": 3,
        "object_label": "tanker",
        "detect_conf": 0.91007620096206665,
        "track_conf": 1.0,
        "object_box": {
          "x": 0.8203125,
          "y": 0.21666669845581055,
          "width": 0.090625002980232239,
          "height": 0.3722221851348877
        },
        "attribute": [
          {
            "name": "chain",
            "value": "yes"
          },
          {
            "name": "OC",
            "value": "IN"
          }
        ]
      }
    ]
  }
}`

func TestJsonUnmarshal(t *testing.T) {
	var output AlgoOutput
	err := protojson.Unmarshal([]byte(testJson), &output)
	require.NoError(t, err)
	t.Log(&output)
}
