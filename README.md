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
	<a href="https://www.gnu.org/software/bash/manual/"><img src="https://img.shields.io/badge/Bash-d.svg?logoWidth=40&labelColor=4eaa25&color=293137&logoColor=ffffff&logo=GNU%20Bash"></a>
	<img src="https://img.shields.io/badge/awesome-%C6%94%F0%9D%9A%BA%C5%9E-235789.svg">
</p>

<p>&nbsp;</p>

## Description
This repo is a customizable script that provides a progress bar for the day or other time period. It calculates what percentage of the day has passed and applies that to the area you reserve for it, thereby showing the day's progress. It uses the [Unicode code points 2588-F](https://www.unicode.org/charts/PDF/U2580.pdf) (1/8 block to full block) but you can use anything, including [Powerline characters](https://github.com/ryanoasis/powerline-extra-symbols#glyphs). This [Stack Overflow answer](https://stackoverflow.com/a/68298090) was very helpful in developing this module.

<p>&nbsp;</p>



## Installation
To install a particular font:
```bash
sudo mkdir -pv /usr/share/fonts/BinaryClock && cd /usr/share/fonts/BinaryClock && sudo curl -JOL https://github.com/jamessouth/polybar-binary-clock-fonts/blob/master/BinaryClockLigatureMono.ttf?raw=true && sudo fc-cache -fv && fc-list | awk '/BinaryClock/ {print $2}'
```
This will create a directory for the Binary Clock font family in `/usr/share/fonts/`, change to it, download the font file, update your font cache, and output the name and style. Replace the filename as needed.



<p>&nbsp;</p>

## Usage
A basic bar with a single color. The color can be set by your polybar config in the normal way (`format-foreground`).

```bash
EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do WHOLE+=█; done
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```
<img alt="solid color" src="mono.jpg">
It is called like this in the module config:

```
exec = IFS=\\. read -a flds <<<$(awk 'BEGIN{split(strftime("%T"),a,":");f=(a[1]/24+a[2]/1440+a[3]/86400)*135;printf "%.6f", f}'); bash ~/.config/polybar/gtmscript.sh ${flds[0]} ${flds[1]} 135
```
Note that the length I am using, 135, appears twice in this command. This will set your `IFS` to period, get the time with `awk`, calculate the portion of the day that has passed, read the number of whole blocks and part of the next block that need to be shown, then call the script with those numbers and the module length. You don't need `tail = true` unless you want to make the module clickable to change the state as described below. To determine the interval for calling the script, use this simple command: `bc -l <<< '86400/(module length in characters*8)'`. This is the number of seconds in a day divided by the product of the number of characters (not pixels) you want to use to display the bar and 8. I use a separate polybar just for this module so it can use the whole width of my display. 135 characters covers almost the whole screen and allows for an interval of 80 seconds. Below I will show how to rotate colors in which case you will want to set a lower interval. This just means the bar will not extend but the colors can still change.


|Characters|Interval|
|:-:|:-:|
|10|1080|
|12|900|
|15|720|
|16|675|
|18|600|
|20|540|
|24|450|
|25|432|
|27|400|
|30|360|
|32|337.5|
|36|300|
|40|270|
|45|240|
|48|225|
|50|216|
|54|200|
|60|180|
|72|150|
|75|144|
|80|135|
|90|120|
|96|112.5|
|100|108|
|108|100|
|120|90|
|125|86.4|
|135|80|
|144|75|
|150|72|
|160|67.5|
|180|60|
|192|56.25|
|200|54|
|216|50|
|225|48|
|240|45|
|250|43.2|
|270|40|
|288|37.5|
|300|36|
|320|33.75|
|360|30|
|375|28.8|
|400|27|
|432|25|
|450|24|
|480|22.5|
|500|21.6|
|540|20|
|576|18.75|
|600|18|
|625|17.28|
|675|16|



```bash
EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do printf -v C '%%{F#a4%04x}' $(($i*49));WHOLE+=${C}█; done
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```
<img alt="time over gradient" src="algo.jpg">




You can send a signal to the script as [shown here](https://github.com/polybar/polybar/wiki/Module:-script#examples) to change the display. You can also just create a separate transparent bar and place it over a bar with this module.


<img alt="time over gradient" src="over.jpg">


```bash
EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do WHOLE+=${RAINB[$i]}█; done
WHOLE+=${RAINB[$1]}
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```
<img alt="gradient" src="rnbw.jpg">




```bash
SPACES=$(($2-$1))
for ((i=0; i<$1; i++)); do
if [[ $((($i+1) % 5)) -eq 0 ]];
then WHOLE+=%{B${PWR[($i/5)+1]}}
  else WHOLE+=%{F${PWR[$i/5]}}█;
fi
 done
WHOLE+=%{B-}
printf '%s%*s' "$WHOLE" $SPACES ''
```
with this command:
```
exec = IFS=\\. read -a flds <<<$(awk 'BEGIN{split(strftime("%T"),a,":");f=(a[1]/24+a[2]/1440+a[3]/86400)*120;printf "%.6f", f}'); bash ~/.config/polybar/gtmscript.sh ${flds[0]} 120
```
<img alt="powerline" src="pwrln.jpg">



```bash
LEN=${#RAINB[@]}
G=$(($(date +%s) % $3))
EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do WHOLE+=${RAIN2[($i+$G) % $LEN]}█; done
WHOLE+=${RAIN2[($1+$G) % $LEN]}
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```


<img alt="moving colors" src="loop.gif">

<p>&nbsp;</p>
