import os

import jwt
from rest_framework.authentication import BaseAuthentication
from rest_framework.exceptions import AuthenticationFailed

from ..authentication.models import User

JWT_SECRET = os.getenv("JWT_KEY")


class CookieJWTAuthentication(BaseAuthentication):
    def authenticate(self, request):
        token = request.COOKIES.get("jwt")
        if not token:
            return None

        if isinstance(token, bytes):
            token = token.decode("utf-8")

        try:
            payload = jwt.decode(token, JWT_SECRET, algorithms=["HS256"])
            user_id = payload.get("sub")
            if not user_id:
                raise AuthenticationFailed("Token payload missing 'sub'")

            user = User.objects.get(pk=int(user_id))
            user.is_authenticated = True

        except jwt.ExpiredSignatureError:
            raise AuthenticationFailed("Token expired")
        except jwt.InvalidTokenError:
            raise AuthenticationFailed("Invalid token")
        except jwt.DecodeError:
            raise AuthenticationFailed("Invalid format")

        except User.DoesNotExist:
            raise AuthenticationFailed("User not found")

        return (user, token)

    def authenticate_header(self, request):
        return "Bearer"
