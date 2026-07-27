"""Scheduler for automatic inquisitor estting up"""

import random
import sys

from apps.authentication.models import User
from apscheduler.schedulers.background import BackgroundScheduler


def set_inquisitor():
    """Find new inquisitor"""
    prev_inq = User.objects.filter(inquisitor=True).first()
    id_list = list(User.objects.values_list("id", flat=True))
    if prev_inq:
        prev_inq.inquisitor = False
        prev_inq.save(update_fields=["inquisitor"])
        id_list.remove(prev_inq.id)

    if not id_list:
        print("There are no candidates for inquisitor role")
        return

    new_inq_id = random.choice(id_list)
    new_inq = User.objects.get(id=new_inq_id)
    new_inq.inquisitor = True
    new_inq.save(update_fields=["inquisitor"])
    print(f"New inquisitor - {new_inq.alias}")


scheduler = BackgroundScheduler()


def start():
    """Start the scheduler"""
    if not scheduler.running:
        scheduler.add_job(
            set_inquisitor,
            "interval",
            hours=24,
            name="set_inquisitor",
            replace_existing=True,
        )
        scheduler.start()
        print("Scheduler for inquisitor sterted", file=sys.stdout)
