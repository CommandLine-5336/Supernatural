"""
ASGI config for core project.

It exposes the ASGI callable as a module-level variable named ``application``.

For more information on this file, see
https://docs.djangoproject.com/en/6.0/howto/deployment/asgi/
"""

import os

from django.core.asgi import get_asgi_application
from starlette.applications import Starlette
from starlette.requests import Request
from starlette.responses import RedirectResponse
from starlette.routing import Mount, Route

os.environ.setdefault("DJANGO_SETTINGS_MODULE", "core.settings")

django_application = get_asgi_application()

from apps.clause.tools import mcp

mcp_app = mcp.http_app(path="/")


def _redirect_trailing_slash(request: Request) -> RedirectResponse:
    return RedirectResponse(str(request.url.replace(path="/mcp/")), status_code=307)


application = Starlette(
    routes=[
        Route(
            "/mcp", endpoint=_redirect_trailing_slash, methods=["GET", "POST", "DELETE"]
        ),
        Mount("/mcp", app=mcp_app),
        Mount("/", app=django_application),
    ],
    lifespan=mcp_app.lifespan,
)
