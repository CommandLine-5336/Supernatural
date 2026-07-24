import os

from django.apps import AppConfig


class VotesConfig(AppConfig):
    name = "apps.votes"

    def ready(self):
        if os.environ.get("RUN_MAIN") != "true":
            return
        print("SCHEDULEEEEE", flush=True)
        from .scheduler.scheduler import start

        start()
