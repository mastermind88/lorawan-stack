// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mac_test

import (
	"context"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	. "go.thethings.network/lorawan-stack/v3/pkg/networkserver/mac"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestNeedsPingSlotChannelReq(t *testing.T) {
	for _, tc := range []struct {
		Name        string
		InputDevice *ttnpb.EndDevice
		Needs       bool
	}{
		{
			Name:        "no MAC state",
			InputDevice: &ttnpb.EndDevice{},
		},
		{
			Name: "current(data-rate-index:1,frequency:123),desired(data-rate-index:1,frequency:123)",
			InputDevice: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{
						PingSlotDataRateIndexValue: &ttnpb.DataRateIndexValue{Value: ttnpb.DataRateIndex_DATA_RATE_1},
						PingSlotFrequency:          123,
					},
					DesiredParameters: &ttnpb.MACParameters{
						PingSlotDataRateIndexValue: &ttnpb.DataRateIndexValue{Value: ttnpb.DataRateIndex_DATA_RATE_1},
						PingSlotFrequency:          123,
					},
				},
			},
		},
		{
			Name: "current(data-rate-index:1,frequency:123),desired(data-rate-index:2,frequency:123)",
			InputDevice: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{
						PingSlotDataRateIndexValue: &ttnpb.DataRateIndexValue{Value: ttnpb.DataRateIndex_DATA_RATE_1},
						PingSlotFrequency:          123,
					},
					DesiredParameters: &ttnpb.MACParameters{
						PingSlotDataRateIndexValue: &ttnpb.DataRateIndexValue{Value: ttnpb.DataRateIndex_DATA_RATE_2},
						PingSlotFrequency:          123,
					},
				},
			},
			Needs: true,
		},
		{
			Name: "current(data-rate-index:1,frequency:123),desired(data-rate-index:1,frequency:124)",
			InputDevice: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{
						PingSlotDataRateIndexValue: &ttnpb.DataRateIndexValue{Value: ttnpb.DataRateIndex_DATA_RATE_1},
						PingSlotFrequency:          123,
					},
					DesiredParameters: &ttnpb.MACParameters{
						PingSlotDataRateIndexValue: &ttnpb.DataRateIndexValue{Value: ttnpb.DataRateIndex_DATA_RATE_1},
						PingSlotFrequency:          124,
					},
				},
			},
			Needs: true,
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := ttnpb.Clone(tc.InputDevice)
				res := DeviceNeedsPingSlotChannelReq(dev)
				if tc.Needs {
					a.So(res, should.BeTrue)
				} else {
					a.So(res, should.BeFalse)
				}
				a.So(dev, should.Resemble, tc.InputDevice)
			},
		})
	}
}

func TestHandlePingSlotChannelAns(t *testing.T) {
	for _, tc := range []struct {
		Name             string
		Device, Expected *ttnpb.EndDevice
		Payload          *ttnpb.MACCommand_PingSlotChannelAns
		Events           events.Builders
		Error            error
	}{
		{
			Name: "nil payload",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Error: ErrNoPayload,
		},
		{
			Name: "no request",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Payload: &ttnpb.MACCommand_PingSlotChannelAns{
				FrequencyAck:     true,
				DataRateIndexAck: true,
			},
			Events: events.Builders{
				EvtReceivePingSlotChannelAnswer.With(events.WithData(&ttnpb.MACCommand_PingSlotChannelAns{
					FrequencyAck:     true,
					DataRateIndexAck: true,
				})),
			},
			Error: ErrRequestNotFound.WithAttributes("cid", ttnpb.MACCommandIdentifier_CID_PING_SLOT_CHANNEL),
		},
		{
			Name: "both ack",
			Device: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{
						(&ttnpb.MACCommand_PingSlotChannelReq{
							Frequency:     42,
							DataRateIndex: 43,
						}).MACCommand(),
					},
					CurrentParameters: &ttnpb.MACParameters{},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Expected: &ttnpb.EndDevice{
				MacState: &ttnpb.MACState{
					PendingRequests: []*ttnpb.MACCommand{},
					CurrentParameters: &ttnpb.MACParameters{
						PingSlotDataRateIndexValue: &ttnpb.DataRateIndexValue{Value: 43},
						PingSlotFrequency:          42,
					},
					DesiredParameters: &ttnpb.MACParameters{},
				},
			},
			Payload: &ttnpb.MACCommand_PingSlotChannelAns{
				FrequencyAck:     true,
				DataRateIndexAck: true,
			},
			Events: events.Builders{
				EvtReceivePingSlotChannelAnswer.With(events.WithData(&ttnpb.MACCommand_PingSlotChannelAns{
					FrequencyAck:     true,
					DataRateIndexAck: true,
				})),
			},
		},
	} {
		tc := tc
		test.RunSubtest(t, test.SubtestConfig{
			Name:     tc.Name,
			Parallel: true,
			Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
				dev := ttnpb.Clone(tc.Device)

				evs, err := HandlePingSlotChannelAns(ctx, dev, tc.Payload)
				if tc.Error != nil && !a.So(err, should.EqualErrorOrDefinition, tc.Error) ||
					tc.Error == nil && !a.So(err, should.BeNil) {
					t.FailNow()
				}
				a.So(dev, should.Resemble, tc.Expected)
				a.So(evs, should.ResembleEventBuilders, tc.Events)
			},
		})
	}
}
