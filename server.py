import asyncio, asyncssh
from session_manager import handle_session

class MySSHServer(asyncssh.SSHServer):
    def connection_made(self, conn):
        print("Someone connected!")

    def begin_auth(self, username):
        return False

async def main():
    await asyncssh.create_server(
        MySSHServer,
        host="localhost",
        port=2222,
        authorized_client_keys=None,
        server_host_keys=["ssh_host_key"],
        process_factory=handle_session,
        kex_algs=["mlkem768x25519-sha256", "curve25519-sha256"],
    )
    print("Server running on port 2222...")
    await asyncio.get_event_loop().create_future()

asyncio.run(main())