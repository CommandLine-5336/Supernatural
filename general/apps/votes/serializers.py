from rest_framework import serializers

from .models import Vote, Vote_res


class VoteSerializer(serializers.ModelSerializer):
    user_alias = serializers.ReadOnlyField(source="user.alias")

    class Meta:
        model = Vote
        fields = "__all__"


class VoteResSerializer(serializers.ModelSerializer):
    class Meta:
        model = Vote_res
        fields = "__all__"
