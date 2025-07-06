#!/usr/bin/env python3
# coding: utf-8


import logging
import typing
import os
from concurrent import futures

import grpc
from digimat.saia import SAIANode, SAIAServer
from digimat.saia.memory import (
    SAIAItemFlag,
    SAIAItemInput,
    SAIAItemOutput,
    SAIAItemRegister,
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

    serve()
