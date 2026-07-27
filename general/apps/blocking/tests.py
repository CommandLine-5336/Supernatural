import json

from django.test import TestCase

from .models import Banned


class TestReportip(TestCase):
    def setUp(self):
        self.url = "/api/blocking/report/"

    def test_get_reject(self):
        response = self.client.get(self.url)
        self.assertEqual(response.status_code, 405)

    def test_noip_provided(self):
        response = self.client.post(
            self.url, data=json.dumps({}), content_type="application/json"
        )
        self.assertEqual(response.status_code, 400)

    def test_ip_sent_as_int(self):
        response = self.client.post(
            self.url,
            data=json.dumps({"ip_address": 123}),
            content_type="application/json",
        )
        self.assertEqual(response.status_code, 400)

    def test_ipv6_or_incorrectip(self):
        response = self.client.post(
            self.url,
            data=json.dumps({"ip_address": "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}),
            content_type="application/json",
        )
        self.assertEqual(response.status_code, 400)

    def test_empty_ip_provided(self):
        response = self.client.post(
            self.url,
            data=json.dumps({"ip_address": ""}),
            content_type="application/json",
        )
        self.assertEqual(response.status_code, 400)

    def test_ip_exists(self):
        Banned.objects.create(ip_address="1.2.3.4")
        response = self.client.post(
            self.url,
            data=json.dumps({"ip_address": "1.2.3.4"}),
            content_type="application/json",
        )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(Banned.objects.filter(ip_address="1.2.3.4").count(), 1)
        Banned.objects.filter(ip_address="1.2.3.4").delete()

    def test_someone_banned_successfully(self):
        response = self.client.post(
            self.url,
            data=json.dumps({"ip_address": "1.2.3.4"}),
            content_type="application/json",
        )
        self.assertEqual(response.status_code, 200)
        self.assertTrue(Banned.objects.filter(ip_address="1.2.3.4").exists())
