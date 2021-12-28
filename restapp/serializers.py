from rest_framework import serializers
from portofillo.models import About, GoodAt, Profile, Projects


class ProfileItemSerializers(serializers.ModelSerializer):
    name = serializers.CharField(max_length=15)
    surname = serializers.CharField(max_length=25)
    description = serializers.CharField(style={'base_template': 'textarea.html'})
    image = serializers.ImageField()
    linkedin = serializers.URLField()
    facebook = serializers.URLField()
    twitter = serializers.URLField()
    github = serializers.URLField()

    class Meta:
        model = Profile
        fields = ('__all__')


class AboutItemSerializers(serializers.ModelSerializer):
    titleL = serializers.CharField(max_length=25)
    descriptionL = serializers.CharField(style={'base_template': 'textarea.html'})
    titleR = serializers.CharField(max_length=25)
    descriptionR = serializers.CharField(style={'base_template': 'textarea.html'})

    class Meta:
        model = About
        fields = ['titleL','descriptionL','titleR','descriptionR']


class GoodAtItemSerializers(serializers.ModelSerializer):
    title = serializers.CharField(max_length=50)
    description = serializers.CharField(style={'base_template': 'textarea.html'})
    
    class Meta:
        model = GoodAt
        fields = ('__all__')


class ProjectItemSerializers(serializers.ModelSerializer):
    title = serializers.CharField(max_length=50)
    description = serializers.CharField(style={'base_template': 'textarea.html'})

    class Meta:
        model = Projects
        fields = ['id','title','description']


