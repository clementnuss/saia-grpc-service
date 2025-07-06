#!/usr/bin/env python3
# coding: utf-8


import logging
import typing
import os
from concurrent import futures

import grpc
from digimat.saia import SAIANode, SAIAServer
from digimat.saia.memory import (
    SAIAItemCounter,
    SAIAItemFlag,
    SAIAItemInput,
    SAIAItemOutput,
    SAIAItemRegister,
    SAIAItemTimer,
)
from gen import saia_pb2, saia_pb2_grpc
from dotenv import load_dotenv

# Enable logging
logging.basicConfig(
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s", level=logging.INFO
)
logger = logging.getLogger(__name__)


class SaiaPcdServicer(saia_pb2_grpc.SaiaPcdServiceServicer):
    def ReadRegister(
        self, request: saia_pb2.ReadRegisterRequest, context: grpc.ServicerContext
    ) -> saia_pb2.ReadRegisterResponse:
        r = typing.cast(SAIAItemRegister, server.registers[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to read register value")

        resp = saia_pb2.ReadRegisterResponse()
        if request.data_type == saia_pb2.REGISTER_DATA_TYPE_INT:
            resp.int_value = r.value
        elif request.data_type == saia_pb2.REGISTER_DATA_TYPE_FLOAT:
            resp.float_value = r.float
        else:
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "missing data_type")

        return resp

    def ReadFlag(
        self, request: saia_pb2.ReadFlagRequest, context: grpc.ServicerContext
    ) -> saia_pb2.ReadFlagResponse:
        r = typing.cast(SAIAItemFlag, server.flags[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to read register value")

        return saia_pb2.ReadFlagResponse(value=r.bool)

    def ReadInput(
        self, request: saia_pb2.ReadInputRequest, context: grpc.ServicerContext
    ) -> saia_pb2.ReadInputResponse:
        r = typing.cast(SAIAItemInput, server.inputs[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to read register value")

        return saia_pb2.ReadInputResponse(value=r.bool)

    def ReadOutput(
        self, request: saia_pb2.ReadOutputRequest, context: grpc.ServicerContext
    ) -> saia_pb2.ReadOutputResponse:
        r = typing.cast(SAIAItemOutput, server.outputs[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to read register value")

        return saia_pb2.ReadOutputResponse(value=r.bool)

    def ReadCounter(
        self, request: saia_pb2.ReadCounterRequest, context: grpc.ServicerContext
    ) -> saia_pb2.ReadCounterResponse:
        r = typing.cast(SAIAItemCounter, server.counters[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to read counter value")

        return saia_pb2.ReadCounterResponse(value=r.value)

    def ReadTimer(
        self, request: saia_pb2.ReadTimerRequest, context: grpc.ServicerContext
    ) -> saia_pb2.ReadTimerResponse:
        r = typing.cast(SAIAItemTimer, server.timers[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to read timer value")

        return saia_pb2.ReadTimerResponse(value=r.value)

    def WriteFlag(
        self, request: saia_pb2.WriteFlagRequest, context: grpc.ServicerContext
    ) -> saia_pb2.WriteFlagResponse:
        r = typing.cast(SAIAItemFlag, server.flags[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to write flag value")

        r.bool = request.value
        return saia_pb2.WriteFlagResponse()

    def WriteRegister(
        self, request: saia_pb2.WriteRegisterRequest, context: grpc.ServicerContext
    ) -> saia_pb2.WriteRegisterResponse:
        r = typing.cast(SAIAItemRegister, server.registers[request.address])

        if r is None or not r.isAlive():
            context.abort(grpc.StatusCode.INTERNAL, "unable to write register value")

        if request.data_type == saia_pb2.REGISTER_DATA_TYPE_INT:
            if not request.HasField("int_value"):
                context.abort(grpc.StatusCode.INVALID_ARGUMENT, "int_value required for INT data type")
            r.value = request.int_value
        elif request.data_type == saia_pb2.REGISTER_DATA_TYPE_FLOAT:
            if not request.HasField("float_value"):
                context.abort(grpc.StatusCode.INVALID_ARGUMENT, "float_value required for FLOAT data type")
            r.float = request.float_value
        else:
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "missing or invalid data_type")

        return saia_pb2.WriteRegisterResponse()


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    saia_pb2_grpc.add_SaiaPcdServiceServicer_to_server(SaiaPcdServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    node = SAIANode(253, logger=logger)

    load_dotenv()

    server_address = os.getenv("SAIA_SERVER_ADDRESS", "")
    server_station = os.getenv("SAIA_SERVER_STATION", "")
    if server_address == "" or server_station == "":
        logger.error("SAIA server address or station not configured")
    node.servers.declare(server_address, int(server_station))

    global server
    server = typing.cast(SAIAServer, node.servers[0])
    server.memory.registers.setRefreshDelay(15)
    server.memory.inputs.setRefreshDelay(15)
    server.memory.outputs.setRefreshDelay(15)
    server.memory.flags.setRefreshDelay(15)
    server.memory.counters.setRefreshDelay(15)
    server.memory.timers.setRefreshDelay(15)

    serve()
