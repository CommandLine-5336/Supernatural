"""Models for general"""

from django.db import models

# pylint: disable=too-few-public-methods


class User(models.Model):
    """Users table"""

    alias = models.CharField(max_length=100)
    email = models.EmailField(unique=True)
    password = models.CharField(max_length=300)
    status = models.CharField(max_length=100)

    class Meta:
        """Meta for Users"""

        db_table = "users"
        verbose_name = "User"
        app_label = "myapp"

    def __str__(self) -> str:
        return f"{self.alias}"


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
        app_label = "myapp"

    def __str__(self) -> str:
        return f"{self.name}"


class Report(models.Model):
    """Reports table"""

    user = models.ForeignKey(User, on_delete=models.SET_NULL, null=True, blank=True)
    post = models.ForeignKey(Post, on_delete=models.SET_NULL, null=True, blank=True)

    class Meta:
        """Meta for Reports"""

        db_table = "reports"
        verbose_name = "Report"
        app_label = "myapp"

    def __str__(self) -> str:
        return f"{self.user_id} {self.post_id}"


class Federal(models.Model):
    """Federals table"""

    ip_address = models.GenericIPAddressField()

    class Meta:
        """Meta for Federals"""

        db_table = "federals"
        verbose_name = "Federal"
        app_label = "myapp"

    def __str__(self) -> str:
        return f"{self.ip_address}"
