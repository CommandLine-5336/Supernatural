"""Models for votes"""

from django.db import models

# pylint: disable=too-few-public-methods

from ..authentification.models import User

class Vote(models.Model):
    """Votes table"""

    type = models.CharField(max_length=100)
    description = models.CharField(max_length=300)
    user = models.ForeignKey(User, on_delete=models.SET_NULL, null=True, blank=True)
    agree = models.IntegerField()
    disagree = models.IntegerField()

    class Meta:
        """Meta for Votes"""

        db_table = "votes"
        verbose_name = "Vote"

    def __str__(self) -> str:
        return f"{self.user_id} {self.type}"