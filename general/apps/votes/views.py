"""Views for votes"""

from rest_framework import filters, status, viewsets
from rest_framework.decorators import action
from rest_framework.exceptions import NotFound, ValidationError
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response

from .models import User, Vote, VoteRes
from .serializers import VoteResSerializer, VoteSerializer
from .token import CookieJWTAuthentication


def execute_vote(vote):
    """Deleting vote and applying its results"""
    user = vote.user
    total_users = User.objects.count()
    if user is not None:
        if vote.type == "promotion":
            if vote.agree > total_users / 2:
                if user.status == "copper":
                    user.status = "silver"
                else:
                    user.status = "gold"
                user.save(update_fields=["status"])

        elif vote.type == "excommunication":
            if vote.agree > total_users * 0.8:
                user.banned = True
                user.save(update_fields=["banned"])

    VoteRes.objects.filter(vote=vote).delete()
    vote.delete()


class VoteViewSet(viewsets.ModelViewSet):
    """ViewSet for Vote"""

    http_method_names = ["get", "post", "put", "patch", "delete"]
    serializer_class = VoteSerializer
    permission_classes = (IsAuthenticated,)
    authentication_classes = [CookieJWTAuthentication]
    filter_backends = [filters.OrderingFilter]
    ordering_fields = ["id", "agree", "disagree"]
    ordering = ["-id"]
    queryset = Vote.objects.all()

    def list(self, request):
        """List of all votes and votes on which specific user has already voted"""
        votes_queryset = Vote.objects.all().select_related("user")
        votes_serializer = VoteSerializer(votes_queryset, many=True)
        user_votes_queryset = VoteRes.objects.filter(user=request.user)
        user_votes_serializer = VoteResSerializer(user_votes_queryset, many=True)
        return Response(
            {
                "votes": votes_serializer.data,
                "user_voted": user_votes_serializer.data,
            }
        )

    def perform_create(self, serializer):
        """Create new vote"""
        user_alias = self.request.data.get("user_alias")
        if not user_alias:
            raise ValidationError({"user_alias": "user_alias is required"})
        try:
            user = User.objects.get(alias=user_alias)
        except User.DoesNotExist as exc:
            raise NotFound(
                {"detail": f"User with alias '{user_alias}' does not exist"}
            ) from exc
        serializer.save(user=user)

    @action(detail=True, methods=["post", "put"])
    def set_vote(self, request, pk=None):  # [unused-argument]
        """Set vote"""
        vote = self.get_object()
        res = request.data.get("res")
        user = request.user
        if VoteRes.objects.filter(user=user, vote=vote).exists():
            return Response(
                {"detail": "You have already set vote here"},
                status=status.HTTP_400_BAD_REQUEST,
            )
        if res == "+":
            vote.agree = vote.agree + 1
        elif res == "-":
            vote.disagree = vote.disagree + 1
        else:
            return Response(
                {"detail": "res can only be + or -"},
                status=status.HTTP_400_BAD_REQUEST,
            )

        vote.save(update_fields=["agree", "disagree"])
        VoteRes.objects.create(user=user, vote=vote)
        serializer = self.get_serializer(vote)
        return Response(serializer.data)

    def destroy(self, _request, *_args, **_kwargs):
        """Delete vote"""
        vote = self.get_object()
        execute_vote(vote)
        return Response(
            {"detail": "Vote is deleted and results are applied"},
            status=status.HTTP_204_NO_CONTENT,
        )
