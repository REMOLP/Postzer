### Postzer aka making static pages faster and as easy as possible.
You should guess the usage easily by yourself just by looking at files in the repo, but here's a quick explanation:
- ```aliases.pz``` is for keeping the actual shorthands for full HTML for all the aliases, separated by a space. The alias can contain n amount of ```{{indexes}}``` starting from zero that will be replaced in the final output HTML file
- ```myfirstpost.pzp``` is your post as the name suggest. Name it however you want
- ```init-template.html``` is your initial state of the HTML page that will consist all the default tags etc. for example for styling and other boring stuff. {{replace:me:here}} is the main custom text that the template file should always keep. init-template.html will be parsed by Postzer by default, but if you want, you can specify optional last argument to your own template.html file (name it as you want of course)

Static binaries under "Releases" tab.

Happy static posts making ^^
