#
# Autogenerated by Thrift Compiler (0.14.2)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#
#  options string: py
#

from thrift.Thrift import TType, TMessageType, TFrozenDict, TException, TApplicationException
from thrift.protocol.TProtocol import TProtocolException
from thrift.TRecursive import fix_spec

import sys
import base_service_v1.ttypes

from thrift.transport import TTransport
all_structs = []


class Location(object):
    """
    Attributes:
     - timestamp_unix_sec
     - latitude_degrees
     - longitude_degrees

    """


    def __init__(self, timestamp_unix_sec=None, latitude_degrees=None, longitude_degrees=None,):
        self.timestamp_unix_sec = timestamp_unix_sec
        self.latitude_degrees = latitude_degrees
        self.longitude_degrees = longitude_degrees

    def read(self, iprot):
        if iprot._fast_decode is not None and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None:
            iprot._fast_decode(self, iprot, [self.__class__, self.thrift_spec])
            return
        iprot.readStructBegin()
        while True:
            (fname, ftype, fid) = iprot.readFieldBegin()
            if ftype == TType.STOP:
                break
            if fid == 1:
                if ftype == TType.DOUBLE:
                    self.timestamp_unix_sec = iprot.readDouble()
                else:
                    iprot.skip(ftype)
            elif fid == 2:
                if ftype == TType.DOUBLE:
                    self.latitude_degrees = iprot.readDouble()
                else:
                    iprot.skip(ftype)
            elif fid == 3:
                if ftype == TType.DOUBLE:
                    self.longitude_degrees = iprot.readDouble()
                else:
                    iprot.skip(ftype)
            else:
                iprot.skip(ftype)
            iprot.readFieldEnd()
        iprot.readStructEnd()

    def write(self, oprot):
        if oprot._fast_encode is not None and self.thrift_spec is not None:
            oprot.trans.write(oprot._fast_encode(self, [self.__class__, self.thrift_spec]))
            return
        oprot.writeStructBegin('Location')
        if self.timestamp_unix_sec is not None:
            oprot.writeFieldBegin('timestamp_unix_sec', TType.DOUBLE, 1)
            oprot.writeDouble(self.timestamp_unix_sec)
            oprot.writeFieldEnd()
        if self.latitude_degrees is not None:
            oprot.writeFieldBegin('latitude_degrees', TType.DOUBLE, 2)
            oprot.writeDouble(self.latitude_degrees)
            oprot.writeFieldEnd()
        if self.longitude_degrees is not None:
            oprot.writeFieldBegin('longitude_degrees', TType.DOUBLE, 3)
            oprot.writeDouble(self.longitude_degrees)
            oprot.writeFieldEnd()
        oprot.writeFieldStop()
        oprot.writeStructEnd()

    def validate(self):
        return

    def __repr__(self):
        L = ['%s=%r' % (key, value)
             for key, value in self.__dict__.items()]
        return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

    def __eq__(self, other):
        return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not (self == other)
all_structs.append(Location)
Location.thrift_spec = (
    None,  # 0
    (1, TType.DOUBLE, 'timestamp_unix_sec', None, None, ),  # 1
    (2, TType.DOUBLE, 'latitude_degrees', None, None, ),  # 2
    (3, TType.DOUBLE, 'longitude_degrees', None, None, ),  # 3
)
fix_spec(all_structs)
del all_structs
