import asyncio
import logging

from concurrent import futures
import grpc
# import time
# import os
import py_client_pb2
import py_client_pb2_grpc

# Coroutines to be invoked when the event loop is shutting down.
_cleanup_coroutines = []

class Request(py_client_pb2_grpc.ClientRequestServicer):

    async def SendRequest(
            self, request: py_client_pb2.Request,
            context: grpc.aio.ServicerContext) -> py_client_pb2.Response:
        print('Got request from a client', request)
        return py_client_pb2.Response(status='Success, %s!' % request.registration, outputUrl='testurl')

    def UploadToServer(self, request_iterator, context):
        data = bytearray()
        filepath = './server_out.txt'
        for request in request_iterator:
            # if request.metadata.filename and request.metadata.extension:
            #     filepath = get_filepath(request.metadata.filename, request.metadata.extension)
            #     continue
            data.extend(request.chunk)
        with open(filepath, 'wb') as f:
            f.write(data)
            print('Finished writing the file onto the server')
            # time.sleep(20)
        return py_client_pb2.UploadResponse(name='Success!')

async def serve() -> None:
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=4))
    py_client_pb2_grpc.add_ClientRequestServicer_to_server(Request(), server)
    listen_addr = '[::]:50051'
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()

    async def server_graceful_shutdown():
        logging.info("Starting graceful shutdown...")
        # Shuts down the server with 5 seconds of grace period. During the
        # grace period, the server won't accept new connections and allow
        # existing RPCs to continue within the grace period.
        await server.stop(5)
    
    _cleanup_coroutines.append(server_graceful_shutdown())
    await server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(serve())
    finally:
        loop.run_until_complete(*_cleanup_coroutines)
        loop.close()
