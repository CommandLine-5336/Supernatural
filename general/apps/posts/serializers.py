"""Serializers for posts"""
from rest_framework import serializers
from .models import Post


class PostSerializer(serializers.ModelSerializer):
    """Serializer for Post model"""

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
        ]
        read_only_fields = ["id", "image_url", "seen_count", "created_at"]