from django.shortcuts import redirect, render
from django.contrib import messages
from portofillo.models import About, GoodAt, Profile, Projects
from .forms import ContactForm

def profile(request):
    profile = Profile.objects.all()
    context = {
        "profiles":profile
    }
    return render(request,"partials/navbar.html",context)

def index(request):

    if request.method == 'POST':
        form = ContactForm(request.POST)
        if form.is_valid():
            form.save()
            messages.success(request,"Mesajınız Başarılı Bir Şekilde Gönderildi!")
            return redirect("index")
        else:
            form = ContactForm()
            messages.error(request,"Bilgileri Eksik Yada Yanlış Girdiniz!")
    
    form = ContactForm()
    about = About.objects.all()
    goodat= GoodAt.objects.all()
    projects = Projects.objects.all()
    profile = Profile.objects.all()
    form = ContactForm()



    context = {
        "abouts":about,
        "goodats":goodat,
        "projects":projects,
        "profiles":profile,
        "form":form

    }
    return render(request,"blog/index.html",context)


def about(request):
    about = About.objects.all()
    context = {
        "abouts":about
    }
    return render(request,"blog/about.html",context)


def my_work(request):
    project = Projects.objects.all()
    context = {
        "projects":project
    }
    return render(request,"blog/mywork.html",context)


def good_at(request):
    goodat=GoodAt.objects.all()
    context = {
        "goodats":goodat
    }
    return render(request,"blog/good_at.html",context)

def handler404(request,exception):
    data = {}
    return render(request,"err/404.html",data)
