from django.shortcuts import get_object_or_404, redirect
from rest_framework import serializers,status
from rest_framework.response import Response
from rest_framework.views import APIView
from portofillo.models import About, GoodAt, Profile, Projects
from .serializers import AboutItemSerializers, GoodAtItemSerializers, ProfileItemSerializers, ProjectItemSerializers


class ProfileREST(APIView):

    def post(self,request):
        serializer = ProfileItemSerializers(data=request.data)
        if serializer.is_valid():
            serializer.save()
            context = {
                "status":"success",
                "data":serializer.data
            }

            return Response(context,status=status.HTTP_200_OK)
        else:
            context = {
                "status":"success",
                "data":serializer.errors
            }
            return Response(context,status=status.HTTP_400_BAD_REQUEST)

    def get(self,request,id=None):
        
            if id:
                item = Profile.objects.get(id=id)
                serializer = ProfileItemSerializers(item)
                context = {
                    "status":"success",
                    "data":serializer.data
                }
                return Response(context,status=status.HTTP_200_OK)
        
        
            item = Profile.objects.all()
            serializer = ProfileItemSerializers(item,many=True)
            context = {
                "status":"success",
                "data":serializer.data
            }
            return Response(context,status=status.HTTP_200_OK)
        

    def patch(self,request,id=None):
        item = Profile.objects.get(id=id)
        serializer = ProfileItemSerializers(item,data=request.data,partial=True)
        if serializer.is_valid():
            context = {
                "status":"succes",
                "data":serializer.data
            }
            return Response(context,status=status.HTTP_200_OK)

        else:
            return Response({"status":"error","data":serializer.errors})

    def delete(self,request,id=None):
        #objeyi getir veya 404 dön ilgili modeldeki id.si eşit olanları
        item = get_object_or_404(Profile,id=id)
        #sil
        item.delete()
        context = {
                "status":"success",
                "data":"Kayıt Silindi"
            }
        #geri dönüş
        return Response(context)


class AboutREST(APIView):

    def get(self,request,id=None):
        
            if id:
                item = About.objects.get(id=id)
                serializer = AboutItemSerializers(item)
                context = {
                    "status":"success",
                    "data":serializer.data
                }
                return Response(context,status=status.HTTP_200_OK)
        
        
            item = About.objects.all()
            serializer = AboutItemSerializers(item,many=True)
            context = {
                "status":"success",
                "data":serializer.data
            }
            return Response(context,status=status.HTTP_200_OK)

class GoodAtREST(APIView):
    def get(self,request,id=None):
        
            if id:
                item = GoodAt.objects.get(id=id)
                serializer = GoodAtItemSerializers(item)
                context = {
                    "status":"success",
                    "data":serializer.data
                }
                return Response(context,status=status.HTTP_200_OK)
        
        
            item = GoodAt.objects.all()
            serializer = GoodAtItemSerializers(item,many=True)
            context = {
                "status":"success",
                "data":serializer.data
            }
            return Response(context,status=status.HTTP_200_OK)

class ProjectREST(APIView):
    def get(self,request,id=None):
        
            if id:
                item = Projects.objects.get(id=id)
                serializer = ProjectItemSerializers(item)
                context = {
                    "status":"success",
                    "data":serializer.data
                }
                return Response(context,status=status.HTTP_200_OK)
        
        
            item = Projects.objects.all()
            serializer = ProjectItemSerializers(item,many=True)
            context = {
                "status":"success",
                "data":serializer.data
            }
            return Response(context,status=status.HTTP_200_OK)




