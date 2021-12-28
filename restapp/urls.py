from django.urls import path
from restapp.views import AboutREST, GoodAtREST, ProfileREST, ProjectREST

urlpatterns = [
    path("profile-api/",ProfileREST.as_view(),name="profile-api"),
    path("profile-api/<int:id>",ProfileREST.as_view()),
    path("profile-api/about",AboutREST.as_view()),
    path("profile-api/goodat",GoodAtREST.as_view()),
    path("profile-api/projects",ProjectREST.as_view())
]
