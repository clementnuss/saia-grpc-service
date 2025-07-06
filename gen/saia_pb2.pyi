from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RegisterDataType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    REGISTER_DATA_TYPE_UNSPECIFIED: _ClassVar[RegisterDataType]
    REGISTER_DATA_TYPE_INT: _ClassVar[RegisterDataType]
    REGISTER_DATA_TYPE_FLOAT: _ClassVar[RegisterDataType]
REGISTER_DATA_TYPE_UNSPECIFIED: RegisterDataType
REGISTER_DATA_TYPE_INT: RegisterDataType
REGISTER_DATA_TYPE_FLOAT: RegisterDataType

class ReadInputRequest(_message.Message):
    __slots__ = ("address",)
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    address: int
    def __init__(self, address: _Optional[int] = ...) -> None: ...

class ReadOutputRequest(_message.Message):
    __slots__ = ("address",)
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    address: int
    def __init__(self, address: _Optional[int] = ...) -> None: ...

class ReadFlagRequest(_message.Message):
    __slots__ = ("address",)
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    address: int
    def __init__(self, address: _Optional[int] = ...) -> None: ...

class ReadCounterRequest(_message.Message):
    __slots__ = ("address",)
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    address: int
    def __init__(self, address: _Optional[int] = ...) -> None: ...

class ReadTimerRequest(_message.Message):
    __slots__ = ("address",)
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    address: int
    def __init__(self, address: _Optional[int] = ...) -> None: ...

class ReadRegisterRequest(_message.Message):
    __slots__ = ("address", "data_type")
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    DATA_TYPE_FIELD_NUMBER: _ClassVar[int]
    address: int
    data_type: RegisterDataType
    def __init__(self, address: _Optional[int] = ..., data_type: _Optional[_Union[RegisterDataType, str]] = ...) -> None: ...

class ReadInputResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: bool
    def __init__(self, value: bool = ...) -> None: ...

class ReadOutputResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: bool
    def __init__(self, value: bool = ...) -> None: ...

class ReadFlagResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: bool
    def __init__(self, value: bool = ...) -> None: ...

class ReadCounterResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: int
    def __init__(self, value: _Optional[int] = ...) -> None: ...

class ReadTimerResponse(_message.Message):
    __slots__ = ("value",)
    VALUE_FIELD_NUMBER: _ClassVar[int]
    value: int
    def __init__(self, value: _Optional[int] = ...) -> None: ...

class ReadRegisterResponse(_message.Message):
    __slots__ = ("int_value", "float_value")
    INT_VALUE_FIELD_NUMBER: _ClassVar[int]
    FLOAT_VALUE_FIELD_NUMBER: _ClassVar[int]
    int_value: int
    float_value: float
    def __init__(self, int_value: _Optional[int] = ..., float_value: _Optional[float] = ...) -> None: ...

class WriteFlagRequest(_message.Message):
    __slots__ = ("address", "value")
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    address: int
    value: bool
    def __init__(self, address: _Optional[int] = ..., value: bool = ...) -> None: ...

class WriteFlagResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class WriteRegisterRequest(_message.Message):
    __slots__ = ("address", "data_type", "int_value", "float_value")
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    DATA_TYPE_FIELD_NUMBER: _ClassVar[int]
    INT_VALUE_FIELD_NUMBER: _ClassVar[int]
    FLOAT_VALUE_FIELD_NUMBER: _ClassVar[int]
    address: int
    data_type: RegisterDataType
    int_value: int
    float_value: float
    def __init__(self, address: _Optional[int] = ..., data_type: _Optional[_Union[RegisterDataType, str]] = ..., int_value: _Optional[int] = ..., float_value: _Optional[float] = ...) -> None: ...

class WriteRegisterResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
