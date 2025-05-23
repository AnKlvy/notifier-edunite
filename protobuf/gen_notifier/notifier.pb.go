// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: notifier.proto

package gen_notifier

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Сообщение для отправки уведомлений
type Notification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Subject       string                 `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Images        []string               `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`
	Metadata      map[string]string      `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // ISO 8601 формат даты/времени
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // ISO 8601 формат даты/времени
	Id            int64                  `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`                               // ID уведомления
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_notifier_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Notification) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Notification) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Notification) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Notification) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Notification) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UsersNotification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UsersIds      []string               `protobuf:"bytes,1,rep,name=users_ids,json=usersIds,proto3" json:"users_ids,omitempty"` // ID пользователя
	Notification  *Notification          `protobuf:"bytes,2,opt,name=notification,proto3" json:"notification,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersNotification) Reset() {
	*x = UsersNotification{}
	mi := &file_notifier_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersNotification) ProtoMessage() {}

func (x *UsersNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersNotification.ProtoReflect.Descriptor instead.
func (*UsersNotification) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{1}
}

func (x *UsersNotification) GetUsersIds() []string {
	if x != nil {
		return x.UsersIds
	}
	return nil
}

func (x *UsersNotification) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

// Запрос на подписку пользователя на канал уведомлений
type SubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                                                                 // ID пользователя
	Channel       string                 `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`                                                                             // Канал (email, fcm, telegram)
	Value         string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`                                                                                 // Токен или email
	Metadata      map[string]string      `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // Метаданные (например, устройство или статус)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_notifier_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SubscribeRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *SubscribeRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SubscribeRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Запрос на отписку пользователя от канала
type UnsubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ID пользователя
	Channel       string                 `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`             // Канал
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	mi := &file_notifier_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{3}
}

func (x *UnsubscribeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UnsubscribeRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

// Ответ на запрос отправки уведомления
type SuccessResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                              // Успешно ли отправлено
	ErrorMessage  string                 `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // Сообщение об ошибке (если есть)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	mi := &file_notifier_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{4}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SuccessResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Настройки уведомлений пользователя
type NotifierSettings struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ID пользователя
	Channel       string                 `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`             // Канал
	Token         string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`                 // Токен или email
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotifierSettings) Reset() {
	*x = NotifierSettings{}
	mi := &file_notifier_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifierSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifierSettings) ProtoMessage() {}

func (x *NotifierSettings) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifierSettings.ProtoReflect.Descriptor instead.
func (*NotifierSettings) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{5}
}

func (x *NotifierSettings) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NotifierSettings) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *NotifierSettings) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Ответ со всеми настройками
type GetAllSettingsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Settings      []*NotifierSettings    `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllSettingsResponse) Reset() {
	*x = GetAllSettingsResponse{}
	mi := &file_notifier_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSettingsResponse) ProtoMessage() {}

func (x *GetAllSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetAllSettingsResponse) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllSettingsResponse) GetSettings() []*NotifierSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

// Ответ со всеми уведомлениями
type GetAllNotificationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Notifications []*Notification        `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllNotificationsResponse) Reset() {
	*x = GetAllNotificationsResponse{}
	mi := &file_notifier_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNotificationsResponse) ProtoMessage() {}

func (x *GetAllNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetAllNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

// Запрос на получение настроек пользователя
type GetUserSettingsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // ID пользователя
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserSettingsRequest) Reset() {
	*x = GetUserSettingsRequest{}
	mi := &file_notifier_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingsRequest) ProtoMessage() {}

func (x *GetUserSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetUserSettingsRequest) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserSettingsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Запрос на получение уведомлений пользователя
type GetUserNotificationsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // ID пользователя
	FromDate      string                 `protobuf:"bytes,2,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"` // Дата начала в формате ISO 8601
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserNotificationsRequest) Reset() {
	*x = GetUserNotificationsRequest{}
	mi := &file_notifier_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserNotificationsRequest) ProtoMessage() {}

func (x *GetUserNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifier_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetUserNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_notifier_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserNotificationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserNotificationsRequest) GetFromDate() string {
	if x != nil {
		return x.FromDate
	}
	return ""
}

var File_notifier_proto protoreflect.FileDescriptor

const file_notifier_proto_rawDesc = "" +
	"\n" +
	"\x0enotifier.proto\x12\x06notify\x1a\x1bgoogle/protobuf/empty.proto\"\xa5\x02\n" +
	"\fNotification\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\x18\n" +
	"\asubject\x18\x02 \x01(\tR\asubject\x12\x16\n" +
	"\x06images\x18\x03 \x03(\tR\x06images\x12>\n" +
	"\bmetadata\x18\x04 \x03(\v2\".notify.Notification.MetadataEntryR\bmetadata\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\tR\tupdatedAt\x12\x0e\n" +
	"\x02id\x18\a \x01(\x03R\x02id\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"j\n" +
	"\x11UsersNotification\x12\x1b\n" +
	"\tusers_ids\x18\x01 \x03(\tR\busersIds\x128\n" +
	"\fnotification\x18\x02 \x01(\v2\x14.notify.NotificationR\fnotification\"\xdc\x01\n" +
	"\x10SubscribeRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x18\n" +
	"\achannel\x18\x02 \x01(\tR\achannel\x12\x14\n" +
	"\x05value\x18\x03 \x01(\tR\x05value\x12B\n" +
	"\bmetadata\x18\x04 \x03(\v2&.notify.SubscribeRequest.MetadataEntryR\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"G\n" +
	"\x12UnsubscribeRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x18\n" +
	"\achannel\x18\x02 \x01(\tR\achannel\"P\n" +
	"\x0fSuccessResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12#\n" +
	"\rerror_message\x18\x02 \x01(\tR\ferrorMessage\"[\n" +
	"\x10NotifierSettings\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x18\n" +
	"\achannel\x18\x02 \x01(\tR\achannel\x12\x14\n" +
	"\x05token\x18\x03 \x01(\tR\x05token\"N\n" +
	"\x16GetAllSettingsResponse\x124\n" +
	"\bsettings\x18\x01 \x03(\v2\x18.notify.NotifierSettingsR\bsettings\"Y\n" +
	"\x1bGetAllNotificationsResponse\x12:\n" +
	"\rnotifications\x18\x01 \x03(\v2\x14.notify.NotificationR\rnotifications\"1\n" +
	"\x16GetUserSettingsRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"S\n" +
	"\x1bGetUserNotificationsRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\tfrom_date\x18\x02 \x01(\tR\bfromDate2\xee\x04\n" +
	"\x13NotificationService\x12>\n" +
	"\tSubscribe\x12\x18.notify.SubscribeRequest\x1a\x17.notify.SuccessResponse\x12B\n" +
	"\vUnsubscribe\x12\x1a.notify.UnsubscribeRequest\x1a\x17.notify.SuccessResponse\x12G\n" +
	"\x0fSendToOneOrMany\x12\x19.notify.UsersNotification\x1a\x19.notify.UsersNotification\x127\n" +
	"\tSendToAll\x12\x14.notify.Notification\x1a\x14.notify.Notification\x12H\n" +
	"\x0eGetAllSettings\x12\x16.google.protobuf.Empty\x1a\x1e.notify.GetAllSettingsResponse\x12R\n" +
	"\x13GetAllNotifications\x12\x16.google.protobuf.Empty\x1a#.notify.GetAllNotificationsResponse\x12Q\n" +
	"\x0fGetUserSettings\x12\x1e.notify.GetUserSettingsRequest\x1a\x1e.notify.GetAllSettingsResponse\x12`\n" +
	"\x14GetUserNotifications\x12#.notify.GetUserNotificationsRequest\x1a#.notify.GetAllNotificationsResponseB8Z6github.com/AnKlvy/notify-service/protobuf/gen_notifierb\x06proto3"

var (
	file_notifier_proto_rawDescOnce sync.Once
	file_notifier_proto_rawDescData []byte
)

func file_notifier_proto_rawDescGZIP() []byte {
	file_notifier_proto_rawDescOnce.Do(func() {
		file_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notifier_proto_rawDesc), len(file_notifier_proto_rawDesc)))
	})
	return file_notifier_proto_rawDescData
}

var file_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_notifier_proto_goTypes = []any{
	(*Notification)(nil),                // 0: notify.Notification
	(*UsersNotification)(nil),           // 1: notify.UsersNotification
	(*SubscribeRequest)(nil),            // 2: notify.SubscribeRequest
	(*UnsubscribeRequest)(nil),          // 3: notify.UnsubscribeRequest
	(*SuccessResponse)(nil),             // 4: notify.SuccessResponse
	(*NotifierSettings)(nil),            // 5: notify.NotifierSettings
	(*GetAllSettingsResponse)(nil),      // 6: notify.GetAllSettingsResponse
	(*GetAllNotificationsResponse)(nil), // 7: notify.GetAllNotificationsResponse
	(*GetUserSettingsRequest)(nil),      // 8: notify.GetUserSettingsRequest
	(*GetUserNotificationsRequest)(nil), // 9: notify.GetUserNotificationsRequest
	nil,                                 // 10: notify.Notification.MetadataEntry
	nil,                                 // 11: notify.SubscribeRequest.MetadataEntry
	(*emptypb.Empty)(nil),               // 12: google.protobuf.Empty
}
var file_notifier_proto_depIdxs = []int32{
	10, // 0: notify.Notification.metadata:type_name -> notify.Notification.MetadataEntry
	0,  // 1: notify.UsersNotification.notification:type_name -> notify.Notification
	11, // 2: notify.SubscribeRequest.metadata:type_name -> notify.SubscribeRequest.MetadataEntry
	5,  // 3: notify.GetAllSettingsResponse.settings:type_name -> notify.NotifierSettings
	0,  // 4: notify.GetAllNotificationsResponse.notifications:type_name -> notify.Notification
	2,  // 5: notify.NotificationService.Subscribe:input_type -> notify.SubscribeRequest
	3,  // 6: notify.NotificationService.Unsubscribe:input_type -> notify.UnsubscribeRequest
	1,  // 7: notify.NotificationService.SendToOneOrMany:input_type -> notify.UsersNotification
	0,  // 8: notify.NotificationService.SendToAll:input_type -> notify.Notification
	12, // 9: notify.NotificationService.GetAllSettings:input_type -> google.protobuf.Empty
	12, // 10: notify.NotificationService.GetAllNotifications:input_type -> google.protobuf.Empty
	8,  // 11: notify.NotificationService.GetUserSettings:input_type -> notify.GetUserSettingsRequest
	9,  // 12: notify.NotificationService.GetUserNotifications:input_type -> notify.GetUserNotificationsRequest
	4,  // 13: notify.NotificationService.Subscribe:output_type -> notify.SuccessResponse
	4,  // 14: notify.NotificationService.Unsubscribe:output_type -> notify.SuccessResponse
	1,  // 15: notify.NotificationService.SendToOneOrMany:output_type -> notify.UsersNotification
	0,  // 16: notify.NotificationService.SendToAll:output_type -> notify.Notification
	6,  // 17: notify.NotificationService.GetAllSettings:output_type -> notify.GetAllSettingsResponse
	7,  // 18: notify.NotificationService.GetAllNotifications:output_type -> notify.GetAllNotificationsResponse
	6,  // 19: notify.NotificationService.GetUserSettings:output_type -> notify.GetAllSettingsResponse
	7,  // 20: notify.NotificationService.GetUserNotifications:output_type -> notify.GetAllNotificationsResponse
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_notifier_proto_init() }
func file_notifier_proto_init() {
	if File_notifier_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notifier_proto_rawDesc), len(file_notifier_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifier_proto_goTypes,
		DependencyIndexes: file_notifier_proto_depIdxs,
		MessageInfos:      file_notifier_proto_msgTypes,
	}.Build()
	File_notifier_proto = out.File
	file_notifier_proto_goTypes = nil
	file_notifier_proto_depIdxs = nil
}
