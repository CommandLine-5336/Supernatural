from rest_framework import viewsets, status
from rest_framework.decorators import action
from rest_framework.response import Response
from .models import Post
from .serializers import PostSerializer
from .storage import upload_image

class PostViewSet(viewsets.ModelViewSet):
    queryset = Post.objects.all().order_by("-created_at")
    serializer_class = PostSerializer

    def create(self, request, *args, **kwargs):
        data = request.data.copy()
        image = request.FILES.get("image")
        serializer = self.get_serializer(data=data)
        serializer.is_valid(raise_exception=True)
        image_url = upload_image(image) if image else None
        serializer.save(image_url=image_url)
        headers = self.get_success_headers(serializer.data)
        return Response(serializer.data, status=status.HTTP_201_CREATED, headers=headers)

    @action(detail=True, methods=["post"])
    def seen(self, request, pk=None):
        post = self.get_object()
        post.seen_count += 1
        post.save(update_fields=["seen_count"])
        return Response({"seen_count": post.seen_count}, status=status.HTTP_200_OK)