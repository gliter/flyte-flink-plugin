// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl-flink/flink.proto

package flyteidl_flink

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Resource_PersistentVolume_Type int32

const (
	Resource_PersistentVolume_PD_STANDARD Resource_PersistentVolume_Type = 0
	Resource_PersistentVolume_PD_SSD      Resource_PersistentVolume_Type = 1
)

var Resource_PersistentVolume_Type_name = map[int32]string{
	0: "PD_STANDARD",
	1: "PD_SSD",
}

var Resource_PersistentVolume_Type_value = map[string]int32{
	"PD_STANDARD": 0,
	"PD_SSD":      1,
}

func (x Resource_PersistentVolume_Type) String() string {
	return proto.EnumName(Resource_PersistentVolume_Type_name, int32(x))
}

func (Resource_PersistentVolume_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0, 1, 0}
}

type Resource struct {
	Cpu                  *Resource_Quantity         `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               *Resource_Quantity         `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	PersistentVolume     *Resource_PersistentVolume `protobuf:"bytes,3,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetCpu() *Resource_Quantity {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Resource) GetMemory() *Resource_Quantity {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Resource) GetPersistentVolume() *Resource_PersistentVolume {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

// Value must be a valid k8s quantity. See
// https://github.com/kubernetes/apimachinery/blob/master/pkg/api/resource/quantity.go#L30-L80
type Resource_Quantity struct {
	String_              string   `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource_Quantity) Reset()         { *m = Resource_Quantity{} }
func (m *Resource_Quantity) String() string { return proto.CompactTextString(m) }
func (*Resource_Quantity) ProtoMessage()    {}
func (*Resource_Quantity) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0, 0}
}

func (m *Resource_Quantity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource_Quantity.Unmarshal(m, b)
}
func (m *Resource_Quantity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource_Quantity.Marshal(b, m, deterministic)
}
func (m *Resource_Quantity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource_Quantity.Merge(m, src)
}
func (m *Resource_Quantity) XXX_Size() int {
	return xxx_messageInfo_Resource_Quantity.Size(m)
}
func (m *Resource_Quantity) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource_Quantity.DiscardUnknown(m)
}

var xxx_messageInfo_Resource_Quantity proto.InternalMessageInfo

func (m *Resource_Quantity) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

type Resource_PersistentVolume struct {
	Type                 Resource_PersistentVolume_Type `protobuf:"varint,1,opt,name=type,proto3,enum=flyteidl_flink.Resource_PersistentVolume_Type" json:"type,omitempty"`
	Size                 *Resource_Quantity             `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Resource_PersistentVolume) Reset()         { *m = Resource_PersistentVolume{} }
func (m *Resource_PersistentVolume) String() string { return proto.CompactTextString(m) }
func (*Resource_PersistentVolume) ProtoMessage()    {}
func (*Resource_PersistentVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0, 1}
}

func (m *Resource_PersistentVolume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource_PersistentVolume.Unmarshal(m, b)
}
func (m *Resource_PersistentVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource_PersistentVolume.Marshal(b, m, deterministic)
}
func (m *Resource_PersistentVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource_PersistentVolume.Merge(m, src)
}
func (m *Resource_PersistentVolume) XXX_Size() int {
	return xxx_messageInfo_Resource_PersistentVolume.Size(m)
}
func (m *Resource_PersistentVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource_PersistentVolume.DiscardUnknown(m)
}

var xxx_messageInfo_Resource_PersistentVolume proto.InternalMessageInfo

func (m *Resource_PersistentVolume) GetType() Resource_PersistentVolume_Type {
	if m != nil {
		return m.Type
	}
	return Resource_PersistentVolume_PD_STANDARD
}

func (m *Resource_PersistentVolume) GetSize() *Resource_Quantity {
	if m != nil {
		return m.Size
	}
	return nil
}

type JobManager struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *JobManager) Reset()         { *m = JobManager{} }
func (m *JobManager) String() string { return proto.CompactTextString(m) }
func (*JobManager) ProtoMessage()    {}
func (*JobManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{1}
}

func (m *JobManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobManager.Unmarshal(m, b)
}
func (m *JobManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobManager.Marshal(b, m, deterministic)
}
func (m *JobManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobManager.Merge(m, src)
}
func (m *JobManager) XXX_Size() int {
	return xxx_messageInfo_JobManager.Size(m)
}
func (m *JobManager) XXX_DiscardUnknown() {
	xxx_messageInfo_JobManager.DiscardUnknown(m)
}

var xxx_messageInfo_JobManager proto.InternalMessageInfo

func (m *JobManager) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type TaskManager struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Replicas             int32     `protobuf:"varint,2,opt,name=replicas,proto3" json:"replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TaskManager) Reset()         { *m = TaskManager{} }
func (m *TaskManager) String() string { return proto.CompactTextString(m) }
func (*TaskManager) ProtoMessage()    {}
func (*TaskManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{2}
}

func (m *TaskManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskManager.Unmarshal(m, b)
}
func (m *TaskManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskManager.Marshal(b, m, deterministic)
}
func (m *TaskManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskManager.Merge(m, src)
}
func (m *TaskManager) XXX_Size() int {
	return xxx_messageInfo_TaskManager.Size(m)
}
func (m *TaskManager) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskManager.DiscardUnknown(m)
}

var xxx_messageInfo_TaskManager proto.InternalMessageInfo

func (m *TaskManager) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TaskManager) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

type JFlyte struct {
	IndexFileLocation    string             `protobuf:"bytes,1,opt,name=index_file_location,json=indexFileLocation,proto3" json:"index_file_location,omitempty"`
	Artifacts            []*JFlyte_Artifact `protobuf:"bytes,2,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JFlyte) Reset()         { *m = JFlyte{} }
func (m *JFlyte) String() string { return proto.CompactTextString(m) }
func (*JFlyte) ProtoMessage()    {}
func (*JFlyte) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{3}
}

func (m *JFlyte) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JFlyte.Unmarshal(m, b)
}
func (m *JFlyte) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JFlyte.Marshal(b, m, deterministic)
}
func (m *JFlyte) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JFlyte.Merge(m, src)
}
func (m *JFlyte) XXX_Size() int {
	return xxx_messageInfo_JFlyte.Size(m)
}
func (m *JFlyte) XXX_DiscardUnknown() {
	xxx_messageInfo_JFlyte.DiscardUnknown(m)
}

var xxx_messageInfo_JFlyte proto.InternalMessageInfo

func (m *JFlyte) GetIndexFileLocation() string {
	if m != nil {
		return m.IndexFileLocation
	}
	return ""
}

func (m *JFlyte) GetArtifacts() []*JFlyte_Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type JFlyte_Artifact struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JFlyte_Artifact) Reset()         { *m = JFlyte_Artifact{} }
func (m *JFlyte_Artifact) String() string { return proto.CompactTextString(m) }
func (*JFlyte_Artifact) ProtoMessage()    {}
func (*JFlyte_Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{3, 0}
}

func (m *JFlyte_Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JFlyte_Artifact.Unmarshal(m, b)
}
func (m *JFlyte_Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JFlyte_Artifact.Marshal(b, m, deterministic)
}
func (m *JFlyte_Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JFlyte_Artifact.Merge(m, src)
}
func (m *JFlyte_Artifact) XXX_Size() int {
	return xxx_messageInfo_JFlyte_Artifact.Size(m)
}
func (m *JFlyte_Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_JFlyte_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_JFlyte_Artifact proto.InternalMessageInfo

func (m *JFlyte_Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JFlyte_Artifact) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

// Custom Proto for Flink Plugin.
type FlinkJob struct {
	JarFiles        []string          `protobuf:"bytes,1,rep,name=jarFiles,proto3" json:"jarFiles,omitempty"`
	MainClass       string            `protobuf:"bytes,2,opt,name=mainClass,proto3" json:"mainClass,omitempty"`
	Args            []string          `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	FlinkProperties map[string]string `protobuf:"bytes,4,rep,name=flinkProperties,proto3" json:"flinkProperties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	JobManager      *JobManager       `protobuf:"bytes,5,opt,name=jobManager,proto3" json:"jobManager,omitempty"`
	TaskManager     *TaskManager      `protobuf:"bytes,6,opt,name=taskManager,proto3" json:"taskManager,omitempty"`
	ServiceAccount  string            `protobuf:"bytes,7,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	Image           string            `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`
	FlinkVersion    string            `protobuf:"bytes,9,opt,name=flinkVersion,proto3" json:"flinkVersion,omitempty"`
	Parallelism     int32             `protobuf:"varint,10,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	// if using experiment flytekit-java this will contain all artifacts required
	// to run the task
	Jflyte               *JFlyte  `protobuf:"bytes,100,opt,name=jflyte,proto3" json:"jflyte,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlinkJob) Reset()         { *m = FlinkJob{} }
func (m *FlinkJob) String() string { return proto.CompactTextString(m) }
func (*FlinkJob) ProtoMessage()    {}
func (*FlinkJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{4}
}

func (m *FlinkJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlinkJob.Unmarshal(m, b)
}
func (m *FlinkJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlinkJob.Marshal(b, m, deterministic)
}
func (m *FlinkJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlinkJob.Merge(m, src)
}
func (m *FlinkJob) XXX_Size() int {
	return xxx_messageInfo_FlinkJob.Size(m)
}
func (m *FlinkJob) XXX_DiscardUnknown() {
	xxx_messageInfo_FlinkJob.DiscardUnknown(m)
}

var xxx_messageInfo_FlinkJob proto.InternalMessageInfo

func (m *FlinkJob) GetJarFiles() []string {
	if m != nil {
		return m.JarFiles
	}
	return nil
}

func (m *FlinkJob) GetMainClass() string {
	if m != nil {
		return m.MainClass
	}
	return ""
}

func (m *FlinkJob) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *FlinkJob) GetFlinkProperties() map[string]string {
	if m != nil {
		return m.FlinkProperties
	}
	return nil
}

func (m *FlinkJob) GetJobManager() *JobManager {
	if m != nil {
		return m.JobManager
	}
	return nil
}

func (m *FlinkJob) GetTaskManager() *TaskManager {
	if m != nil {
		return m.TaskManager
	}
	return nil
}

func (m *FlinkJob) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *FlinkJob) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *FlinkJob) GetFlinkVersion() string {
	if m != nil {
		return m.FlinkVersion
	}
	return ""
}

func (m *FlinkJob) GetParallelism() int32 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

func (m *FlinkJob) GetJflyte() *JFlyte {
	if m != nil {
		return m.Jflyte
	}
	return nil
}

type JobExecutionInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobExecutionInfo) Reset()         { *m = JobExecutionInfo{} }
func (m *JobExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*JobExecutionInfo) ProtoMessage()    {}
func (*JobExecutionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{5}
}

func (m *JobExecutionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobExecutionInfo.Unmarshal(m, b)
}
func (m *JobExecutionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobExecutionInfo.Marshal(b, m, deterministic)
}
func (m *JobExecutionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobExecutionInfo.Merge(m, src)
}
func (m *JobExecutionInfo) XXX_Size() int {
	return xxx_messageInfo_JobExecutionInfo.Size(m)
}
func (m *JobExecutionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JobExecutionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JobExecutionInfo proto.InternalMessageInfo

func (m *JobExecutionInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JobManagerExecutionInfo struct {
	IngressURLs          []string `protobuf:"bytes,1,rep,name=ingressURLs,proto3" json:"ingressURLs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobManagerExecutionInfo) Reset()         { *m = JobManagerExecutionInfo{} }
func (m *JobManagerExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*JobManagerExecutionInfo) ProtoMessage()    {}
func (*JobManagerExecutionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{6}
}

func (m *JobManagerExecutionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobManagerExecutionInfo.Unmarshal(m, b)
}
func (m *JobManagerExecutionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobManagerExecutionInfo.Marshal(b, m, deterministic)
}
func (m *JobManagerExecutionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobManagerExecutionInfo.Merge(m, src)
}
func (m *JobManagerExecutionInfo) XXX_Size() int {
	return xxx_messageInfo_JobManagerExecutionInfo.Size(m)
}
func (m *JobManagerExecutionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JobManagerExecutionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JobManagerExecutionInfo proto.InternalMessageInfo

func (m *JobManagerExecutionInfo) GetIngressURLs() []string {
	if m != nil {
		return m.IngressURLs
	}
	return nil
}

type FlinkExecutionInfo struct {
	Job                  *JobExecutionInfo        `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	JobManager           *JobManagerExecutionInfo `protobuf:"bytes,2,opt,name=jobManager,proto3" json:"jobManager,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FlinkExecutionInfo) Reset()         { *m = FlinkExecutionInfo{} }
func (m *FlinkExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*FlinkExecutionInfo) ProtoMessage()    {}
func (*FlinkExecutionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{7}
}

func (m *FlinkExecutionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlinkExecutionInfo.Unmarshal(m, b)
}
func (m *FlinkExecutionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlinkExecutionInfo.Marshal(b, m, deterministic)
}
func (m *FlinkExecutionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlinkExecutionInfo.Merge(m, src)
}
func (m *FlinkExecutionInfo) XXX_Size() int {
	return xxx_messageInfo_FlinkExecutionInfo.Size(m)
}
func (m *FlinkExecutionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FlinkExecutionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FlinkExecutionInfo proto.InternalMessageInfo

func (m *FlinkExecutionInfo) GetJob() *JobExecutionInfo {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *FlinkExecutionInfo) GetJobManager() *JobManagerExecutionInfo {
	if m != nil {
		return m.JobManager
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl_flink.Resource_PersistentVolume_Type", Resource_PersistentVolume_Type_name, Resource_PersistentVolume_Type_value)
	proto.RegisterType((*Resource)(nil), "flyteidl_flink.Resource")
	proto.RegisterType((*Resource_Quantity)(nil), "flyteidl_flink.Resource.Quantity")
	proto.RegisterType((*Resource_PersistentVolume)(nil), "flyteidl_flink.Resource.PersistentVolume")
	proto.RegisterType((*JobManager)(nil), "flyteidl_flink.JobManager")
	proto.RegisterType((*TaskManager)(nil), "flyteidl_flink.TaskManager")
	proto.RegisterType((*JFlyte)(nil), "flyteidl_flink.JFlyte")
	proto.RegisterType((*JFlyte_Artifact)(nil), "flyteidl_flink.JFlyte.Artifact")
	proto.RegisterType((*FlinkJob)(nil), "flyteidl_flink.FlinkJob")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl_flink.FlinkJob.FlinkPropertiesEntry")
	proto.RegisterType((*JobExecutionInfo)(nil), "flyteidl_flink.JobExecutionInfo")
	proto.RegisterType((*JobManagerExecutionInfo)(nil), "flyteidl_flink.JobManagerExecutionInfo")
	proto.RegisterType((*FlinkExecutionInfo)(nil), "flyteidl_flink.FlinkExecutionInfo")
}

func init() { proto.RegisterFile("flyteidl-flink/flink.proto", fileDescriptor_13f39d29671649d9) }

var fileDescriptor_13f39d29671649d9 = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xee, 0xf8, 0x2f, 0xeb, 0xe3, 0xca, 0x5d, 0x86, 0x8a, 0xac, 0xcc, 0x05, 0x61, 0x0b, 0xc5,
	0x69, 0xf0, 0xb6, 0x72, 0x40, 0xa2, 0x45, 0x15, 0xca, 0x92, 0xa4, 0x22, 0x4d, 0x51, 0x98, 0xa6,
	0xad, 0x44, 0x14, 0xa2, 0xf1, 0x7a, 0x6c, 0x26, 0xd9, 0x3f, 0xcd, 0xcc, 0x46, 0x35, 0x97, 0x3c,
	0x01, 0xf4, 0x79, 0x90, 0x78, 0x04, 0x9e, 0x83, 0x57, 0xf0, 0x15, 0x9a, 0xd9, 0xf5, 0xcf, 0xba,
	0x89, 0x14, 0x89, 0x1b, 0x6b, 0x7c, 0xce, 0xf7, 0x7d, 0xe7, 0xcc, 0x39, 0xdf, 0xee, 0x42, 0x67,
	0x14, 0x4e, 0x14, 0xe3, 0xc3, 0xb0, 0x37, 0x0a, 0x79, 0x7c, 0xf1, 0xd0, 0xfc, 0x7a, 0xa9, 0x48,
	0x54, 0x82, 0xdb, 0xb3, 0xdc, 0x99, 0x89, 0x76, 0xd6, 0x2f, 0x69, 0xc8, 0x87, 0x54, 0xb1, 0x87,
	0xb3, 0x43, 0x0e, 0x74, 0xff, 0xad, 0x82, 0x45, 0x98, 0x4c, 0x32, 0x11, 0x30, 0xbc, 0x0d, 0xd5,
	0x20, 0xcd, 0x1c, 0xb4, 0x81, 0xba, 0xad, 0xfe, 0xa7, 0x5e, 0x59, 0xc3, 0x9b, 0xc1, 0xbc, 0x9f,
	0x32, 0x1a, 0x2b, 0xae, 0x26, 0x44, 0xa3, 0xf1, 0x63, 0x68, 0x44, 0x2c, 0x4a, 0xc4, 0xc4, 0xa9,
	0xdc, 0x94, 0x57, 0x10, 0xf0, 0x2b, 0xb0, 0x53, 0x26, 0x24, 0x97, 0x8a, 0xc5, 0xea, 0x75, 0x12,
	0x66, 0x11, 0x73, 0xaa, 0x46, 0x64, 0xf3, 0x5a, 0x91, 0xa3, 0x15, 0x02, 0x79, 0x4f, 0xa2, 0xf3,
	0x06, 0xac, 0x59, 0x29, 0xfc, 0x1c, 0x1a, 0x52, 0x09, 0x1e, 0x8f, 0xcd, 0xad, 0x9a, 0xfe, 0xf6,
	0xd4, 0x7f, 0x24, 0xbc, 0xfe, 0x97, 0xbf, 0x74, 0x4f, 0xb6, 0x7a, 0xa7, 0xdf, 0x9d, 0x3c, 0xea,
	0x3d, 0xf6, 0x4e, 0xb7, 0x36, 0xbb, 0x27, 0x6c, 0x8f, 0xc7, 0x59, 0x74, 0xf1, 0xfc, 0xc5, 0xb3,
	0xe3, 0xa3, 0xd3, 0x07, 0x27, 0xbd, 0xad, 0x3c, 0x79, 0xfa, 0x60, 0xf3, 0x33, 0x52, 0x48, 0x74,
	0xfe, 0x42, 0x60, 0xaf, 0xd6, 0xc7, 0x87, 0x50, 0x53, 0x93, 0x94, 0x19, 0xfd, 0x76, 0xdf, 0xbb,
	0x71, 0xe3, 0xde, 0xf1, 0x24, 0x65, 0xbe, 0x35, 0xf5, 0xeb, 0xbf, 0xa3, 0x8a, 0x8d, 0x88, 0x51,
	0xc1, 0x5f, 0x43, 0x4d, 0xf2, 0xdf, 0xd8, 0xcd, 0x67, 0x69, 0xe0, 0xee, 0x3d, 0xa8, 0x69, 0x39,
	0x7c, 0x07, 0x5a, 0x47, 0xbb, 0x67, 0x2f, 0x8f, 0x77, 0x7e, 0xdc, 0xdd, 0x21, 0xbb, 0xf6, 0x2d,
	0x0c, 0xd0, 0xd0, 0x81, 0x97, 0xbb, 0x36, 0x72, 0x7d, 0x80, 0x83, 0x64, 0xf0, 0x82, 0xc6, 0x74,
	0xcc, 0x04, 0xfe, 0x0a, 0x2c, 0x51, 0xa8, 0x15, 0x1b, 0x77, 0xae, 0xab, 0x46, 0xe6, 0x48, 0xf7,
	0x57, 0x68, 0x1d, 0x53, 0x79, 0xf1, 0xbf, 0x44, 0xf0, 0x3d, 0xcd, 0x4a, 0x43, 0x1e, 0x50, 0x69,
	0x2e, 0x5a, 0xf7, 0xd7, 0xa6, 0x7e, 0xad, 0x53, 0xe9, 0xde, 0x22, 0xf3, 0x84, 0xfb, 0x37, 0x82,
	0xc6, 0xc1, 0xbe, 0xd6, 0xc2, 0x1e, 0x7c, 0xc8, 0xe3, 0x21, 0x7b, 0x7b, 0x36, 0xe2, 0x21, 0x3b,
	0x0b, 0x93, 0x80, 0x2a, 0x9e, 0xc4, 0xf9, 0x46, 0xc9, 0x07, 0x26, 0xb5, 0xcf, 0x43, 0x76, 0x58,
	0x24, 0xf0, 0x53, 0x68, 0x52, 0xa1, 0xf8, 0x88, 0x06, 0x4a, 0x17, 0xa8, 0x76, 0x5b, 0xfd, 0x4f,
	0x56, 0xdb, 0xca, 0xa5, 0xbd, 0x9d, 0x02, 0x47, 0x16, 0x8c, 0xce, 0x3e, 0x58, 0xb3, 0x30, 0xc6,
	0x50, 0x8b, 0x69, 0xc4, 0x8a, 0x5a, 0xe6, 0x8c, 0xef, 0x83, 0x35, 0xef, 0xa1, 0x62, 0x5c, 0x05,
	0x53, 0x7f, 0x4d, 0xd4, 0x6d, 0xf4, 0x07, 0x42, 0x64, 0x9e, 0x73, 0xff, 0xa9, 0x81, 0xb5, 0xaf,
	0x8b, 0x1d, 0x24, 0x03, 0xbc, 0x0d, 0xd6, 0x39, 0x15, 0xba, 0x4d, 0xe9, 0xa0, 0x8d, 0x6a, 0xb7,
	0xe9, 0xaf, 0x4f, 0xfd, 0xfa, 0x3b, 0x54, 0x71, 0xd0, 0xd4, 0xbf, 0xfd, 0x0e, 0x35, 0xdd, 0x85,
	0xc2, 0x0c, 0x88, 0x3f, 0x87, 0x66, 0x44, 0x79, 0xfc, 0x7d, 0x48, 0xa5, 0x2c, 0x4a, 0xe9, 0x49,
	0x09, 0xed, 0x97, 0x45, 0x46, 0x37, 0x49, 0xc5, 0x58, 0x3a, 0x55, 0xad, 0x4b, 0xcc, 0x19, 0xbf,
	0x81, 0x3b, 0xe6, 0xa2, 0x47, 0x22, 0x49, 0x99, 0x50, 0x9c, 0x49, 0xa7, 0x66, 0x26, 0xd1, 0x5b,
	0x9d, 0xc4, 0xac, 0xc5, 0xfc, 0xb0, 0xc0, 0xef, 0xc5, 0x4a, 0x4c, 0xc8, 0xaa, 0x0a, 0x7e, 0x02,
	0x70, 0x3e, 0x77, 0x91, 0x53, 0x37, 0x4b, 0xef, 0xbc, 0x37, 0xdd, 0x39, 0x82, 0x2c, 0xa1, 0xf1,
	0x53, 0x68, 0xa9, 0x85, 0x7b, 0x9c, 0x86, 0x21, 0x7f, 0xbc, 0x4a, 0x5e, 0x32, 0x18, 0x59, 0xc6,
	0xe3, 0xfb, 0xd0, 0x96, 0x4c, 0x5c, 0xf2, 0x80, 0xed, 0x04, 0x41, 0x92, 0xc5, 0xca, 0x59, 0x33,
	0x6b, 0x59, 0x89, 0xe2, 0xbb, 0x50, 0xe7, 0x11, 0x1d, 0x33, 0xc7, 0x32, 0xe9, 0xfc, 0x0f, 0x76,
	0xe1, 0xb6, 0xd1, 0x7f, 0xad, 0x1f, 0xc4, 0x24, 0x76, 0x9a, 0x26, 0x59, 0x8a, 0xe1, 0x4d, 0x68,
	0xa5, 0x54, 0xd0, 0x30, 0x64, 0x21, 0x97, 0x91, 0x03, 0x65, 0x73, 0x2e, 0xe7, 0xb0, 0x07, 0x8d,
	0x73, 0xd3, 0xb8, 0x33, 0x34, 0xd7, 0xf8, 0xe8, 0x6a, 0x87, 0x91, 0x02, 0xd5, 0xf1, 0xe1, 0xee,
	0x55, 0x03, 0xc6, 0x36, 0x54, 0x2f, 0xd8, 0xa4, 0x30, 0x98, 0x3e, 0xea, 0xf6, 0x2f, 0x69, 0x98,
	0xe5, 0x2f, 0x81, 0x26, 0xc9, 0xff, 0x3c, 0xa9, 0x7c, 0x83, 0x5c, 0x17, 0xec, 0x83, 0x64, 0xb0,
	0xf7, 0x96, 0x05, 0x99, 0x76, 0xd8, 0x0f, 0xf1, 0x28, 0xc1, 0x6d, 0xa8, 0xf0, 0x61, 0x41, 0xaf,
	0xf0, 0xa1, 0xfb, 0x2d, 0xac, 0x2f, 0xa6, 0x5f, 0x86, 0x6e, 0x40, 0x8b, 0xc7, 0x63, 0xc1, 0xa4,
	0x7c, 0x45, 0x0e, 0x0b, 0x1b, 0x92, 0xe5, 0x90, 0xfb, 0x27, 0x02, 0x6c, 0xba, 0x2c, 0x13, 0xfb,
	0x50, 0x3d, 0x4f, 0x06, 0xc5, 0x13, 0xbe, 0x71, 0xc5, 0xb2, 0x4b, 0x70, 0xa2, 0xc1, 0xf8, 0x59,
	0xc9, 0x27, 0xf9, 0xfb, 0xec, 0x8b, 0xeb, 0x7d, 0x52, 0x56, 0x58, 0xa2, 0xfa, 0xf6, 0xcf, 0xed,
	0xf2, 0x97, 0x6e, 0xd0, 0x30, 0xdf, 0xae, 0xed, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x44,
	0xad, 0x44, 0x02, 0x07, 0x00, 0x00,
}
