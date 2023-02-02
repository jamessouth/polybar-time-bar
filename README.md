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
This repo is a progress bar module that shows how much of the day (or other time period) has passed. It uses the [Unicode code points 2588-F](https://www.unicode.org/charts/PDF/U2580.pdf) (1/8 block to full block) but you can use anything, including [Powerline characters](https://github.com/ryanoasis/powerline-extra-symbols#glyphs). This [Stack Overflow answer](https://stackoverflow.com/a/68298090) was helpful in developing this module.

<p>&nbsp;</p>



## Installation
Just copy and paste the examples below into a file like your other polybar scripts. Make it executable with `chmod +x filename.sh`. 

## Usage
In your module file:
```
type = custom/script
exec = IFS=\\. read -a flds <<<$(awk 'BEGIN{split(strftime("%T"),a,":");len=135;f=(a[1]/24+a[2]/1440+a[3]/86400)*len;printf "%.6f.%d", f, len}'); bash ~/.config/polybar/timebarscript.sh ${flds[0]} ${flds[1]} ${flds[2]}
interval = 80
format = <label>
;format-foreground = ${colors.red}
label = %{T1}%output%
```
The `exec` command uses `awk` and `read` to calculate the number of whole blocks (█)
needed and what fraction of the next one to show. The script is then called with these values and the length in characters you use to display this module. Note that you set this (len) in the `awk` command. In this example I am using 135 characters, which at font size 13, covers nearly the whole width of my display and allows for an integer interval (80). The more characters you use, the more often it will update. Since the overhead is so low, I use a second polybar just for this. 

To determine the optimal interval divide 86400 (number of seconds in a day) by the product of 8 and how many characters you use to display the time bar - `86400/(module length in characters*8)`. The following table lists character lengths that work out to succinct interval times. These pairs of figures are reversible; 100 characters will update every 108 seconds and 108 characters every 100 seconds.

|Characters|Interval|Characters|Interval|
|:-:|:-:|:-:|:-:|
|10|1080|40|270|
|12|900|45|240|
|15|720|48|225|
|16|675|50|216|
|18|600|54|200|
|20|540|60|180|
|24|450|72|150|
|25|432|75|144|
|27|400|80|135|
|30|360|90|120|
|36|300|100|108|

<p>&nbsp;</p>

## Examples
### Default bar, single color

```bash
#!/usr/bin/env bash
EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do WHOLE+=█; done
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```
<img alt="solid color" src="mono.jpg">

The color can be set by your polybar config in the normal way (`format-foreground`).

### Multi-color

```bash
#!/usr/bin/env bash
EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do printf -v C '%%{F#a4%04x}' $(($i*49));WHOLE+=${C}█; done
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```
<img alt="time over gradient" src="algo.jpg">

This one adds a polybar [format tag](https://github.com/polybar/polybar/wiki/Formatting#foreground-color-f) before each block. The `printf` command takes the loop variable (`$i`) multiplied by an arbitrary value (`49`) and injects it as a hex value to create a color in the format tag, stored in the variable `C`. `C` is then concatenated to `WHOLE` before the block, setting its color. The format tag works like this:
| | |
|:-:|:-:|
|%%|a literal %|
|{F#|open brace, F for foreground, # for hex color|
|a4|arbitrary red value|
|%04x|printf formatting for 0-padded hex value, width 4|
|}|close brace|

The width of the hex value must be enough to accomodate whatever number you want to put there to create a color algorithmically. 




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
