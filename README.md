# agenda-maker
Agenda Maker for Mary

# Notes

The "Read for Meeting" section is generated based on the value of the column "Circulate". If the value for it on a given
row is "yes" then the row will be added to this section and removed from the preliminary report.

# Usage

- From the release section of this repository, download the binary that's targeted to work on your computer. agenda-maker-x64 for X64 architectures, or agenda-maker-arm for ARM ones (like M1 Macs).
- Open the Terminal application of your computer.
- In the Terminal, navigate to the folder where you downloaded the program. For example, if the program is located in your "Downloads" folder type: `cd ~/Downloads`
- Allow the file to be executed by running `chmod +x agenda-maker-arm`. If you are on a Mac you also need to [allow apps from "unindentified developers to run.](https://www.howtogeek.com/205393/gatekeeper-101-why-your-mac-only-allows-apple-approved-software-by-default)"
- Place the CSV file that you want to use to generate the report in the same directory as the downloaded release.
- Run the program by typing: `./agenda-maker-arm --file=NAMEOFTHEFILE.CSV`

- Example:
`./agenda-maker-arm --file=reports.csv`

- Open the generated Agenda file in Word.
- If prompted, select Unicode(UTF8) as encoding.
- Edit with desired style and save as docx.
