import ipaddress
import json
import os

import jwt
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_POST

from .models import Banned


def get_user_from_JWT(request):
    """Since Maksym stores user information inside JWT i need to get information from JWT cookie
    https://pyjwt.readthedocs.io/en/stable/usage.html#encoding-decoding-tokens-with-hs256
    https://www.youtube.com/watch?v=T14xCFlAUuM
    https://stackoverflow.com/questions/5113660/how-to-set-or-get-a-cookie-value-in-django

    returns entire payload to be used in the future.
    """
    token = request.COOKIES.get("jwt")
    if not token:
        return None

    try:
        payload = jwt.decode(token, os.environ["JWT_SECRET"], algorithms=["HS256"])
    except jwt.PyJWTError:
        return None

    return payload


@csrf_exempt
@require_POST
def report_ip(request):
    """This is where well be getting our IP and putting it into a DB
    First i need to parse the jwt and get a user status
    My guess is that it will be great to get IP as a JSON this is the way i found to decode a json request
    https://stackoverflow.com/questions/19573747/parsing-json-fields-in-python
    www.geeksforgeeks.org/python/creating-a-json-response-using-django-and-python/
    https://docs.djangoproject.com/en/6.0/ref/models/querysets/
    https://stackoverflow.com/questions/32848472/better-option-to-check-if-a-particular-instance-exists-django
    https://www.youtube.com/watch?v=Da5abtjf0Bg&t=22s

    """
    payload = get_user_from_JWT(request)
    if payload is None:
        return JsonResponse(
            {"status": "error", "message": "unauthorized or invalid cookie"}, status=401
        )

    jwtstatus = payload.get("status")
    if jwtstatus not in ("silver", "gold", "inquisitor"):
        return JsonResponse(
            {"status": "error", "message": "forbidden, not enough priviliges"},
            status=403,
        )

    try:
        jsonbody = json.loads(request.body)
        ip_address = jsonbody["ip_address"]
    except (json.JSONDecodeError, KeyError):
        return JsonResponse(
            {"status": "error", "message": "Bad request, No ipv4 was provided"},
            status=400,
        )

    try:
        ipaddress.IPv4Address(ip_address)

    except ValueError:
        return JsonResponse(
            {"status": "error", "message": "Bad request, Invalid ip address"},
            status=400,
        )

    if Banned.objects.filter(ip_address=ip_address).exists():
        return JsonResponse(
            {"status": "ok", "message": "IP already exists in DB"}, status=200
        )

    Banned.objects.create(ip_address=ip_address)
    return JsonResponse(
        {"status": "ok", "message": "IP reported and banned"}, status=200
    )
