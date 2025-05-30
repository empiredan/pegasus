// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package admin

import (
	"bytes"
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/apache/incubator-pegasus/go-client/idl/base"
	"github.com/apache/incubator-pegasus/go-client/idl/replication"
	"github.com/apache/thrift/lib/go/thrift"
	"reflect"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = base.GoUnusedProtection__
var _ = replication.GoUnusedProtection__

type PartitionStatus int64

const (
	PartitionStatus_PS_INVALID             PartitionStatus = 0
	PartitionStatus_PS_INACTIVE            PartitionStatus = 1
	PartitionStatus_PS_ERROR               PartitionStatus = 2
	PartitionStatus_PS_PRIMARY             PartitionStatus = 3
	PartitionStatus_PS_SECONDARY           PartitionStatus = 4
	PartitionStatus_PS_POTENTIAL_SECONDARY PartitionStatus = 5
	PartitionStatus_PS_PARTITION_SPLIT     PartitionStatus = 6
)

func (p PartitionStatus) String() string {
	switch p {
	case PartitionStatus_PS_INVALID:
		return "PS_INVALID"
	case PartitionStatus_PS_INACTIVE:
		return "PS_INACTIVE"
	case PartitionStatus_PS_ERROR:
		return "PS_ERROR"
	case PartitionStatus_PS_PRIMARY:
		return "PS_PRIMARY"
	case PartitionStatus_PS_SECONDARY:
		return "PS_SECONDARY"
	case PartitionStatus_PS_POTENTIAL_SECONDARY:
		return "PS_POTENTIAL_SECONDARY"
	case PartitionStatus_PS_PARTITION_SPLIT:
		return "PS_PARTITION_SPLIT"
	}
	return "<UNSET>"
}

func PartitionStatusFromString(s string) (PartitionStatus, error) {
	switch s {
	case "PS_INVALID":
		return PartitionStatus_PS_INVALID, nil
	case "PS_INACTIVE":
		return PartitionStatus_PS_INACTIVE, nil
	case "PS_ERROR":
		return PartitionStatus_PS_ERROR, nil
	case "PS_PRIMARY":
		return PartitionStatus_PS_PRIMARY, nil
	case "PS_SECONDARY":
		return PartitionStatus_PS_SECONDARY, nil
	case "PS_POTENTIAL_SECONDARY":
		return PartitionStatus_PS_POTENTIAL_SECONDARY, nil
	case "PS_PARTITION_SPLIT":
		return PartitionStatus_PS_PARTITION_SPLIT, nil
	}
	return PartitionStatus(0), fmt.Errorf("not a valid PartitionStatus string")
}

func PartitionStatusPtr(v PartitionStatus) *PartitionStatus { return &v }

func (p PartitionStatus) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *PartitionStatus) UnmarshalText(text []byte) error {
	q, err := PartitionStatusFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *PartitionStatus) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = PartitionStatus(v)
	return nil
}

func (p *PartitionStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type SplitStatus int64

const (
	SplitStatus_NOT_SPLIT SplitStatus = 0
	SplitStatus_SPLITTING SplitStatus = 1
	SplitStatus_PAUSING   SplitStatus = 2
	SplitStatus_PAUSED    SplitStatus = 3
	SplitStatus_CANCELING SplitStatus = 4
)

func (p SplitStatus) String() string {
	switch p {
	case SplitStatus_NOT_SPLIT:
		return "NOT_SPLIT"
	case SplitStatus_SPLITTING:
		return "SPLITTING"
	case SplitStatus_PAUSING:
		return "PAUSING"
	case SplitStatus_PAUSED:
		return "PAUSED"
	case SplitStatus_CANCELING:
		return "CANCELING"
	}
	return "<UNSET>"
}

func SplitStatusFromString(s string) (SplitStatus, error) {
	switch s {
	case "NOT_SPLIT":
		return SplitStatus_NOT_SPLIT, nil
	case "SPLITTING":
		return SplitStatus_SPLITTING, nil
	case "PAUSING":
		return SplitStatus_PAUSING, nil
	case "PAUSED":
		return SplitStatus_PAUSED, nil
	case "CANCELING":
		return SplitStatus_CANCELING, nil
	}
	return SplitStatus(0), fmt.Errorf("not a valid SplitStatus string")
}

func SplitStatusPtr(v SplitStatus) *SplitStatus { return &v }

func (p SplitStatus) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *SplitStatus) UnmarshalText(text []byte) error {
	q, err := SplitStatusFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *SplitStatus) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = SplitStatus(v)
	return nil
}

func (p *SplitStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type DiskStatus int64

const (
	DiskStatus_NORMAL             DiskStatus = 0
	DiskStatus_SPACE_INSUFFICIENT DiskStatus = 1
	DiskStatus_IO_ERROR           DiskStatus = 2
)

func (p DiskStatus) String() string {
	switch p {
	case DiskStatus_NORMAL:
		return "NORMAL"
	case DiskStatus_SPACE_INSUFFICIENT:
		return "SPACE_INSUFFICIENT"
	case DiskStatus_IO_ERROR:
		return "IO_ERROR"
	}
	return "<UNSET>"
}

func DiskStatusFromString(s string) (DiskStatus, error) {
	switch s {
	case "NORMAL":
		return DiskStatus_NORMAL, nil
	case "SPACE_INSUFFICIENT":
		return DiskStatus_SPACE_INSUFFICIENT, nil
	case "IO_ERROR":
		return DiskStatus_IO_ERROR, nil
	}
	return DiskStatus(0), fmt.Errorf("not a valid DiskStatus string")
}

func DiskStatusPtr(v DiskStatus) *DiskStatus { return &v }

func (p DiskStatus) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *DiskStatus) UnmarshalText(text []byte) error {
	q, err := DiskStatusFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *DiskStatus) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = DiskStatus(v)
	return nil
}

func (p *DiskStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ManualCompactionStatus int64

const (
	ManualCompactionStatus_IDLE     ManualCompactionStatus = 0
	ManualCompactionStatus_QUEUING  ManualCompactionStatus = 1
	ManualCompactionStatus_RUNNING  ManualCompactionStatus = 2
	ManualCompactionStatus_FINISHED ManualCompactionStatus = 3
)

func (p ManualCompactionStatus) String() string {
	switch p {
	case ManualCompactionStatus_IDLE:
		return "IDLE"
	case ManualCompactionStatus_QUEUING:
		return "QUEUING"
	case ManualCompactionStatus_RUNNING:
		return "RUNNING"
	case ManualCompactionStatus_FINISHED:
		return "FINISHED"
	}
	return "<UNSET>"
}

func ManualCompactionStatusFromString(s string) (ManualCompactionStatus, error) {
	switch s {
	case "IDLE":
		return ManualCompactionStatus_IDLE, nil
	case "QUEUING":
		return ManualCompactionStatus_QUEUING, nil
	case "RUNNING":
		return ManualCompactionStatus_RUNNING, nil
	case "FINISHED":
		return ManualCompactionStatus_FINISHED, nil
	}
	return ManualCompactionStatus(0), fmt.Errorf("not a valid ManualCompactionStatus string")
}

func ManualCompactionStatusPtr(v ManualCompactionStatus) *ManualCompactionStatus { return &v }

func (p ManualCompactionStatus) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *ManualCompactionStatus) UnmarshalText(text []byte) error {
	q, err := ManualCompactionStatusFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *ManualCompactionStatus) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = ManualCompactionStatus(v)
	return nil
}

func (p *ManualCompactionStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Attributes:
//  - Name
//  - Size
//  - Md5
type FileMeta struct {
	Name string `thrift:"name,1" db:"name" json:"name"`
	Size int64  `thrift:"size,2" db:"size" json:"size"`
	Md5  string `thrift:"md5,3" db:"md5" json:"md5"`
}

func NewFileMeta() *FileMeta {
	return &FileMeta{}
}

func (p *FileMeta) GetName() string {
	return p.Name
}

func (p *FileMeta) GetSize() int64 {
	return p.Size
}

func (p *FileMeta) GetMd5() string {
	return p.Md5
}
func (p *FileMeta) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FileMeta) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *FileMeta) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Size = v
	}
	return nil
}

func (p *FileMeta) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Md5 = v
	}
	return nil
}

func (p *FileMeta) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("file_meta"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FileMeta) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
	}
	return err
}

func (p *FileMeta) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("size", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:size: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Size)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.size (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:size: ", p), err)
	}
	return err
}

func (p *FileMeta) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("md5", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:md5: ", p), err)
	}
	if err := oprot.WriteString(string(p.Md5)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.md5 (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:md5: ", p), err)
	}
	return err
}

func (p *FileMeta) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FileMeta(%+v)", *p)
}

// Attributes:
//  - Pid
//  - Ballot
//  - Primary
//  - Status
//  - LearnerSignature
//  - PopAll
//  - SplitSyncToChild
//  - HpPrimary
type ReplicaConfiguration struct {
	Pid              *base.Gpid       `thrift:"pid,1" db:"pid" json:"pid"`
	Ballot           int64            `thrift:"ballot,2" db:"ballot" json:"ballot"`
	Primary          *base.RPCAddress `thrift:"primary,3" db:"primary" json:"primary"`
	Status           PartitionStatus  `thrift:"status,4" db:"status" json:"status"`
	LearnerSignature int64            `thrift:"learner_signature,5" db:"learner_signature" json:"learner_signature"`
	PopAll           bool             `thrift:"pop_all,6" db:"pop_all" json:"pop_all"`
	SplitSyncToChild bool             `thrift:"split_sync_to_child,7" db:"split_sync_to_child" json:"split_sync_to_child"`
	HpPrimary        *base.HostPort   `thrift:"hp_primary,8" db:"hp_primary" json:"hp_primary,omitempty"`
}

func NewReplicaConfiguration() *ReplicaConfiguration {
	return &ReplicaConfiguration{
		Status: 0,
	}
}

var ReplicaConfiguration_Pid_DEFAULT *base.Gpid

func (p *ReplicaConfiguration) GetPid() *base.Gpid {
	if !p.IsSetPid() {
		return ReplicaConfiguration_Pid_DEFAULT
	}
	return p.Pid
}

func (p *ReplicaConfiguration) GetBallot() int64 {
	return p.Ballot
}

var ReplicaConfiguration_Primary_DEFAULT *base.RPCAddress

func (p *ReplicaConfiguration) GetPrimary() *base.RPCAddress {
	if !p.IsSetPrimary() {
		return ReplicaConfiguration_Primary_DEFAULT
	}
	return p.Primary
}

func (p *ReplicaConfiguration) GetStatus() PartitionStatus {
	return p.Status
}

func (p *ReplicaConfiguration) GetLearnerSignature() int64 {
	return p.LearnerSignature
}

var ReplicaConfiguration_PopAll_DEFAULT bool = false

func (p *ReplicaConfiguration) GetPopAll() bool {
	return p.PopAll
}

var ReplicaConfiguration_SplitSyncToChild_DEFAULT bool = false

func (p *ReplicaConfiguration) GetSplitSyncToChild() bool {
	return p.SplitSyncToChild
}

var ReplicaConfiguration_HpPrimary_DEFAULT *base.HostPort

func (p *ReplicaConfiguration) GetHpPrimary() *base.HostPort {
	if !p.IsSetHpPrimary() {
		return ReplicaConfiguration_HpPrimary_DEFAULT
	}
	return p.HpPrimary
}
func (p *ReplicaConfiguration) IsSetPid() bool {
	return p.Pid != nil
}

func (p *ReplicaConfiguration) IsSetPrimary() bool {
	return p.Primary != nil
}

func (p *ReplicaConfiguration) IsSetPopAll() bool {
	return p.PopAll != ReplicaConfiguration_PopAll_DEFAULT
}

func (p *ReplicaConfiguration) IsSetSplitSyncToChild() bool {
	return p.SplitSyncToChild != ReplicaConfiguration_SplitSyncToChild_DEFAULT
}

func (p *ReplicaConfiguration) IsSetHpPrimary() bool {
	return p.HpPrimary != nil
}

func (p *ReplicaConfiguration) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField5(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField6(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 7:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField7(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 8:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField8(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField1(iprot thrift.TProtocol) error {
	p.Pid = &base.Gpid{}
	if err := p.Pid.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Pid), err)
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Ballot = v
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField3(iprot thrift.TProtocol) error {
	p.Primary = &base.RPCAddress{}
	if err := p.Primary.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Primary), err)
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		temp := PartitionStatus(v)
		p.Status = temp
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.LearnerSignature = v
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.PopAll = v
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.SplitSyncToChild = v
	}
	return nil
}

func (p *ReplicaConfiguration) ReadField8(iprot thrift.TProtocol) error {
	p.HpPrimary = &base.HostPort{}
	if err := p.HpPrimary.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.HpPrimary), err)
	}
	return nil
}

func (p *ReplicaConfiguration) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("replica_configuration"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
		if err := p.writeField7(oprot); err != nil {
			return err
		}
		if err := p.writeField8(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ReplicaConfiguration) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("pid", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pid: ", p), err)
	}
	if err := p.Pid.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Pid), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pid: ", p), err)
	}
	return err
}

func (p *ReplicaConfiguration) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ballot", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ballot: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Ballot)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ballot (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ballot: ", p), err)
	}
	return err
}

func (p *ReplicaConfiguration) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("primary", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:primary: ", p), err)
	}
	if err := p.Primary.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Primary), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:primary: ", p), err)
	}
	return err
}

func (p *ReplicaConfiguration) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:status: ", p), err)
	}
	return err
}

func (p *ReplicaConfiguration) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("learner_signature", thrift.I64, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:learner_signature: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.LearnerSignature)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.learner_signature (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:learner_signature: ", p), err)
	}
	return err
}

func (p *ReplicaConfiguration) writeField6(oprot thrift.TProtocol) (err error) {
	if p.IsSetPopAll() {
		if err := oprot.WriteFieldBegin("pop_all", thrift.BOOL, 6); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:pop_all: ", p), err)
		}
		if err := oprot.WriteBool(bool(p.PopAll)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.pop_all (6) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 6:pop_all: ", p), err)
		}
	}
	return err
}

func (p *ReplicaConfiguration) writeField7(oprot thrift.TProtocol) (err error) {
	if p.IsSetSplitSyncToChild() {
		if err := oprot.WriteFieldBegin("split_sync_to_child", thrift.BOOL, 7); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:split_sync_to_child: ", p), err)
		}
		if err := oprot.WriteBool(bool(p.SplitSyncToChild)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.split_sync_to_child (7) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 7:split_sync_to_child: ", p), err)
		}
	}
	return err
}

func (p *ReplicaConfiguration) writeField8(oprot thrift.TProtocol) (err error) {
	if p.IsSetHpPrimary() {
		if err := oprot.WriteFieldBegin("hp_primary", thrift.STRUCT, 8); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:hp_primary: ", p), err)
		}
		if err := p.HpPrimary.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.HpPrimary), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 8:hp_primary: ", p), err)
		}
	}
	return err
}

func (p *ReplicaConfiguration) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReplicaConfiguration(%+v)", *p)
}

// Attributes:
//  - Pid
//  - Ballot
//  - Status
//  - LastCommittedDecree
//  - LastPreparedDecree
//  - LastDurableDecree
//  - AppType
//  - DiskTag
//  - ManualCompactStatus
type ReplicaInfo struct {
	Pid                 *base.Gpid              `thrift:"pid,1" db:"pid" json:"pid"`
	Ballot              int64                   `thrift:"ballot,2" db:"ballot" json:"ballot"`
	Status              PartitionStatus         `thrift:"status,3" db:"status" json:"status"`
	LastCommittedDecree int64                   `thrift:"last_committed_decree,4" db:"last_committed_decree" json:"last_committed_decree"`
	LastPreparedDecree  int64                   `thrift:"last_prepared_decree,5" db:"last_prepared_decree" json:"last_prepared_decree"`
	LastDurableDecree   int64                   `thrift:"last_durable_decree,6" db:"last_durable_decree" json:"last_durable_decree"`
	AppType             string                  `thrift:"app_type,7" db:"app_type" json:"app_type"`
	DiskTag             string                  `thrift:"disk_tag,8" db:"disk_tag" json:"disk_tag"`
	ManualCompactStatus *ManualCompactionStatus `thrift:"manual_compact_status,9" db:"manual_compact_status" json:"manual_compact_status,omitempty"`
}

func NewReplicaInfo() *ReplicaInfo {
	return &ReplicaInfo{}
}

var ReplicaInfo_Pid_DEFAULT *base.Gpid

func (p *ReplicaInfo) GetPid() *base.Gpid {
	if !p.IsSetPid() {
		return ReplicaInfo_Pid_DEFAULT
	}
	return p.Pid
}

func (p *ReplicaInfo) GetBallot() int64 {
	return p.Ballot
}

func (p *ReplicaInfo) GetStatus() PartitionStatus {
	return p.Status
}

func (p *ReplicaInfo) GetLastCommittedDecree() int64 {
	return p.LastCommittedDecree
}

func (p *ReplicaInfo) GetLastPreparedDecree() int64 {
	return p.LastPreparedDecree
}

func (p *ReplicaInfo) GetLastDurableDecree() int64 {
	return p.LastDurableDecree
}

func (p *ReplicaInfo) GetAppType() string {
	return p.AppType
}

func (p *ReplicaInfo) GetDiskTag() string {
	return p.DiskTag
}

var ReplicaInfo_ManualCompactStatus_DEFAULT ManualCompactionStatus

func (p *ReplicaInfo) GetManualCompactStatus() ManualCompactionStatus {
	if !p.IsSetManualCompactStatus() {
		return ReplicaInfo_ManualCompactStatus_DEFAULT
	}
	return *p.ManualCompactStatus
}
func (p *ReplicaInfo) IsSetPid() bool {
	return p.Pid != nil
}

func (p *ReplicaInfo) IsSetManualCompactStatus() bool {
	return p.ManualCompactStatus != nil
}

func (p *ReplicaInfo) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField3(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField4(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField5(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField6(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField7(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 8:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField8(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		case 9:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField9(iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ReplicaInfo) ReadField1(iprot thrift.TProtocol) error {
	p.Pid = &base.Gpid{}
	if err := p.Pid.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Pid), err)
	}
	return nil
}

func (p *ReplicaInfo) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Ballot = v
	}
	return nil
}

func (p *ReplicaInfo) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := PartitionStatus(v)
		p.Status = temp
	}
	return nil
}

func (p *ReplicaInfo) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.LastCommittedDecree = v
	}
	return nil
}

func (p *ReplicaInfo) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.LastPreparedDecree = v
	}
	return nil
}

func (p *ReplicaInfo) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.LastDurableDecree = v
	}
	return nil
}

func (p *ReplicaInfo) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.AppType = v
	}
	return nil
}

func (p *ReplicaInfo) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 8: ", err)
	} else {
		p.DiskTag = v
	}
	return nil
}

func (p *ReplicaInfo) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 9: ", err)
	} else {
		temp := ManualCompactionStatus(v)
		p.ManualCompactStatus = &temp
	}
	return nil
}

func (p *ReplicaInfo) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("replica_info"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
		if err := p.writeField2(oprot); err != nil {
			return err
		}
		if err := p.writeField3(oprot); err != nil {
			return err
		}
		if err := p.writeField4(oprot); err != nil {
			return err
		}
		if err := p.writeField5(oprot); err != nil {
			return err
		}
		if err := p.writeField6(oprot); err != nil {
			return err
		}
		if err := p.writeField7(oprot); err != nil {
			return err
		}
		if err := p.writeField8(oprot); err != nil {
			return err
		}
		if err := p.writeField9(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ReplicaInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("pid", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pid: ", p), err)
	}
	if err := p.Pid.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Pid), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pid: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ballot", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ballot: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Ballot)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ballot (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ballot: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("status", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:status: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Status)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.status (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:status: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("last_committed_decree", thrift.I64, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:last_committed_decree: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.LastCommittedDecree)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.last_committed_decree (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:last_committed_decree: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("last_prepared_decree", thrift.I64, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:last_prepared_decree: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.LastPreparedDecree)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.last_prepared_decree (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:last_prepared_decree: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("last_durable_decree", thrift.I64, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:last_durable_decree: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.LastDurableDecree)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.last_durable_decree (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:last_durable_decree: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("app_type", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:app_type: ", p), err)
	}
	if err := oprot.WriteString(string(p.AppType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.app_type (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:app_type: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("disk_tag", thrift.STRING, 8); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:disk_tag: ", p), err)
	}
	if err := oprot.WriteString(string(p.DiskTag)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.disk_tag (8) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 8:disk_tag: ", p), err)
	}
	return err
}

func (p *ReplicaInfo) writeField9(oprot thrift.TProtocol) (err error) {
	if p.IsSetManualCompactStatus() {
		if err := oprot.WriteFieldBegin("manual_compact_status", thrift.I32, 9); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:manual_compact_status: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.ManualCompactStatus)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.manual_compact_status (9) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 9:manual_compact_status: ", p), err)
		}
	}
	return err
}

func (p *ReplicaInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReplicaInfo(%+v)", *p)
}
