from datetime import timedelta

from django.utils import timezone

from .models import Vote
from .views import execute_vote


def delete_votes():
    time = timezone.now() - timedelta(days=1)
    expired_votes = Vote.objects.filter(time_created__lt=time)
    for vote in expired_votes:
        try:
            execute_vote(vote)
        except Exception as e:
            print("Couldn't delete the vote: ", e)
