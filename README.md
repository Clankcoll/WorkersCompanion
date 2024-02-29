```[tasklist]
### My tasks
- [ ] test
- [ ] test2
```
Please keep in mind the Website is not yet always up to date until i know how i could automate this to build its own Docker image rn i need to do it by hand.

Also this Project and looks of it may and probably will change alot until i am statisfied for a personal full release

As you can see in the main.go file there is no handler for https and no cert set this is because in my setup i am running this site behind a reverse proxy which handles the Certs and HTTPS connections.
I should als add the information or a Disclaimer yes internally if u use it like a internal tool there is no Security and encryption as it only works with http cause we are not saving or serving personal data,
this should be fine for me atleast. 
