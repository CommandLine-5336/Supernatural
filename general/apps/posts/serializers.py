"""Serializers for posts"""

from rest_framework import serializers
from .models import Post, Report

# pylint: disable=too-few-public-methods


class PostSerializer(serializers.ModelSerializer):
    """Serializer for Post model"""

    already_seen = serializers.SerializerMethodField()

    class Meta:
        """Meta options for PostSerializer"""

        model = Post
        fields = [
            "id",
            "name",
            "description",
            "latitude",
            "longitude",
            "image_url",
            "visibility_level",
            "seen_count",
            "created_at",
            "already_seen",
        ]
        read_only_fields = [
            "id",
            "image_url",
            "seen_count",
            "created_at",
            "already_seen",
        ]

    def get_already_seen(self, obj):
        """Check whether the current authenticated user already reported seeing this post"""
        request = self.context.get("request")
        user = getattr(request, "user", None)
        if not user or not getattr(user, "is_authenticated", False):
            return False
        return Report.objects.filter(user=user, post=obj).exists()


class ReportSerializer(serializers.ModelSerializer):
    """Serializer for Report model"""

    class Meta:
        """Meta options for ReportSerializer"""

        model = Report
        fields = ["id", "user", "post"]
