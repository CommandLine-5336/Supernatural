"""Tools for ai agent"""

import ipaddress

import requests
from django.core.exceptions import PermissionDenied

from ..authentication.models import *
from ..blocking.models import *
from ..posts.models import *
from ..posts.views import viewsets
from ..votes.models import *


def seen_post(user, post_id: int) -> str:
    """Set seen in post"""
    post = Post.objects.get(pk=post_id)
    if Report.objects.filter(user=user, post=post).exists():
        return "You have already marked this post as seen"

    Report.objects.create(user=user, post=post)
    post.seen_count += 1
    post.save(update_fields=["seen_count"])
    return f"Marked post {post.name} as seen"


def create_post(
    user, name: str, description: str, latitude: str, longitude: str
) -> str:
    """Create post"""
    if user.status == "copper":
        raise PermissionDenied("You cannot create posts")

    Post.objects.create(
        name=name,
        description=description,
        latitude=latitude,
        longitude=longitude,
    )
    return f"Post created"


def set_vote(user, vote_id: int, stat: str) -> str:
    """Set vote"""
    vote = Vote.objects.get(pk=vote_id)
    if VoteRes.objects.filter(user=user, vote=vote).exists():
        return "You have already voted on this"

    if stat == "+":
        vote.agree = vote.agree + 1
    else:
        vote.disagree = vote.disagree + 1

    vote.save(update_fields=["agree", "disagree"])
    VoteRes.objects.create(user=user, vote=vote)
    return f"Vote {stat} set"


def create_vote(user, vote_type: str, user_alias: str, description: str) -> str:
    """Create vote"""

    try:
        user_v = User.objects.get(alias=user_alias)
    except User.DoesNotExist:
        return f"There is no user with alias '{user_alias}'"

    if vote_type == "promotion":
        if user.status == "copper":
            raise PermissionDenied("You cannot create votes")

    elif vote_type == "excommunication":
        if not user.inquisitor:
            raise PermissionDenied("You cannot create excommunication votes")

    elif vote_type == "architect":
        if not user.is_architect:
            raise PermissionDenied("You cannot create architect votes")
        if user_v.status != "gold":
            raise PermissionDenied("only gold users can be nominated architect")

    else:
        return f"Unknown vote type: {vote_type}"

    Vote.objects.create(
        type=vote_type,
        user=user_v,
        description=description,
        agree=0,
        disagree=0,
    )
    return f"{vote_type} vote for '{user_alias}' successfully created"


def grade(user, user_alias: str, stat: str) -> str:
    """Grade user"""
    if not user.is_architect:
        raise PermissionDenied("You cannot upgrade/downgrade users")

    try:
        user_g = User.objects.get(alias=user_alias)
    except User.DoesNotExist:
        return f"There is no user with alias '{user_alias}'"

    if stat == "up":
        if user_g.status == "copper":
            user_g.status = "silver"
        else:
            user_g.status = "gold"
    else:
        if user_g.status == "gold":
            user_g.status = "silver"
        else:
            user_g.status = "copper"

    user_g.save(update_fields=["status"])
    return f"{user_alias} is now {user_g.status}."


def invite(user, email: str) -> str:
    """Invite user"""
    if user.status != "gold":
        raise PermissionDenied("You cannot send invites")

    try:
        resp = requests.post(
            f"http://mail_service:8074/invite",
            json={"email": email},
            timeout=5,
        )
        resp.raise_for_status()
    except requests.RequestException as exc:
        return f"Could not send invite: {exc}"

    return f"Invite successfully sent to {email}"


def report_ip(user, ip_address: str) -> str:
    """Report ip"""
    if user.status == "copper":
        raise PermissionDenied("You cannot report IP addresses")

    try:
        ipaddress.IPv4Address(ip_address)
    except ValueError:
        return "Invalid IPv4 address"

    if Banned.objects.filter(ip_address=ip_address).exists():
        return "IP is already banned"

    Banned.objects.create(ip_address=ip_address)
    return f"IP address {ip_address} successfully reported"


def compromised(user):
    """Cleanup"""
    raise PermissionDenied(
        "This action can only be done manually. Please go to admin page to do this"
    )
