import os

import jwt
from django.test import TestCase

from ..authentication.models import User
from .models import Vote, Vote_res

JWT_SECRET = os.getenv("JWT_KEY")


class VoteTestCase(TestCase):
    def setUp(self):
        """Creating new post"""
        self.Ben = User.objects.create(
            alias="Ben",
            email="he@gmail.com",
            password="1234",
            status="silver",
            inquisitor=False,
        )
        self.Donna = User.objects.create(
            alias="Donna",
            email="she@gmail.com",
            password="1234",
            status="copper",
            inquisitor=False,
        )

        payload = {
            "sub": str(self.Ben.id),
            "status": self.Ben.status,
        }

        self.token = jwt.encode(payload, JWT_SECRET, algorithm="HS256")

        if isinstance(self.token, bytes):
            self.token = self.token.decode("utf-8")

        self.client.cookies["jwt"] = self.token

        self.Promo = Vote.objects.create(
            type="promotion",
            description="promoting",
            user=self.Ben,
            agree=2,
            disagree=0,
        )
        self.Ex = Vote.objects.create(
            type="excommunication",
            description="executing",
            user=self.Donna,
            agree=1,
            disagree=0,
        )

    def test_set_vote(self):
        """Setting agree vote"""
        url = f"/api/votes/{self.Ex.id}/set_vote/"
        data = {"res": "+"}
        response = self.client.post(url, data, content_type="application/json")
        self.assertEqual(response.status_code, 200)
        self.Ex.refresh_from_db()
        self.assertEqual(self.Ex.agree, 2)
        self.assertTrue(Vote_res.objects.filter(user=self.Ben, vote=self.Ex).exists())

    def test_delete_ex_vote(self):
        """Delete excommunication vote"""
        url = f"/api/votes/{self.Ex.id}/"
        response = self.client.delete(url)
        self.assertEqual(response.status_code, 204)
        self.assertFalse(Vote.objects.filter(id=self.Ex.id).exists())
        self.assertFalse(User.objects.filter(id=self.Donna.id).exists())

    def test_delete_promo_vote(self):
        """Delete promo vote"""
        url = f"/api/votes/{self.Promo.id}/"
        response = self.client.delete(url)
        self.assertEqual(response.status_code, 204)
        self.Ben.refresh_from_db()
        self.assertFalse(Vote.objects.filter(id=self.Promo.id).exists())
        self.assertEqual(self.Ben.status, "gold")
