# wsl-iperf
<i> Issue: By default powershell commands are not allowed to run in elevation, a user would need to manually set the execution policy to allow unsigned scripts 
<br>
The iPerf test section should be working though as does not require elevation.
</i>
<br><br>
This project is built as a user-friendly version of running iPerf3 tests using the Windows Subsystem for Linux. <br>
It's still new, no clean-up has been done, just getting things to work and still need to test in regards to the features for install as need to set up a blank VM. <br>

### Features still coming:
* <s>Separate iPerf into new window</s><i> tab system implemented. </i>
* <s>iPerf window needs to contain all options for an iPerf test that are usually used</s>
* Include an option in the iperf window to add your own arguments
* Include an option to: directly dump result to file and to <s>clipboard.</s>
* Include an option to "fancy" format the results so it's easier to look at (maybe auto syntax highlighter) <i>waiting on fyne implementation of md viewer </i>
* Maybe a clean-up option to remove ubuntu install
* Different icon
* <s>Placeholder text for Entries</s>
* Swap between dark and light theme as Fyne has it built-in

#### Maybe: https://github.com/RoliSoft/WSL-Distribution-Switcher to run docker instance instead
<br>

Way down the line I'll consider adding a MacOS variant. 

<br>

### Changes once WSL 2 gets released
* Add features like mtr, traceroute, etc. that currently cannot be used due to WSL not being a full linux kernel. 
