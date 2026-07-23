from rest_framework import filters, status, viewsets
from rest_framework.decorators import action
from rest_framework.exceptions import NotFound, ValidationError
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response

from .models import User, Vote, Vote_res
from .serializers import VoteResSerializer, VoteSerializer
from .token import CookieJWTAuthentication


class VoteViewSet(viewsets.ModelViewSet):
    http_method_names = ["get", "post", "put", "patch", "delete"]
    serializer_class = VoteSerializer
    permission_classes = (IsAuthenticated,)
    authentication_classes = [CookieJWTAuthentication]
    filter_backends = [filters.OrderingFilter]
    ordering_fields = ["id", "agree", "disagree"]
    ordering = ["-id"]
    queryset = Vote.objects.all()

    def list(self, request):
        votes_queryset = Vote.objects.all().select_related("user")
        votes_serializer = VoteSerializer(votes_queryset, many=True)
        user_votes_queryset = Vote_res.objects.filter(user=request.user)
        user_votes_serializer = VoteResSerializer(user_votes_queryset, many=True)
        return Response(
            {
                "votes": votes_serializer.data,
                "user_voted": user_votes_serializer.data,
                "me": {
                    "id": request.user.id,
                    "status": request.user.status,
                },
            }
        )

    def perform_create(self, serializer):
        user_alias = self.request.data.get("user_alias")
        if not user_alias:
            raise ValidationError({"user_alias": "user_alias is required"})
        try:
            user = User.objects.get(alias=user_alias)
        except User.DoesNotExist:
            raise NotFound({"detail": f"User with alias '{user_alias}' does not exist"})
        serializer.save(user=user)

    @action(detail=True, methods=["post", "put"])
    def set_vote(self, request, pk=None):
        vote = self.get_object()
        res = request.data.get("res")
        user = request.user
        if Vote_res.objects.filter(user=user, vote=vote).exists():
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
        Vote_res.objects.create(user=user, vote=vote)
        serializer = self.get_serializer(vote)
        return Response(serializer.data)

    def destroy(self, request, *args, **kwargs):
        vote = self.get_object()

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
                    user.delete()

        Vote_res.objects.filter(vote=vote).delete()
        self.perform_destroy(vote)

        return Response(
            {"detail": "Vote is deleted and are results applied"},
            status=status.HTTP_204_NO_CONTENT,
        )
