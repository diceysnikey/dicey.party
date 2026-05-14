import asyncio, asyncssh

async def handle_session(process):
    process.stdout.write("Hello! Welcome to my server.\n")
    process.exit(0)