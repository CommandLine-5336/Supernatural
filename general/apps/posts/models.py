"""Models for posts"""

from django.db import models

# pylint: disable=too-few-public-methods

from ..authentification.models import User

class Post(models.Model):
    """Posts table"""

    name = models.CharField(max_length=100)
    description = models.CharField(max_length=300)
    latitude = models.CharField(max_length=100)
    longitude = models.CharField(max_length=100)

    class Meta:
        """Meta for Posts"""

        db_table = "posts"
        verbose_name = "Post"

    def __str__(self) -> str:
        return f"{self.name}"


class Report(models.Model):
    """Reports table"""

    user = models.ForeignKey(User, on_delete=models.CASCADE, null=True, blank=True)
    post = models.ForeignKey(Post, on_delete=models.CASCADE, null=True, blank=True)

    class Meta:
        """Meta for Reports"""

        db_table = "reports"
        verbose_name = "Report"

    def __str__(self) -> str:
        return f"{self.user_id} {self.post_id}"
