from django.urls import path
from . import views


urlpatterns = [
    path("",views.index,name="index"),
    path("about/",views.about,name="about"),
    path("goodat/",views.good_at,name="goodat"),
    path("mywork/",views.my_work,name="mywork"),


]
