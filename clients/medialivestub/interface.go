// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7

package medialivestub

import (
	"github.com/aws/aws-sdk-go/service/medialive"
	"go.uber.org/cadence/workflow"

	"github.com/banzaicloud/cadence-aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptInputDeviceTransfer(ctx workflow.Context, input *medialive.AcceptInputDeviceTransferInput) (*medialive.AcceptInputDeviceTransferOutput, error)
	AcceptInputDeviceTransferAsync(ctx workflow.Context, input *medialive.AcceptInputDeviceTransferInput) *MediaLiveAcceptInputDeviceTransferFuture

	BatchDelete(ctx workflow.Context, input *medialive.BatchDeleteInput) (*medialive.BatchDeleteOutput, error)
	BatchDeleteAsync(ctx workflow.Context, input *medialive.BatchDeleteInput) *MediaLiveBatchDeleteFuture

	BatchStart(ctx workflow.Context, input *medialive.BatchStartInput) (*medialive.BatchStartOutput, error)
	BatchStartAsync(ctx workflow.Context, input *medialive.BatchStartInput) *MediaLiveBatchStartFuture

	BatchStop(ctx workflow.Context, input *medialive.BatchStopInput) (*medialive.BatchStopOutput, error)
	BatchStopAsync(ctx workflow.Context, input *medialive.BatchStopInput) *MediaLiveBatchStopFuture

	BatchUpdateSchedule(ctx workflow.Context, input *medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error)
	BatchUpdateScheduleAsync(ctx workflow.Context, input *medialive.BatchUpdateScheduleInput) *MediaLiveBatchUpdateScheduleFuture

	CancelInputDeviceTransfer(ctx workflow.Context, input *medialive.CancelInputDeviceTransferInput) (*medialive.CancelInputDeviceTransferOutput, error)
	CancelInputDeviceTransferAsync(ctx workflow.Context, input *medialive.CancelInputDeviceTransferInput) *MediaLiveCancelInputDeviceTransferFuture

	CreateChannel(ctx workflow.Context, input *medialive.CreateChannelInput) (*medialive.CreateChannelOutput, error)
	CreateChannelAsync(ctx workflow.Context, input *medialive.CreateChannelInput) *MediaLiveCreateChannelFuture

	CreateInput(ctx workflow.Context, input *medialive.CreateInputInput) (*medialive.CreateInputOutput, error)
	CreateInputAsync(ctx workflow.Context, input *medialive.CreateInputInput) *MediaLiveCreateInputFuture

	CreateInputSecurityGroup(ctx workflow.Context, input *medialive.CreateInputSecurityGroupInput) (*medialive.CreateInputSecurityGroupOutput, error)
	CreateInputSecurityGroupAsync(ctx workflow.Context, input *medialive.CreateInputSecurityGroupInput) *MediaLiveCreateInputSecurityGroupFuture

	CreateMultiplex(ctx workflow.Context, input *medialive.CreateMultiplexInput) (*medialive.CreateMultiplexOutput, error)
	CreateMultiplexAsync(ctx workflow.Context, input *medialive.CreateMultiplexInput) *MediaLiveCreateMultiplexFuture

	CreateMultiplexProgram(ctx workflow.Context, input *medialive.CreateMultiplexProgramInput) (*medialive.CreateMultiplexProgramOutput, error)
	CreateMultiplexProgramAsync(ctx workflow.Context, input *medialive.CreateMultiplexProgramInput) *MediaLiveCreateMultiplexProgramFuture

	CreateTags(ctx workflow.Context, input *medialive.CreateTagsInput) (*medialive.CreateTagsOutput, error)
	CreateTagsAsync(ctx workflow.Context, input *medialive.CreateTagsInput) *MediaLiveCreateTagsFuture

	DeleteChannel(ctx workflow.Context, input *medialive.DeleteChannelInput) (*medialive.DeleteChannelOutput, error)
	DeleteChannelAsync(ctx workflow.Context, input *medialive.DeleteChannelInput) *MediaLiveDeleteChannelFuture

	DeleteInput(ctx workflow.Context, input *medialive.DeleteInputInput) (*medialive.DeleteInputOutput, error)
	DeleteInputAsync(ctx workflow.Context, input *medialive.DeleteInputInput) *MediaLiveDeleteInputFuture

	DeleteInputSecurityGroup(ctx workflow.Context, input *medialive.DeleteInputSecurityGroupInput) (*medialive.DeleteInputSecurityGroupOutput, error)
	DeleteInputSecurityGroupAsync(ctx workflow.Context, input *medialive.DeleteInputSecurityGroupInput) *MediaLiveDeleteInputSecurityGroupFuture

	DeleteMultiplex(ctx workflow.Context, input *medialive.DeleteMultiplexInput) (*medialive.DeleteMultiplexOutput, error)
	DeleteMultiplexAsync(ctx workflow.Context, input *medialive.DeleteMultiplexInput) *MediaLiveDeleteMultiplexFuture

	DeleteMultiplexProgram(ctx workflow.Context, input *medialive.DeleteMultiplexProgramInput) (*medialive.DeleteMultiplexProgramOutput, error)
	DeleteMultiplexProgramAsync(ctx workflow.Context, input *medialive.DeleteMultiplexProgramInput) *MediaLiveDeleteMultiplexProgramFuture

	DeleteReservation(ctx workflow.Context, input *medialive.DeleteReservationInput) (*medialive.DeleteReservationOutput, error)
	DeleteReservationAsync(ctx workflow.Context, input *medialive.DeleteReservationInput) *MediaLiveDeleteReservationFuture

	DeleteSchedule(ctx workflow.Context, input *medialive.DeleteScheduleInput) (*medialive.DeleteScheduleOutput, error)
	DeleteScheduleAsync(ctx workflow.Context, input *medialive.DeleteScheduleInput) *MediaLiveDeleteScheduleFuture

	DeleteTags(ctx workflow.Context, input *medialive.DeleteTagsInput) (*medialive.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *medialive.DeleteTagsInput) *MediaLiveDeleteTagsFuture

	DescribeChannel(ctx workflow.Context, input *medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error)
	DescribeChannelAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) *MediaLiveDescribeChannelFuture

	DescribeInput(ctx workflow.Context, input *medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error)
	DescribeInputAsync(ctx workflow.Context, input *medialive.DescribeInputInput) *MediaLiveDescribeInputFuture

	DescribeInputDevice(ctx workflow.Context, input *medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error)
	DescribeInputDeviceAsync(ctx workflow.Context, input *medialive.DescribeInputDeviceInput) *MediaLiveDescribeInputDeviceFuture

	DescribeInputDeviceThumbnail(ctx workflow.Context, input *medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error)
	DescribeInputDeviceThumbnailAsync(ctx workflow.Context, input *medialive.DescribeInputDeviceThumbnailInput) *MediaLiveDescribeInputDeviceThumbnailFuture

	DescribeInputSecurityGroup(ctx workflow.Context, input *medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error)
	DescribeInputSecurityGroupAsync(ctx workflow.Context, input *medialive.DescribeInputSecurityGroupInput) *MediaLiveDescribeInputSecurityGroupFuture

	DescribeMultiplex(ctx workflow.Context, input *medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error)
	DescribeMultiplexAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) *MediaLiveDescribeMultiplexFuture

	DescribeMultiplexProgram(ctx workflow.Context, input *medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error)
	DescribeMultiplexProgramAsync(ctx workflow.Context, input *medialive.DescribeMultiplexProgramInput) *MediaLiveDescribeMultiplexProgramFuture

	DescribeOffering(ctx workflow.Context, input *medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error)
	DescribeOfferingAsync(ctx workflow.Context, input *medialive.DescribeOfferingInput) *MediaLiveDescribeOfferingFuture

	DescribeReservation(ctx workflow.Context, input *medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error)
	DescribeReservationAsync(ctx workflow.Context, input *medialive.DescribeReservationInput) *MediaLiveDescribeReservationFuture

	DescribeSchedule(ctx workflow.Context, input *medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error)
	DescribeScheduleAsync(ctx workflow.Context, input *medialive.DescribeScheduleInput) *MediaLiveDescribeScheduleFuture

	ListChannels(ctx workflow.Context, input *medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error)
	ListChannelsAsync(ctx workflow.Context, input *medialive.ListChannelsInput) *MediaLiveListChannelsFuture

	ListInputDeviceTransfers(ctx workflow.Context, input *medialive.ListInputDeviceTransfersInput) (*medialive.ListInputDeviceTransfersOutput, error)
	ListInputDeviceTransfersAsync(ctx workflow.Context, input *medialive.ListInputDeviceTransfersInput) *MediaLiveListInputDeviceTransfersFuture

	ListInputDevices(ctx workflow.Context, input *medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error)
	ListInputDevicesAsync(ctx workflow.Context, input *medialive.ListInputDevicesInput) *MediaLiveListInputDevicesFuture

	ListInputSecurityGroups(ctx workflow.Context, input *medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error)
	ListInputSecurityGroupsAsync(ctx workflow.Context, input *medialive.ListInputSecurityGroupsInput) *MediaLiveListInputSecurityGroupsFuture

	ListInputs(ctx workflow.Context, input *medialive.ListInputsInput) (*medialive.ListInputsOutput, error)
	ListInputsAsync(ctx workflow.Context, input *medialive.ListInputsInput) *MediaLiveListInputsFuture

	ListMultiplexPrograms(ctx workflow.Context, input *medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error)
	ListMultiplexProgramsAsync(ctx workflow.Context, input *medialive.ListMultiplexProgramsInput) *MediaLiveListMultiplexProgramsFuture

	ListMultiplexes(ctx workflow.Context, input *medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error)
	ListMultiplexesAsync(ctx workflow.Context, input *medialive.ListMultiplexesInput) *MediaLiveListMultiplexesFuture

	ListOfferings(ctx workflow.Context, input *medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error)
	ListOfferingsAsync(ctx workflow.Context, input *medialive.ListOfferingsInput) *MediaLiveListOfferingsFuture

	ListReservations(ctx workflow.Context, input *medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error)
	ListReservationsAsync(ctx workflow.Context, input *medialive.ListReservationsInput) *MediaLiveListReservationsFuture

	ListTagsForResource(ctx workflow.Context, input *medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *medialive.ListTagsForResourceInput) *MediaLiveListTagsForResourceFuture

	PurchaseOffering(ctx workflow.Context, input *medialive.PurchaseOfferingInput) (*medialive.PurchaseOfferingOutput, error)
	PurchaseOfferingAsync(ctx workflow.Context, input *medialive.PurchaseOfferingInput) *MediaLivePurchaseOfferingFuture

	RejectInputDeviceTransfer(ctx workflow.Context, input *medialive.RejectInputDeviceTransferInput) (*medialive.RejectInputDeviceTransferOutput, error)
	RejectInputDeviceTransferAsync(ctx workflow.Context, input *medialive.RejectInputDeviceTransferInput) *MediaLiveRejectInputDeviceTransferFuture

	StartChannel(ctx workflow.Context, input *medialive.StartChannelInput) (*medialive.StartChannelOutput, error)
	StartChannelAsync(ctx workflow.Context, input *medialive.StartChannelInput) *MediaLiveStartChannelFuture

	StartMultiplex(ctx workflow.Context, input *medialive.StartMultiplexInput) (*medialive.StartMultiplexOutput, error)
	StartMultiplexAsync(ctx workflow.Context, input *medialive.StartMultiplexInput) *MediaLiveStartMultiplexFuture

	StopChannel(ctx workflow.Context, input *medialive.StopChannelInput) (*medialive.StopChannelOutput, error)
	StopChannelAsync(ctx workflow.Context, input *medialive.StopChannelInput) *MediaLiveStopChannelFuture

	StopMultiplex(ctx workflow.Context, input *medialive.StopMultiplexInput) (*medialive.StopMultiplexOutput, error)
	StopMultiplexAsync(ctx workflow.Context, input *medialive.StopMultiplexInput) *MediaLiveStopMultiplexFuture

	TransferInputDevice(ctx workflow.Context, input *medialive.TransferInputDeviceInput) (*medialive.TransferInputDeviceOutput, error)
	TransferInputDeviceAsync(ctx workflow.Context, input *medialive.TransferInputDeviceInput) *MediaLiveTransferInputDeviceFuture

	UpdateChannel(ctx workflow.Context, input *medialive.UpdateChannelInput) (*medialive.UpdateChannelOutput, error)
	UpdateChannelAsync(ctx workflow.Context, input *medialive.UpdateChannelInput) *MediaLiveUpdateChannelFuture

	UpdateChannelClass(ctx workflow.Context, input *medialive.UpdateChannelClassInput) (*medialive.UpdateChannelClassOutput, error)
	UpdateChannelClassAsync(ctx workflow.Context, input *medialive.UpdateChannelClassInput) *MediaLiveUpdateChannelClassFuture

	UpdateInput(ctx workflow.Context, input *medialive.UpdateInputInput) (*medialive.UpdateInputOutput, error)
	UpdateInputAsync(ctx workflow.Context, input *medialive.UpdateInputInput) *MediaLiveUpdateInputFuture

	UpdateInputDevice(ctx workflow.Context, input *medialive.UpdateInputDeviceInput) (*medialive.UpdateInputDeviceOutput, error)
	UpdateInputDeviceAsync(ctx workflow.Context, input *medialive.UpdateInputDeviceInput) *MediaLiveUpdateInputDeviceFuture

	UpdateInputSecurityGroup(ctx workflow.Context, input *medialive.UpdateInputSecurityGroupInput) (*medialive.UpdateInputSecurityGroupOutput, error)
	UpdateInputSecurityGroupAsync(ctx workflow.Context, input *medialive.UpdateInputSecurityGroupInput) *MediaLiveUpdateInputSecurityGroupFuture

	UpdateMultiplex(ctx workflow.Context, input *medialive.UpdateMultiplexInput) (*medialive.UpdateMultiplexOutput, error)
	UpdateMultiplexAsync(ctx workflow.Context, input *medialive.UpdateMultiplexInput) *MediaLiveUpdateMultiplexFuture

	UpdateMultiplexProgram(ctx workflow.Context, input *medialive.UpdateMultiplexProgramInput) (*medialive.UpdateMultiplexProgramOutput, error)
	UpdateMultiplexProgramAsync(ctx workflow.Context, input *medialive.UpdateMultiplexProgramInput) *MediaLiveUpdateMultiplexProgramFuture

	UpdateReservation(ctx workflow.Context, input *medialive.UpdateReservationInput) (*medialive.UpdateReservationOutput, error)
	UpdateReservationAsync(ctx workflow.Context, input *medialive.UpdateReservationInput) *MediaLiveUpdateReservationFuture

	WaitUntilChannelCreated(ctx workflow.Context, input *medialive.DescribeChannelInput) error
	WaitUntilChannelCreatedAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) *clients.VoidFuture

	WaitUntilChannelDeleted(ctx workflow.Context, input *medialive.DescribeChannelInput) error
	WaitUntilChannelDeletedAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) *clients.VoidFuture

	WaitUntilChannelRunning(ctx workflow.Context, input *medialive.DescribeChannelInput) error
	WaitUntilChannelRunningAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) *clients.VoidFuture

	WaitUntilChannelStopped(ctx workflow.Context, input *medialive.DescribeChannelInput) error
	WaitUntilChannelStoppedAsync(ctx workflow.Context, input *medialive.DescribeChannelInput) *clients.VoidFuture

	WaitUntilInputAttached(ctx workflow.Context, input *medialive.DescribeInputInput) error
	WaitUntilInputAttachedAsync(ctx workflow.Context, input *medialive.DescribeInputInput) *clients.VoidFuture

	WaitUntilInputDeleted(ctx workflow.Context, input *medialive.DescribeInputInput) error
	WaitUntilInputDeletedAsync(ctx workflow.Context, input *medialive.DescribeInputInput) *clients.VoidFuture

	WaitUntilInputDetached(ctx workflow.Context, input *medialive.DescribeInputInput) error
	WaitUntilInputDetachedAsync(ctx workflow.Context, input *medialive.DescribeInputInput) *clients.VoidFuture

	WaitUntilMultiplexCreated(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexCreatedAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) *clients.VoidFuture

	WaitUntilMultiplexDeleted(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexDeletedAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) *clients.VoidFuture

	WaitUntilMultiplexRunning(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexRunningAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) *clients.VoidFuture

	WaitUntilMultiplexStopped(ctx workflow.Context, input *medialive.DescribeMultiplexInput) error
	WaitUntilMultiplexStoppedAsync(ctx workflow.Context, input *medialive.DescribeMultiplexInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
