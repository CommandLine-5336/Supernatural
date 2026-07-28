"""Models for authentication"""

from django.db import models

# pylint: disable=too-few-public-methods


class User(models.Model):
    """Users table"""

    alias = models.CharField(max_length=100)
    email = models.EmailField(unique=True)
    password = models.CharField(max_length=300)
    status = models.CharField(max_length=100)
    inquisitor = models.BooleanField(default=False)
    banned = models.BooleanField(default=False)

    class Meta:
        """Meta for Users"""

        db_table = "users"
        verbose_name = "User"

    def __str__(self) -> str:
        return f"{self.alias}"
