from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

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
    __slots__ = ("address", "as_int", "as_float")
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    AS_INT_FIELD_NUMBER: _ClassVar[int]
    AS_FLOAT_FIELD_NUMBER: _ClassVar[int]
    address: int
    as_int: bool
    as_float: bool
    def __init__(self, address: _Optional[int] = ..., as_int: bool = ..., as_float: bool = ...) -> None: ...

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
    __slots__ = ("address", "int_value", "float_value")
    ADDRESS_FIELD_NUMBER: _ClassVar[int]
    INT_VALUE_FIELD_NUMBER: _ClassVar[int]
    FLOAT_VALUE_FIELD_NUMBER: _ClassVar[int]
    address: int
    int_value: int
    float_value: float
    def __init__(self, address: _Optional[int] = ..., int_value: _Optional[int] = ..., float_value: _Optional[float] = ...) -> None: ...

class WriteRegisterResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
