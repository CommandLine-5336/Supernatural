import os

import jwt as pyjwt
from fastmcp.exceptions import ToolError
from fastmcp.server.dependencies import get_http_headers
from fastmcp.server.middleware import Middleware, MiddlewareContext

from ..authentication.models import User

JWT_SECRET = os.getenv("JWT_KEY")


class UserAuthMiddleware(Middleware):
    """Gets user from JWT token"""

    async def on_call_tool(self, context: MiddlewareContext, call_next):
        headers = get_http_headers()
        token = headers.get("authorization", "").removeprefix("Bearer ").strip()
        if not token:
            raise ToolError("User hasn't been authorized")

        try:
            payload = pyjwt.decode(token, JWT_SECRET, algorithms=["HS256"])
            user = await User.objects.aget(pk=int(payload["sub"]))
        except Exception as exc:
            raise ToolError("Token is invalid or expired") from exc

        context.fastmcp_context.set_state("user", user)
        return await call_next(context)
