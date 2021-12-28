from django.db import models

class Profile(models.Model):
    name = models.CharField(max_length=15)
    surname = models.CharField(max_length=25)
    description = models.TextField(max_length=255)
    image = models.ImageField(upload_to="portofillo",blank=True,null=True)
    linkedin = models.URLField(blank=True,null=True)
    facebook = models.URLField(blank=True,null=True)
    twitter = models.URLField(blank=True,null=True)
    github = models.URLField(blank=True,null=True)

    def __str__(self) -> str:
        return self.name

class About(models.Model):
    titleL = models.CharField(max_length=25,blank=True,null=True)
    descriptionL = models.TextField(max_length=255,blank=True,null=True)
    imageL = models.ImageField(upload_to="portofillo",blank=True,null=True)
    titleR = models.CharField(max_length=25,blank=True,null=True)
    descriptionR = models.TextField(max_length=255,blank=True,null=True)
    imageR = models.ImageField(upload_to="portofillo",blank=True,null=True)

    def __str__(self):
        self.title = "Burası Anasayfa Hakkımda"
        return self.title



class GoodAt(models.Model):
    title = models.CharField(max_length=15,blank=True,null=True)
    description = models.TextField(max_length=130,blank=True,null=True)

    def __str__(self) -> str:
        return self.title


class Projects(models.Model):
    title = models.CharField(max_length=50)
    description = models.TextField(max_length=255)
    image = models.ImageField(upload_to="portofillo")

    def __str__(self) -> str:
        return self.title


class Contact(models.Model):
    name = models.CharField(max_length=100)
    email =models.EmailField()
    subject = models.CharField(max_length=244)
    message = models.TextField(blank=True,null=True)


    def __str__(self):
        return self.email