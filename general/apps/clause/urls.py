"""URLS for clause ai agent"""

from django.urls import path

from views import *

urlpatterns = [
    path("clause/", ClauseView.as_view(), name="clause"),
]
