# String 2 Png in GoLang [![Maintenance](https://img.shields.io/badge/Maintained%3F-no-red.svg)](https://GitHub.com/wsdt/GoLang_StringToPng/graphs/commit-activity) [![Generic badge](https://img.shields.io/badge/In-GO-BLUE.svg)](https://golang.org/) [![GitHub license](https://img.shields.io/github/license/wsdt/GoLang_StringToPng.svg)](https://github.com/wsdt/GoLang_StringToPng/blob/master/LICENSE)
First of all this is my first Go project, so please bare with me. 

## What is it? [![Open Source Love svg2](https://badges.frapsoft.com/os/v2/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)]
This program converts arbitrary string into a colorful .PNG. 

Example pictures: [pic1](https://github.com/wsdt/GoLang_StringToPng/blob/master/image_0.png), [pic2](https://github.com/wsdt/GoLang_StringToPng/blob/master/image_1.png), 
[pic3](https://github.com/wsdt/GoLang_StringToPng/blob/master/image_2.png)

![alt text](https://github.com/wsdt/GoLang_StringToPng/blob/master/image_0.png "Example image") 
![alt text](https://github.com/wsdt/GoLang_StringToPng/blob/master/image_1.png "Example image") 
![alt text](https://github.com/wsdt/GoLang_StringToPng/blob/master/image_2.png "Example image")


E.g. a random string like "5dfs45654sfefsfdfdsfew09´" or even a useful information like "Some secret geeky message." can be encoded into a picture. 

## How does it work?
The colors itself do not contain any information about the inputted string and are completely randomized. 

Therefore, each char of the string/information will be converted to unicode (numeric representation -> e.g. 'H' = 72, ...). This numeric representation is going to be used as an indicator of how long a color is supposed to stay. 

Consequently, a char like 'H' with the representation of 72 will cause 72 pixels in a random color. When the message is too short for the image-size the information will be repeated (so the image might contain the information multiple times). 

As the numeric representation of each char is saved into the image, the image could be theoretically converted back to a string. 

## Some thoughts
As images, videos, sound-files etc. consist of several chars (e.g. if you open one of those files via a text-editor) you could also encode a video/sound-file as such a image (for sure, the image would need to have the right size). 

But yeah, this whole project is just supposed to be something funny or whatever. To be honest I just wanted to use go-routines and the image-library :)
