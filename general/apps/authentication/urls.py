"""URL routes for votes"""

from django.urls import include, path
from rest_framework.routers import DefaultRouter

from . import views

urlpatterns = [
    path("change_status/", views.change_status, name="change_status"),
    path("architectors/", views.architectors, name="architectors"),
]
