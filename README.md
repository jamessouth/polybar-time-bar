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
exec = IFS=\\. read -a flds <<< $(awk 'BEGIN{split(strftime("%T"),a,":");len=135;f=(a[1]/24+a[2]/1440+a[3]/86400)*len;printf "%.6f.%d", f, len}'); bash ~/.config/polybar/timebarscript.sh ${flds[0]} ${flds[1]} ${flds[2]}
interval = 80
format = <label>
;format-foreground = ${colors.red}
label = %{T1}%output%
```
The `exec` command uses `awk` and `read` to calculate the number of whole blocks (`█`) needed and what fraction of the next one to show. The script is then called with these values and the length in characters you use to display this module. Note that you set this (len) in the `awk` command. In this example I am using 135 characters, which at font size 13, covers nearly the whole width of my display and allows for an integer interval (80). The more characters you use, the more often it will update. Since the overhead is so low, I use a second polybar just for this. 

To determine the optimal interval divide 86400 (number of seconds in a day) by the product of 8 and how many characters you use to display the time bar - `86400/(module length in characters*8)`. The following table lists lengths in characters that have integer interval times. Past 100, just flip the numbers around; e.g., a length of 150 would update every 72 seconds.

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
<p>&nbsp;</p>

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

This one adds a polybar [format tag](https://github.com/polybar/polybar/wiki/Formatting#foreground-color-f) before each block. The `printf` command takes the loop variable (`$i`) multiplied by an arbitrary value (`49`) and injects it as a hex value to create a color in the format tag, stored in the variable `C`. `C` is then concatenated to `WHOLE` before the block, setting its color. The format tag `%%{F#a4%04x}` works like this:
| | |
|:-:|:-:|
|%%|a literal %|
|{F#|open brace, F for foreground, # for hex color|
|a4|arbitrary red value|
|%04x|printf formatting for 0-padded hex value, width 4|
|}|close brace|

The width of the zero-padded hex value must be enough to accomodate whatever number you want to put there to create a color algorithmically.
<p>&nbsp;</p>

### Display date/time with this module
You can use the multiple-state functionality [shown here](https://github.com/polybar/polybar/wiki/Module:-script#examples) to send a signal to the script and change the display. You can also just create a separate transparent bar with such functionality and place it over a bar displaying this module.

<img alt="time with gradient background" src="over.jpg">
<p>&nbsp;</p>

### Custom gradient
```bash
#!/usr/bin/env bash

declare -a RAINB=(
"%{F#ff0000}"
"%{F#ff0200}"
...132 omitted...
"%{F#f800ff}"
)

EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do WHOLE+=${RAINB[$i]}█; done
WHOLE+=${RAINB[$1]}
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```
<img alt="gradient" src="rnbw.jpg">

You can create custom arrays of colors formatted as polybar [format tags](https://github.com/polybar/polybar/wiki/Formatting#format-tags). The length should be equal to the number of characters you use to display the bar. Here I declare the `RAINB` array (edited to save space) with 135 colors for my 135-character-wide time bar. In the loop, a color is applied to each full block, then the color at the `$1` index (number of full blocks) is applied to the partial block. No color is applied to the spaces, therefore your polybar's background color will appear.
<p>&nbsp;</p>

### Powerline symbols
```bash
#!/usr/bin/env bash

declare -a PWR=(
"#ff0e00"
"#ff2e00"
...21 omitted...
"#ff7100"
)

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
<img alt="powerline" src="pwrln.jpg">

To mimic powerline I made a few changes to this example. The partial block characters don't look very good with powerline symbols, so I removed that code and am not printing the `PORTION` character like in the other example. I also have only 24 colors in the `PWR` array, 1 for each hour of the day. Also note they are just hex colors; they are not yet formatted for polybar because we have to set both foreground and background colors to get the powerline effect. This script is called from the module file with this adjusted command:
```
exec = IFS=\\. read -a flds <<< $(awk 'BEGIN{split(strftime("%T"),a,":");len=120;f=(a[1]/24+a[2]/1440+a[3]/86400)*len;printf "%.6f.%d", f, len}'); bash ~/.config/polybar/timebarscript.sh ${flds[0]} ${flds[2]}
```
Note the length of 120 (multiple of 24) and the omission of `flds[1]` (the partial block) from the arguments list the script is called with. In the script, we group by 5 (120 / 24) so that for each hour/color, four full blocks and one powerline symbol are printed. By omitting partial blocks and only printing 5 characters per hour, an interval of 720 seconds (12 minutes) can be used.
<p>&nbsp;</p>

### Rotate colors
```bash
#!/usr/bin/env bash

declare -a RAIN2=(
"%{F#00fff9}"
"%{F#00fff0}"
...132 omitted...
"%{F#00faff}"
)

LEN=${#RAIN2[@]}
UNIX=$(($(date +%s) % $3))
EIGHTH=$((10#$2*8/1000000))
SPACES=$(($3-$1-1))
(($EIGHTH)) && printf -v PORTION '\\U258%X' $((16 - $EIGHTH)) || PORTION=" "
for ((i=0; i<$1; i++)); do WHOLE+=${RAIN2[($i+$UNIX) % $LEN]}█; done
WHOLE+=${RAIN2[($1+$UNIX) % $LEN]}
printf '%s%b%*s' "$WHOLE" "$PORTION" $SPACES ''
```
<img alt="moving colors" src="loop.gif">

To animate the bar we can use a continuous color array where the final color is similar to the first color so that it can be seamlessly looped. We then get the Unix time in seconds, modulo by the length we are using for the bar, and add that to the loop variable, then modulo by the length of the bar again. With a 1 second interval, the colors will rotate through the blocks of the bar.
<p>&nbsp;</p>

## Creating arrays of colors
There are a lot of color/gradient/CSS tools available but they don't generate lists of 100+ colors, so you would probably have to modify their code in order to use their color interpolation algorithms for this purpose. [This tool](https://medialab.github.io/iwanthue/) is the only one I have found that is ready to use. It generates long lists of hex colors and you can sort the colors by properties such as hue and chroma. I have gone up to 300 colors with 'hard' mode and it was pretty quick. Note that some settings will throw a silent error, so check the console if you think it didn't work.

What I have done so far to generate gradient arrays is more manual but works very well:
1. Find or generate a gradient image online.
2. Download or screenshot it.
3. Cut, modify, and/or scale its dimensions to have a `width` equal to the number of characters you are using for this module.
4. Run the `getColors` program on the image to get the colors printed out in a ready-to-copy format. Run `getColors -h` for usage. The source code is in the [gradient folder](/gradient/).

The following [ImageMagick](https://imagemagick.org/index.php) commands may be helpful:
* `magick -size 135x1 -define gradient:angle=90 gradient:#f1bb12-#1234be image.png` will generate a two-color, 135x1 gradient image ready to provide to `getColors`.
* `magick image.png -resize 135x output.png` will scale an image, preserving its aspect ratio. I used this command on a 235px gradient I created with the above command to reduce it to my time bar size (135 characters) and it generated almost the exact result I got by using [GIMP](https://www.gimp.org/) to scale.
* `mogrify -crop 200x20+300+748 ./image.jpg` crops an image in place. This example moves right 300px, down 748px from the upper left corner, then cuts a 200x20 strip and overwrites the original.