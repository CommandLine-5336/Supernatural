"""Serializers for votes"""

from rest_framework import serializers

from .models import Vote, VoteRes


class VoteSerializer(serializers.ModelSerializer):
    """Serializer for Vote model"""

    user_alias = serializers.ReadOnlyField(source="user.alias")

    class Meta:
        model = Vote
        fields = "__all__"


class VoteResSerializer(serializers.ModelSerializer):
    """Serializer for VoteRes model"""

    class Meta:
        model = VoteRes
        fields = "__all__"
