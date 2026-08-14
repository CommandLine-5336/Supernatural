from django.core.exceptions import PermissionDenied
from fastmcp import Context, FastMCP
from fastmcp.exceptions import ToolError

from . import services
from .middleware import UserAuthMiddleware

mcp = FastMCP("Circle Tools", middleware=[UserAuthMiddleware()])


def _call(service_fn, ctx: Context, **kwargs) -> str:
    user = ctx.get_state("user")
    try:
        return service_fn(user, **kwargs)
    except PermissionDenied as exc:
        raise ToolError(str(exc)) from exc


@mcp.tool
def create_vote(ctx: Context, vote_type: str, user_alias: str, description: str) -> str:
    """Create vote"""
    return _call(
        services.create_vote,
        ctx,
        vote_type=vote_type,
        user_alias=user_alias,
        description=description,
    )


@mcp.tool
def set_vote(ctx: Context, vote_id: int, stat: str) -> str:
    """Set vote"""
    return _call(services.set_vote, ctx, vote_id=vote_id, stat=stat)


@mcp.tool
def grade(ctx: Context, user_alias: str, stat: str) -> str:
    """Upgrade/downgrade user"""
    return _call(services.grade, ctx, user_alias=user_alias, stat=stat)


@mcp.tool
def create_post(
    ctx: Context, name: str, description: str, latitude: str, longitude: str
) -> str:
    """Create post"""
    return _call(
        services.create_post,
        ctx,
        name=name,
        description=description,
        latitude=latitude,
        longitude=longitude,
    )


@mcp.tool
def seen_post(ctx: Context, post_id: int) -> str:
    """Mark post as seen"""
    return _call(services.seen_post, ctx, post_id=post_id)


@mcp.tool
def report_ip(ctx: Context, ip_address: str) -> str:
    """Report ip address"""
    return _call(services.report_ip, ctx, ip_address=ip_address)


@mcp.tool
def invite(ctx: Context, email: str) -> str:
    """Send invite"""
    return _call(services.invite, ctx, email=email)
