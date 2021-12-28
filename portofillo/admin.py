from django.contrib import admin

from portofillo.forms import ContactForm

from .models import About, Contact,GoodAt, Profile, Projects

admin.site.register(About)
admin.site.register(GoodAt)
admin.site.register(Projects)
admin.site.register(Profile)
admin.site.register(Contact)
