# JetPack Project

### Introduction

The JetPack project is an attempt to create a new ecommerce environment entirely rooted in the Go Language and designed with 4 key goals in mind:

1.	Speed - Of course we want to be the fastest ecommerce on the planet
2.	Componentized - Each jetPack component will be able to be used as a standalone component so if you come from another system you can use Jetpack Pricing for example without having to migrate in a "Big Bang" fashion.  We believe that once you see the power you will never want to go back to your old style ecommerce platform.
3.	Extensibility - Once you have a JetPack component you can add-on as much as you want even your own home developped extensions.
4.	Stability - What good is all the rest of the qualities mentionned above if the thing crashes every 5 minutes :)

We also will be using Docker as a deployment strategy for JetPack which will allow you to be up and running easily


### Developper Install and Prep (5 mins or less)

1. Create a Go Workspace (JetPack/src, JetPack/pkg, JetPack/bin)
2. Clone this Git repository in the src directory of your workspace
3. Export your GOPATH to the JetPack directory you created above
4. cd to the JetPack directory created above
4. Run the unit tests
	* go test ./...
5. Run the application
	* go run
 

### Install a JetPack Component




