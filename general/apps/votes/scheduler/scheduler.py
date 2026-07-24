import sys
from datetime import timedelta

from apscheduler.schedulers.background import BackgroundScheduler
from django.utils import timezone

from ..models import Vote
from ..views import execute_vote


def delete_votes():
    time = timezone.now() - timedelta(days=1)
    expired_votes = Vote.objects.filter(time_created__lt=time)
    count = 0
    for vote in expired_votes:
        try:
            execute_vote(vote)
            count += 1
        except Exception as e:
            print("Couldn't delete the vote: ", e)
    print(f"Deleted {count} votes")


scheduler = BackgroundScheduler()


def start():
    if not scheduler.running:
        scheduler.add_job(
            delete_votes,
            "interval",
            minutes=1,
            name="delete_votes",
            replace_existing=True,
        )
        scheduler.start()
        print("Scheduler started...", file=sys.stdout)
