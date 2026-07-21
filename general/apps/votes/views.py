from rest_framework import filters, status, viewsets
from rest_framework.decorators import action
from rest_framework.permissions import IsAuthenticated
from rest_framework.response import Response

from ..authentication.models import User
from .models import Vote
from .serializers import VoteSerializer


class VoteViewSet(viewsets.ModelViewSet):
    http_method_names = ["get", "post", "put", "patch", "delete"]
    serializer_class = VoteSerializer
    permission_classes = (IsAuthenticated,)
    filter_backends = [filters.OrderingFilter]
    ordering_fields = ["id", "agree", "disagree"]
    ordering = ["-id"]

    def get_queryset(self):
        return Vote.objects.all()

    def perform_create(self, serializer):
        serializer.save(user=self.request.user)

    @action(detail=True, methods=["post", "put"])
    def set_vote(self, request, pk=None):
        vote = self.get_object()
        res = request.data.get("res")
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
        serializer = self.get_serializer(vote)
        return Response(serializer.data)

    def destroy(self, request, *args, **kwargs):
        vote = self.get_object()

        user = vote.user
        total_users = User.objects.count()
        if user is not None:
            if vote.agree > total_users / 2:
                if vote.type == "promotion":
                    if user.status == "copper":
                        user.status = "silver"
                    else:
                        user.status = "gold"
                    user.save(update_fields=["status"])

                elif vote.type == "excommunication":
                    user.delete()

        self.perform_destroy(vote)

        return Response(
            {"detail": "Vote deleted and resolution applied successfully"},
            status=status.HTTP_204_NO_CONTENT,
        )
