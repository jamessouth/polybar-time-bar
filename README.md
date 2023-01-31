<h1 align="center">polybar-time-bar</h1>

<p>&nbsp;</p>

A time module for your

<div align="center">
	<picture>
 	 <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/polybar/polybar/master/doc/_static/banner-dark-mode.png">
 	 <img alt="polybar logo" src="https://raw.githubusercontent.com/polybar/polybar/master/doc/_static/banner.png">
	</picture>
</div>

<p>&nbsp;</p>

<p align="center">
	<a href="https://github.com/jamessouth/polybar-time-bar/blob/master/LICENSE"><img src="https://img.shields.io/github/license/jamessouth/polybar-time-bar"></a>
	<a href="https://archlinux.org/"><img src="https://img.shields.io/badge/Linux-d.svg?logoWidth=40&labelColor=d35e49&color=E3C567&logoColor=000000&logo=Linux"></a>
	<a href="https://www.gnu.org/software/bash/manual/"><img src="https://img.shields.io/badge/bash-d.svg?logoWidth=40&labelColor=4eaa25&color=293137&logoColor=ffffff&logo=GNU%20Bash"></a>
	<img src="https://img.shields.io/badge/awesome-%C6%94%F0%9D%9A%BA%C5%9E-red.svg">
</p>

<p>&nbsp;</p>

## Description
This repo is a collection of fonts that represent numbers in base 2 as columns of dots like on a [binary clock](https://en.wikipedia.org/wiki/Binary_clock). The only glyphs are 0-9. Other characters will display in the first font that contains them.

<p>&nbsp;</p>

## Appearance
On my 20px polybar with font size 15, vertical offset 4:
|Font Table|Running Clock|
|:-:|:-:|
|<picture><source media="(prefers-color-scheme: dark)" srcset="montagedark.jpg"><img alt="table of fonts" src="montagelight.jpg"></picture>|<img alt="binary clock" src="vid.gif">|

<p>&nbsp;</p>

## Installation
To install a particular font:
```bash
sudo mkdir -pv /usr/share/fonts/BinaryClock && cd /usr/share/fonts/BinaryClock && sudo curl -JOL https://github.com/jamessouth/polybar-binary-clock-fonts/blob/master/BinaryClockLigatureMono.ttf?raw=true && sudo fc-cache -fv && fc-list | awk '/BinaryClock/ {print $2}'
```
This will create a directory for the Binary Clock font family in `/usr/share/fonts/`, change to it, download the font file, update your font cache, and output the name and style. Replace the filename as needed.

To install all of them:
```bash
curl -JOL https://github.com/jamessouth/polybar-binary-clock-fonts/blob/master/fonts.zip?raw=true && sudo mkdir -pv /usr/share/fonts/BinaryClock && sudo unzip fonts.zip -d /usr/share/fonts/BinaryClock && sudo fc-cache -fv && fc-list | awk '/BinaryClock/ {print $2}'
```
This will download the zip file with all 10 font files and the LICENSE, create the Binary Clock directory, unzip the files into it, update your font cache, and output the names and styles.

<p>&nbsp;</p>

## Usage
Take the output of the installation command and put it in your polybar config with your other fonts: `font-3 = BinaryClock:style=BoldMono:size=15;4`. In your date module, use [time-alt/date-alt settings](https://github.com/polybar/polybar/wiki/Module:-date#basic-settings) to switch to your normal format. Use [format tags](https://github.com/polybar/polybar/wiki/Formatting#format-tags) to show time in a different color from the date.

<p>&nbsp;</p>
