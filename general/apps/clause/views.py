"""Views for AI agent"""

import asyncio

from django.core import signing
from fastmcp import Client
from google import genai
from google.genai import types
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response
from rest_framework.views import APIView

from ..token import CookieJWTAuthentication


def _load_history(request) -> list[dict]:
    raw = request.COOKIES.get("history")
    if not raw:
        return []
    try:
        return signing.loads(raw, max_age=60 * 60 * 5)
    except signing.BadSignature:
        return []


def _save_history(response, history: list[dict]) -> None:
    response.set_cookie(
        "history",
        signing.dumps(history),
        max_age=60 * 60 * 5,
        httponly=True,
        secure=True,
        samesite="Lax",
    )


async def _run_chat_turn(message: str, history: list[dict], jwt_token: str) -> str:
    async with Client(
        "http://general-service:4040/mcp/",
        headers={"Authorization": f"Bearer {jwt_token}"},
    ) as mcp_client:
        client = genai.Client()
        contents = [
            types.Content(role=h["role"], parts=[types.Part(text=h["text"])])
            for h in history
        ]
        contents.append(types.Content(role="user", parts=[types.Part(text=message)]))

        response = await client.aio.models.generate_content(
            model="gemini-flash-2.5",
            contents=contents,
            config=types.GenerateContentConfig(tools=[mcp_client]),
        )
        return response.text


class ClauseView(APIView):
    authentication_classes = [CookieJWTAuthentication]
    permission_classes = [IsAuthenticated]

    def post(self, request):
        message = request.data.get("message")
        if not message:
            return Response({"message": "'there is no message"}, status=400)

        jwt_token = request.auth
        history = _load_history(request)

        reply_text = asyncio.run(_run_chat_turn(message, history, jwt_token))

        history.append({"role": "user", "text": message})
        history.append({"role": "model", "text": reply_text})

        response = Response({"status": "ok", "reply": reply_text})
        _save_history(response, history)
        return response
