"""S3 storage for post images"""

import uuid
import boto3
from django.conf import settings


def _get_s3_client():
    """Returns a boto3 S3 client configured for AWS"""
    return boto3.client(
        "s3",
        region_name=settings.AWS_REGION,
        aws_access_key_id=settings.AWS_ACCESS_KEY_ID,
        aws_secret_access_key=settings.AWS_SECRET_ACCESS_KEY,
    )


def upload_image(file) -> str:
    """Uploads an image and returns its public URL"""
    s3 = _get_s3_client()
    key = f"posts/{uuid.uuid4()}-{file.name}"
    s3.upload_fileobj(
        file,
        settings.AWS_STORAGE_BUCKET_NAME,
        key,
        ExtraArgs={"ContentType": file.content_type},
    )
    bucket = settings.AWS_STORAGE_BUCKET_NAME
    region = settings.AWS_REGION
    return f"https://{bucket}.s3.{region}.amazonaws.com/{key}"
