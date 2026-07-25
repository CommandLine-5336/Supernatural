import json
import uuid
import boto3
from django.conf import settings


def _get_s3_client():
    return boto3.client(
        "s3",
        endpoint_url=settings.AWS_S3_ENDPOINT_URL,
        aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
        aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
    )


def ensure_public_bucket():
    s3 = _get_s3_client()
    bucket = settings.AWS_STORAGE_BUCKET_NAME

    existing = [b["Name"] for b in s3.list_buckets().get("Buckets", [])]
    if bucket not in existing:
        s3.create_bucket(Bucket=bucket)

    policy = {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Allow",
                "Principal": "*",
                "Action": "s3:GetObject",
                "Resource": f"arn:aws:s3:::{bucket}/*",
            }
        ],
    }
    s3.put_bucket_policy(Bucket=bucket, Policy=json.dumps(policy))


def upload_image(file) -> str:
    s3 = _get_s3_client()
    ensure_public_bucket()
    key = f"posts/{uuid.uuid4()}-{file.name}"
    s3.upload_fileobj(
        file,
        settings.AWS_STORAGE_BUCKET_NAME,
        key,
        ExtraArgs={"ContentType": file.content_type},
    )
    return f"{settings.AWS_S3_PUBLIC_URL}/{settings.AWS_STORAGE_BUCKET_NAME}/{key}"
